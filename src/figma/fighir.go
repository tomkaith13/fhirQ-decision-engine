package figma

import "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/questionnaire_go_proto"

func FigmaToFhirQuestionnaireConvertor(tree FigmaDoc) *questionnaire_go_proto.Questionnaire {
	q := &questionnaire_go_proto.Questionnaire{}
	return q
}
