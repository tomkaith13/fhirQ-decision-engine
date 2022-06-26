package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/fhir/go/fhirversion"
	"github.com/google/fhir/go/jsonformat"
	r4pb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/bundle_and_contained_resource_go_proto"
	"github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/questionnaire_go_proto"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/questionnaire", func(w http.ResponseWriter, r *http.Request) {
		questionnaireJson, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}
		unmarshaller, err := jsonformat.NewUnmarshaller(time.UTC.String(), fhirversion.R4)

		unmarshalled, err := unmarshaller.Unmarshal(questionnaireJson)
		if err != nil {
			return
		}
		containedResource := unmarshalled.(*r4pb.ContainedResource)
		q := containedResource.GetQuestionnaire()
		for _, item := range q.GetItem() {
			nestedItemParser(item)
		}

		id := q.GetId().Value
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(id))
	})

	http.ListenAndServe(":8080", r)
}

func nestedItemParser(item *questionnaire_go_proto.Questionnaire_Item) {
	fmt.Println(item.GetText())
	fmt.Println(item.GetEnableWhen())
	if item.GetItem() != nil {
		for _, i := range item.GetItem() {
			nestedItemParser(i)
		}
	}
}
