package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/fhir/go/fhirversion"
	"github.com/google/fhir/go/jsonformat"
	"github.com/google/fhir/go/proto/google/fhir/proto/r4/core/datatypes_go_proto"
	r4pb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/bundle_and_contained_resource_go_proto"
	"github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/questionnaire_go_proto"
	"github.com/tomkaith13/fhirQuestionnaireEngine/src/questionnaire_collection"
	"github.com/tomkaith13/fhirQuestionnaireEngine/src/questionnaire_resp_model"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	qMap := make(map[string]*questionnaire_go_proto.Questionnaire)

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
		if _, ok := qMap[q.Id.Value]; ok {
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
		if _, ok := questionnaire_collection.FhirQmap[id]; !ok {
			questionnaire_collection.FhirQmap[id] = *q
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
		qMap[q.Id.Value] = q
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(resp))
	})

	r.Get("/questionnaire/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		questionnaire, ok := qMap[id]
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

		qr.Identifier.Value = "test123"

		// Insert decision logic for a Questionnaire here!!!

		fmt.Println(qr.GetId())
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("{\"id\": \"" + qr.GetId().Value + "\"}"))
	})
	http.ListenAndServe(":8080", r)
}
