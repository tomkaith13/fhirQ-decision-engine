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

	r.Get("/mocked-questionnaire/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		q := &questionnaire_go_proto.Questionnaire{}
		switch id {
		case "hm1":
			q.Id = &datatypes_go_proto.Id{Value: "hm1"}
			// Title of page
			q.Title = &datatypes_go_proto.String{
				Value: "Select your plan",
			}
			semVer, _ := semver.NewVersion("1.0.0")
			// subtitle
			// qi1 := &questionnaire_go_proto.Questionnaire_Item{
			// 	LinkId:   &datatypes_go_proto.String{Value: semVer.String()},
			// 	Required: &datatypes_go_proto.Boolean{Value: true},
			// 	Text:     &datatypes_go_proto.String{Value: "Select your plan"},
			// 	Type: &questionnaire_go_proto.Questionnaire_Item_TypeCode{
			// 		Value: codes_go_proto.QuestionnaireItemTypeCode_DISPLAY,
			// 	},
			// }
			// q.Item = append(q.Item, qi1)

			// secondary text
			// subtitle
			qi1 := &questionnaire_go_proto.Questionnaire_Item{
				LinkId:   &datatypes_go_proto.String{Value: semVer.String()},
				Required: &datatypes_go_proto.Boolean{Value: true},
				Text:     &datatypes_go_proto.String{Value: "To help ensure we connect you with the right agent, let us know the plan this is related to"},
				Type: &questionnaire_go_proto.Questionnaire_Item_TypeCode{
					Value: codes_go_proto.QuestionnaireItemTypeCode_DISPLAY,
				},
			}
			q.Item = append(q.Item, qi1)

			// Radio buttons
			semVer1 := semVer.IncMinor()
			semVer2 := semVer1.IncMinor()
			semVer3 := semVer2.IncMinor()
			semVer4 := semVer3.IncMinor()

			answers := []*questionnaire_go_proto.Questionnaire_Item_AnswerOption{
				&questionnaire_go_proto.Questionnaire_Item_AnswerOption{
					Value: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX{
						Choice: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX_Coding{
							Coding: &datatypes_go_proto.Coding{
								Id: &datatypes_go_proto.String{Value: semVer2.String()},
								Code: &datatypes_go_proto.Code{
									Value: "dental-wellness-product",
								},
								Display: &datatypes_go_proto.String{Value: "Dental Wellness Product"},
							},
						},
					},
				},
				&questionnaire_go_proto.Questionnaire_Item_AnswerOption{
					Value: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX{
						Choice: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX_Coding{
							Coding: &datatypes_go_proto.Coding{
								Id: &datatypes_go_proto.String{Value: semVer3.String()},
								Code: &datatypes_go_proto.Code{
									Value: "drug-ppo-blue",
								},
								Display: &datatypes_go_proto.String{Value: "Drug-PPO-Blue"},
							},
						},
					},
				},
				&questionnaire_go_proto.Questionnaire_Item_AnswerOption{
					Value: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX{
						Choice: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX_Coding{
							Coding: &datatypes_go_proto.Coding{
								Id: &datatypes_go_proto.String{Value: semVer4.String()},
								Code: &datatypes_go_proto.Code{
									Value: "drug-ppo-red",
								},
								Display: &datatypes_go_proto.String{Value: "Drug-PPO-Red"},
							},
						},
					},
				},
			}
			qi3 := &questionnaire_go_proto.Questionnaire_Item{
				LinkId:   &datatypes_go_proto.String{Value: semVer1.String()},
				Required: &datatypes_go_proto.Boolean{Value: true},
				Type: &questionnaire_go_proto.Questionnaire_Item_TypeCode{
					Value: codes_go_proto.QuestionnaireItemTypeCode_CHOICE,
				},
				Text: &datatypes_go_proto.String{
					Value: "Plans",
				},
				AnswerOption: answers,
				Extension: []*datatypes_go_proto.Extension{
					&datatypes_go_proto.Extension{
						Url: &datatypes_go_proto.Uri{
							Value: "http://hl7.org/fhir/StructureDefinition/questionnaire-itemControl",
						},
						Value: &datatypes_go_proto.Extension_ValueX{
							Choice: &datatypes_go_proto.Extension_ValueX_CodeableConcept{
								CodeableConcept: &datatypes_go_proto.CodeableConcept{
									Coding: []*datatypes_go_proto.Coding{
										{
											System: &datatypes_go_proto.Uri{
												Value: "http://hl7.org/fhir/ValueSet/questionnaire-item-control",
											},
											Code: &datatypes_go_proto.Code{
												Value: "radio-button",
											},
										},
									},
								},
							},
						},
					},
				},
			}
			q.Item = append(q.Item, qi3)
		case "hm2":
			q.Id = &datatypes_go_proto.Id{Value: "hm2"}
			// Title of page
			q.Title = &datatypes_go_proto.String{
				Value: "Callback Preference",
			}
			semVer, _ := semver.NewVersion("1.0.0")
			// subtitle
			// qi1 := &questionnaire_go_proto.Questionnaire_Item{
			// 	LinkId:   &datatypes_go_proto.String{Value: semVer.String()},
			// 	Required: &datatypes_go_proto.Boolean{Value: true},
			// 	Text:     &datatypes_go_proto.String{Value: "Callback Preference"},
			// 	Type: &questionnaire_go_proto.Questionnaire_Item_TypeCode{
			// 		Value: codes_go_proto.QuestionnaireItemTypeCode_DISPLAY,
			// 	},
			// }
			// q.Item = append(q.Item, qi1)

			// secondary text
			// subtitle
			qi1 := &questionnaire_go_proto.Questionnaire_Item{
				LinkId:   &datatypes_go_proto.String{Value: semVer.String()},
				Required: &datatypes_go_proto.Boolean{Value: true},
				Text:     &datatypes_go_proto.String{Value: "We may need to call you about personal health information. If we cannot reach you via our message centre, let us know the best time to reach you"},
				Type: &questionnaire_go_proto.Questionnaire_Item_TypeCode{
					Value: codes_go_proto.QuestionnaireItemTypeCode_DISPLAY,
				},
			}
			q.Item = append(q.Item, qi1)

			// phone number
			semVer1 := semVer.IncMinor()
			qi3 := &questionnaire_go_proto.Questionnaire_Item{
				LinkId:   &datatypes_go_proto.String{Value: semVer1.String()},
				Required: &datatypes_go_proto.Boolean{Value: true},
				Text:     &datatypes_go_proto.String{Value: "Phone"},
				Type: &questionnaire_go_proto.Questionnaire_Item_TypeCode{
					Value: codes_go_proto.QuestionnaireItemTypeCode_STRING,
				},
				Initial: []*questionnaire_go_proto.Questionnaire_Item_Initial{&questionnaire_go_proto.Questionnaire_Item_Initial{
					Value: &questionnaire_go_proto.Questionnaire_Item_Initial_ValueX{
						Choice: &questionnaire_go_proto.Questionnaire_Item_Initial_ValueX_StringValue{
							StringValue: &datatypes_go_proto.String{Value: "123-456-7890"},
						}},
				},
				},
			}
			q.Item = append(q.Item, qi3)

			// dropbox
			semVer2 := semVer1.IncMinor()
			semVer3 := semVer2.IncMinor()
			semVer4 := semVer3.IncMinor()
			semVer5 := semVer4.IncMinor()

			answers := []*questionnaire_go_proto.Questionnaire_Item_AnswerOption{
				&questionnaire_go_proto.Questionnaire_Item_AnswerOption{
					Value: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX{
						Choice: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX_Coding{
							Coding: &datatypes_go_proto.Coding{
								Id: &datatypes_go_proto.String{Value: semVer3.String()},
								Code: &datatypes_go_proto.Code{
									Value: "8-a.m.---11-a.m.-est",
								},
								Display: &datatypes_go_proto.String{Value: "8 a.m - 11 a.m EST"},
							},
						},
					},
				},
				&questionnaire_go_proto.Questionnaire_Item_AnswerOption{
					Value: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX{
						Choice: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX_Coding{
							Coding: &datatypes_go_proto.Coding{
								Id: &datatypes_go_proto.String{Value: semVer4.String()},
								Code: &datatypes_go_proto.Code{
									Value: "11-a.m.---2-p.m.-est",
								},
								Display: &datatypes_go_proto.String{Value: "11 a.m - 2 p.m EST"},
							},
						},
					},
				},
				&questionnaire_go_proto.Questionnaire_Item_AnswerOption{
					Value: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX{
						Choice: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX_Coding{
							Coding: &datatypes_go_proto.Coding{
								Id: &datatypes_go_proto.String{Value: semVer5.String()},
								Code: &datatypes_go_proto.Code{
									Value: "2-p.m.---5-p.m.-est",
								},
								Display: &datatypes_go_proto.String{Value: "2 p.m - 5 p.m EST"},
							},
						},
					},
				},
			}
			qi4 := &questionnaire_go_proto.Questionnaire_Item{
				LinkId:   &datatypes_go_proto.String{Value: semVer2.String()},
				Required: &datatypes_go_proto.Boolean{Value: true},
				Type: &questionnaire_go_proto.Questionnaire_Item_TypeCode{
					Value: codes_go_proto.QuestionnaireItemTypeCode_CHOICE,
				},
				Text: &datatypes_go_proto.String{
					Value: "Best time to call",
				},
				AnswerOption: answers,
				Extension: []*datatypes_go_proto.Extension{
					&datatypes_go_proto.Extension{
						Url: &datatypes_go_proto.Uri{
							Value: "http://hl7.org/fhir/StructureDefinition/questionnaire-itemControl",
						},
						Value: &datatypes_go_proto.Extension_ValueX{
							Choice: &datatypes_go_proto.Extension_ValueX_CodeableConcept{
								CodeableConcept: &datatypes_go_proto.CodeableConcept{
									Coding: []*datatypes_go_proto.Coding{
										{
											System: &datatypes_go_proto.Uri{
												Value: "http://hl7.org/fhir/ValueSet/questionnaire-item-control",
											},
											Code: &datatypes_go_proto.Code{
												Value: "drop-down",
											},
										},
									},
								},
							},
						},
					},
				},
			}
			q.Item = append(q.Item, qi4)

			// confirmation radio-button
			semVer6 := semVer5.IncMinor()
			// answers2 := []*questionnaire_go_proto.Questionnaire_Item_AnswerOption{
			// 	&questionnaire_go_proto.Questionnaire_Item_AnswerOption{
			// 		Value: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX{
			// 			Choice: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX_StringValue{
			// 				StringValue: &datatypes_go_proto.String{
			// 					Value: "Yes",
			// 				},
			// 			},
			// 		},
			// 	},
			// 	&questionnaire_go_proto.Questionnaire_Item_AnswerOption{
			// 		Value: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX{
			// 			Choice: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX_StringValue{
			// 				StringValue: &datatypes_go_proto.String{
			// 					Value: "No",
			// 				},
			// 			},
			// 		},
			// 	},
			// }
			qi5 := &questionnaire_go_proto.Questionnaire_Item{
				LinkId:   &datatypes_go_proto.String{Value: semVer6.String()},
				Required: &datatypes_go_proto.Boolean{Value: true},
				Type: &questionnaire_go_proto.Questionnaire_Item_TypeCode{
					Value: codes_go_proto.QuestionnaireItemTypeCode_BOOLEAN,
				},
				Text: &datatypes_go_proto.String{
					Value: "May we leave a message if you’re not available",
				},
				// AnswerOption: answers2,
				Extension: []*datatypes_go_proto.Extension{
					&datatypes_go_proto.Extension{
						Url: &datatypes_go_proto.Uri{
							Value: "http://hl7.org/fhir/StructureDefinition/questionnaire-itemControl",
						},
						Value: &datatypes_go_proto.Extension_ValueX{
							Choice: &datatypes_go_proto.Extension_ValueX_CodeableConcept{
								CodeableConcept: &datatypes_go_proto.CodeableConcept{
									Coding: []*datatypes_go_proto.Coding{
										{
											System: &datatypes_go_proto.Uri{
												Value: "http://hl7.org/fhir/ValueSet/questionnaire-item-control",
											},
											Code: &datatypes_go_proto.Code{
												Value: "radio-button",
											},
										},
									},
								},
							},
						},
					},
				},
			}
			q.Item = append(q.Item, qi5)
		case "opm":
			q.Id = &datatypes_go_proto.Id{Value: "opm"}
			// Title of page
			q.Title = &datatypes_go_proto.String{
				Value: "Select Participants",
			}

			// secondary text
			// Subtitle
			semVer, _ := semver.NewVersion("1.0.0")
			qi1 := &questionnaire_go_proto.Questionnaire_Item{
				LinkId:   &datatypes_go_proto.String{Value: semVer.String()},
				Required: &datatypes_go_proto.Boolean{Value: true},
				Text:     &datatypes_go_proto.String{Value: "Select people on your plan that you’d like to include in this message."},
				Type: &questionnaire_go_proto.Questionnaire_Item_TypeCode{
					Value: codes_go_proto.QuestionnaireItemTypeCode_DISPLAY,
				},
			}
			q.Item = append(q.Item, qi1)

			//EditText
			semVer1 := semVer.IncMinor()
			qi2 := &questionnaire_go_proto.Questionnaire_Item{
				LinkId:   &datatypes_go_proto.String{Value: semVer1.String()},
				Required: &datatypes_go_proto.Boolean{Value: false},
				Text:     &datatypes_go_proto.String{Value: "Message subject"},
				Type: &questionnaire_go_proto.Questionnaire_Item_TypeCode{
					Value: codes_go_proto.QuestionnaireItemTypeCode_STRING,
				},
			}
			q.Item = append(q.Item, qi2)

			// Check boxes
			semVer2 := semVer1.IncMinor()
			semVer3 := semVer2.IncMinor()
			semVer4 := semVer3.IncMinor()
			semVer5 := semVer4.IncMinor()
			semVer6 := semVer5.IncMinor()

			answers := []*questionnaire_go_proto.Questionnaire_Item_AnswerOption{
				&questionnaire_go_proto.Questionnaire_Item_AnswerOption{
					Value: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX{
						Choice: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX_Coding{
							Coding: &datatypes_go_proto.Coding{
								Id: &datatypes_go_proto.String{Value: semVer3.String()},
								Code: &datatypes_go_proto.Code{
									Value: "cathy-c.",
								},
								Display: &datatypes_go_proto.String{Value: "Cathy C."},
							},
						},
					},
				},
				&questionnaire_go_proto.Questionnaire_Item_AnswerOption{
					Value: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX{
						Choice: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX_Coding{
							Coding: &datatypes_go_proto.Coding{
								Id: &datatypes_go_proto.String{Value: semVer4.String()},
								Code: &datatypes_go_proto.Code{
									Value: "arjun-n",
								},
								Display: &datatypes_go_proto.String{Value: "Arjun N"},
							},
						},
					},
				},
				&questionnaire_go_proto.Questionnaire_Item_AnswerOption{
					Value: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX{
						Choice: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX_Coding{
							Coding: &datatypes_go_proto.Coding{
								Id: &datatypes_go_proto.String{Value: semVer5.String()},
								Code: &datatypes_go_proto.Code{
									Value: "bibin-thomas",
								},
								Display: &datatypes_go_proto.String{Value: "Bibin Thomas"},
							},
						},
					},
				},
				&questionnaire_go_proto.Questionnaire_Item_AnswerOption{
					Value: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX{
						Choice: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX_Coding{
							Coding: &datatypes_go_proto.Coding{
								Id: &datatypes_go_proto.String{Value: semVer6.String()},
								Code: &datatypes_go_proto.Code{
									Value: "evan-holtrop",
								},
								Display: &datatypes_go_proto.String{Value: "Evan Holtrop"},
							},
						},
					},
				},
			}
			qi3 := &questionnaire_go_proto.Questionnaire_Item{
				LinkId:   &datatypes_go_proto.String{Value: semVer2.String()},
				Required: &datatypes_go_proto.Boolean{Value: true},
				Type: &questionnaire_go_proto.Questionnaire_Item_TypeCode{
					Value: codes_go_proto.QuestionnaireItemTypeCode_CHOICE,
				},
				Text: &datatypes_go_proto.String{
					Value: "Participants",
				},
				AnswerOption: answers,
				Extension: []*datatypes_go_proto.Extension{
					&datatypes_go_proto.Extension{
						Url: &datatypes_go_proto.Uri{
							Value: "http://hl7.org/fhir/StructureDefinition/questionnaire-itemControl",
						},
						Value: &datatypes_go_proto.Extension_ValueX{
							Choice: &datatypes_go_proto.Extension_ValueX_CodeableConcept{
								CodeableConcept: &datatypes_go_proto.CodeableConcept{
									Coding: []*datatypes_go_proto.Coding{
										{
											System: &datatypes_go_proto.Uri{
												Value: "http://hl7.org/fhir/ValueSet/questionnaire-item-control",
											},
											Code: &datatypes_go_proto.Code{
												Value: "check-box",
											},
										},
									},
								},
							},
						},
					},
				},
			}
			q.Item = append(q.Item, qi3)
		case "generic":
			// TODO

		}
		//formatter - the efficient one
		// pMarshaller, err := jsonformat.NewPrettyMarshaller(fhirversion.R4)
		marshaller, err := jsonformat.NewMarshaller(false, "", "", fhirversion.R4)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		rsc := r4pb.ContainedResource{OneofResource: &r4pb.ContainedResource_Questionnaire{Questionnaire: q}}
		// resp, err := pMarshaller.MarshalResourceToString(q)
		resp, err := marshaller.Marshal(&rsc)
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
					// How to make custom extensions
					//
					// Value: &datatypes_go_proto.Extension_ValueX{
					// 	Choice: &datatypes_go_proto.Extension_ValueX_StringValue{
					// 		StringValue: &datatypes_go_proto.String{
					// 			Value: "style:check-box",
					// 		},
					// 	},
					// },
					Value: &datatypes_go_proto.Extension_ValueX{
						// How to use Codings directly
						//
						// Choice: &datatypes_go_proto.Extension_ValueX_Coding{
						// 	Coding: &datatypes_go_proto.Coding{
						// 		Code: &datatypes_go_proto.Code{
						// 			Value: stu3Codes.QuestionnaireItemUIControlCodesCode_CHECK_BOX.String(),
						// 		},
						// 	},
						// },
						Choice: &datatypes_go_proto.Extension_ValueX_CodeableConcept{
							// How to use CodeableConcepts
							CodeableConcept: &datatypes_go_proto.CodeableConcept{
								Coding: []*datatypes_go_proto.Coding{
									{
										System: &datatypes_go_proto.Uri{
											Value: "http://hl7.org/fhir/ValueSet/questionnaire-item-control",
										},
										Code: &datatypes_go_proto.Code{
											Value: stu3Codes.QuestionnaireItemUIControlCodesCode_CHECK_BOX.String(),
										},
									},
								},
							},
						},
					},
					// Url: &datatypes_go_proto.Uri{Value: "http://hl7.org/fhir/StructureDefinition/questionnaire-itemControl"},
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
