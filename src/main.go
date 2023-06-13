package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/fhir/go/fhirversion"
	"github.com/google/fhir/go/jsonformat"
	"github.com/google/fhir/go/proto/google/fhir/proto/r4/core/codes_go_proto"
	"github.com/google/fhir/go/proto/google/fhir/proto/r4/core/datatypes_go_proto"
	r4pb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/bundle_and_contained_resource_go_proto"
	"github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/questionnaire_go_proto"
	stu3Codes "github.com/google/fhir/go/proto/google/fhir/proto/stu3/codes_go_proto"
	"github.com/masterminds/semver"
	"github.com/tomkaith13/fhirQuestionnaireEngine/src/figma"
	"github.com/tomkaith13/fhirQuestionnaireEngine/src/questionnaire_collection"
	"github.com/tomkaith13/fhirQuestionnaireEngine/src/questionnaire_resp_model"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// questionnaire_collection.QMap := make(map[string]*questionnaire_go_proto.Questionnaire)

	r.Post("/questionnaire", func(w http.ResponseWriter, r *http.Request) {
		questionnaireJson, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}
		unmarshaller, err := jsonformat.NewUnmarshaller(time.UTC.String(), fhirversion.R4)

		unmarshalled, err := unmarshaller.Unmarshal(questionnaireJson)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Cant unmarshall FHIR Questionnaire"))
			return
		}
		containedResource := unmarshalled.(*r4pb.ContainedResource)

		q := containedResource.GetQuestionnaire()
		if _, ok := questionnaire_collection.QMap[q.Id.Value]; ok {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("id already exists"))
			return
		}
		custom := "my own metadata"
		q.Item = append(q.Item, &questionnaire_go_proto.Questionnaire_Item{
			Id: &datatypes_go_proto.String{
				Value: custom,
			},
		})
		id := q.GetId().Value
		if _, ok := questionnaire_collection.QMap[id]; !ok {
			questionnaire_collection.QMap[id] = q
		}
		pMarshaller, err := jsonformat.NewPrettyMarshaller(fhirversion.R4)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		resp, err := pMarshaller.MarshalResourceToString(q)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		questionnaire_collection.QMap[q.Id.Value] = q
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(resp))
	})

	r.Get("/questionnaire/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		questionnaire, ok := questionnaire_collection.QMap[id]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("questionnaire does not exist"))
			return
		}
		pMarshaller, err := jsonformat.NewPrettyMarshaller(fhirversion.R4)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		resp, err := pMarshaller.MarshalResourceToString(questionnaire)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp))
	})

	r.Get("/transform-from-json-1", func(w http.ResponseWriter, r *http.Request) {
		tenantJson := `{
			"da_reasons": null,
			"workflow_node": {
			  "da_workflow_node_id": 0,
			  "workflow_questions": [
				{
				  "da_workflow_question_id": 0,
				  "workflow_question_text": "How old are you",
				  "workflow_question_ui": "date",
				  "workflow_question_tag": "DOB",
				  "workflow_question_answer_data_type": "string",
				  "workflow_question_answers": [
					{
					  "da_workflow_question_answer_id": 0,
					  "workflow_question_answer_text": "string",
					  "workflow_question_answer_value": "string"
					}
				  ]
				}
			  ]
			},
			"da_time": null,
			"message": null
		  }`

		tenantQ := &questionnaire_collection.DaQuesionnaire{}
		err := json.Unmarshal([]byte(tenantJson), tenantQ)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		fmt.Printf("%+v", tenantQ)

		q := &questionnaire_go_proto.Questionnaire{}
		Qnode := tenantQ.QNode
		sId := strconv.Itoa(Qnode.Id)
		q.Id = &datatypes_go_proto.Id{Value: sId}

		// now add questions as items
		questions := Qnode.QQuestions

		for _, question := range questions {
			qId := strconv.Itoa(question.Id)
			qItem := &questionnaire_go_proto.Questionnaire_Item{
				Id: &datatypes_go_proto.String{
					Value: qId,
				},
				Text: &datatypes_go_proto.String{
					Value: question.QText,
				},
			}
			// add custom mappers
			if question.QUICode == "date" {
				qItem.Extension = append(qItem.Extension, &datatypes_go_proto.Extension{
					Id: &datatypes_go_proto.String{Value: "0.1"},
					Value: &datatypes_go_proto.Extension_ValueX{
						Choice: &datatypes_go_proto.Extension_ValueX_StringValue{StringValue: &datatypes_go_proto.String{
							Value: "date",
						}},
					},
				})
			}
			q.Item = append(q.Item, qItem)
		}

		fmt.Printf("%+v", q)

		// formatter
		pMarshaller, err := jsonformat.NewPrettyMarshaller(fhirversion.R4)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		resp, err := pMarshaller.MarshalResourceToString(q)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp))
	})

	r.Get("/transform-from-json-2", func(w http.ResponseWriter, r *http.Request) {
		tenantJson := `
		{
			"familyMembers": [
				{
					"middleName": null,
					"lastName": "Sukla",
					"firstName": "Ahuja"
				},
				{
					"middleName": null,
					"lastName": "Sukla",
					"firstName": "Ankith"
				},
				{
					"middleName": null,
					"lastName": "shukla",
					"firstName": "ranjith"
				}
			],
			"careAdvisor": {
				"role": "Care Advisor",
				"lastName": "Care Advisor",
				"firstName": "Sheelrathan"
			}
		}
		`
		tenantQ := &questionnaire_collection.FamilyAccount{}
		err := json.Unmarshal([]byte(tenantJson), tenantQ)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		fmt.Printf("%+v", *tenantQ)

		q := &questionnaire_go_proto.Questionnaire{}
		q.Id = &datatypes_go_proto.Id{Value: "1"}
		q.Title = &datatypes_go_proto.String{
			Value: "New Message",
		}
		q.Extension = append(q.Extension, &datatypes_go_proto.Extension{
			Value: &datatypes_go_proto.Extension_ValueX{
				Choice: &datatypes_go_proto.Extension_ValueX_StringValue{
					StringValue: &datatypes_go_proto.String{
						Value: "submit-text:Create message",
					},
				},
			},
		})
		q.Extension = append(q.Extension, &datatypes_go_proto.Extension{
			Value: &datatypes_go_proto.Extension_ValueX{
				Choice: &datatypes_go_proto.Extension_ValueX_StringValue{
					StringValue: &datatypes_go_proto.String{
						Value: "button-location:bottom",
					},
				},
			},
		})

		semVer, _ := semver.NewVersion("1.1.0")

		q.Item = append(q.Item, &questionnaire_go_proto.Questionnaire_Item{
			Id:       &datatypes_go_proto.String{Value: semVer.String()},
			Required: &datatypes_go_proto.Boolean{Value: false},
			Text:     &datatypes_go_proto.String{Value: "Message Subject"},
			Type: &questionnaire_go_proto.Questionnaire_Item_TypeCode{
				Value: codes_go_proto.QuestionnaireItemTypeCode_TEXT,
			},
			Initial: []*questionnaire_go_proto.Questionnaire_Item_Initial{&questionnaire_go_proto.Questionnaire_Item_Initial{
				Value: &questionnaire_go_proto.Questionnaire_Item_Initial_ValueX{
					Choice: &questionnaire_go_proto.Questionnaire_Item_Initial_ValueX_StringValue{
						StringValue: &datatypes_go_proto.String{Value: "Set the message topic"},
					}},
			},
			},
		})

		// now adding care advisor
		answers := []*questionnaire_go_proto.Questionnaire_Item_AnswerOption{
			&questionnaire_go_proto.Questionnaire_Item_AnswerOption{
				Value: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX{
					Choice: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX_StringValue{
						StringValue: &datatypes_go_proto.String{
							Value: tenantQ.CareAdvisor.FirstName + " " + tenantQ.CareAdvisor.LastName,
						},
					},
				},
				Extension: []*datatypes_go_proto.Extension{
					&datatypes_go_proto.Extension{
						Value: &datatypes_go_proto.Extension_ValueX{
							Choice: &datatypes_go_proto.Extension_ValueX_StringValue{
								StringValue: &datatypes_go_proto.String{
									Value: "secondary-text:Care Advisor",
								},
							},
						},
					},
				},
			},
		}

		for _, member := range tenantQ.FamilyMembers {
			answers = append(answers, &questionnaire_go_proto.Questionnaire_Item_AnswerOption{
				Value: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX{
					Choice: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX_StringValue{
						StringValue: &datatypes_go_proto.String{
							Value: member.FirstName + " " + member.LastName,
						},
					},
				},
			})
		}

		newVersion := semVer.IncMinor()

		q.Item = append(q.Item, &questionnaire_go_proto.Questionnaire_Item{
			Id:       &datatypes_go_proto.String{Value: newVersion.String()},
			Required: &datatypes_go_proto.Boolean{Value: true},
			Type: &questionnaire_go_proto.Questionnaire_Item_TypeCode{
				Value: codes_go_proto.QuestionnaireItemTypeCode_OPEN_CHOICE,
			},
			Text: &datatypes_go_proto.String{
				Value: "To",
			},
			Extension: []*datatypes_go_proto.Extension{
				&datatypes_go_proto.Extension{
					// Value: &datatypes_go_proto.Extension_ValueX{
					// 	Choice: &datatypes_go_proto.Extension_ValueX_StringValue{
					// 		StringValue: &datatypes_go_proto.String{
					// 			Value: "style:check-box",
					// 		},
					// 	},
					// },
					Value: &datatypes_go_proto.Extension_ValueX{
						Choice: &datatypes_go_proto.Extension_ValueX_Coding{
							Coding: &datatypes_go_proto.Coding{
								Code: &datatypes_go_proto.Code{
									Value: stu3Codes.QuestionnaireItemUIControlCodesCode_CHECK_BOX.String(),
								},
							},
						},
					},
					Url: &datatypes_go_proto.Uri{Value: "http://hl7.org/fhir/ValueSet/questionnaire-item-control"},
				},
			},
			AnswerOption: answers,
		})

		// formatter
		pMarshaller, err := jsonformat.NewPrettyMarshaller(fhirversion.R4)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		resp, err := pMarshaller.MarshalResourceToString(q)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp))

	})
	r.Post("/questionnaire-resp", func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		qrBody := questionnaire_resp_model.QRBody{}

		err := dec.Decode(&qrBody)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Cant unmarshall custom json"))
			return

		}
		fmt.Println("ids::::::::", qrBody.QuestionnaireIds)

		regJsonMarshalledData, err := json.Marshal(qrBody.QuestionnaireResponse)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Cant unmarshall custom json"))
			return
		}
		unmarshaller, err := jsonformat.NewUnmarshaller(time.UTC.String(), fhirversion.R4)

		unmarshalled, err := unmarshaller.Unmarshal(regJsonMarshalledData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Cant unmarshall FHIR Questionnaire"))
			return
		}

		cr := unmarshalled.(*r4pb.ContainedResource)
		qr := cr.GetQuestionnaireResponse()

		// Insert decision logic for a Questionnaire here!!!
		for _, item := range qr.Item {
			for _, ans := range item.Answer {
				fmt.Println(ans.Value.GetCoding().Code.Value)
				answerVal := ans.Value.GetCoding().Code.Value
				if answerVal == "urgent-care" {
					w.WriteHeader(http.StatusAccepted)
					w.Write([]byte("Next questionnaire id: urgent-menu"))
					return
				}
			}
		}

		fmt.Println(qr.GetId())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\" Unable to process resp id\": \"" + qr.GetId().Value + "\"}"))
	})

	r.Post("/fighir", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
			return
		}

		tree := figma.FigmaDoc{}
		err = json.Unmarshal(body, &tree)
		if err != nil {
			log.Fatalln(err)
			return
		}
		q := figma.FigmaToFhirQuestionnaireConvertor(tree)

		pMarshaller, err := jsonformat.NewPrettyMarshaller(fhirversion.R4)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		resp, err := pMarshaller.MarshalResourceToString(q)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf("%+v", resp)))

	})
	http.ListenAndServe(":8080", r)
}
