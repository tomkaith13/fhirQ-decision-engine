package figma

import (
	"github.com/google/fhir/go/proto/google/fhir/proto/r4/core/codes_go_proto"
	"github.com/google/fhir/go/proto/google/fhir/proto/r4/core/datatypes_go_proto"
	"github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/questionnaire_go_proto"
)

func dfs(node *FigmaNode) *questionnaire_go_proto.Questionnaire_Item {
	if node == nil {
		return nil
	}

	qItem := &questionnaire_go_proto.Questionnaire_Item{}
	qItem.Id = &datatypes_go_proto.String{Value: node.Id}

	switch node.Type {
	case "FRAME":
		qItem.Type = &questionnaire_go_proto.Questionnaire_Item_TypeCode{}
		qItem.Type.Value = codes_go_proto.QuestionnaireItemTypeCode_GROUP

		if node.Name == "fhir_choice" {
			// now we need to not recurse and build more nodes by add answerOption
			qItem.AnswerOption = []*questionnaire_go_proto.Questionnaire_Item_AnswerOption{}
			for _, innerC := range node.Children {
				option := questionnaire_go_proto.Questionnaire_Item_AnswerOption{
					Value: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX{
						Choice: &questionnaire_go_proto.Questionnaire_Item_AnswerOption_ValueX_StringValue{
							StringValue: &datatypes_go_proto.String{
								Value: innerC.Name}}},
				}
				qItem.AnswerOption = append(qItem.AnswerOption, &option)
			}
			return qItem

		}

	case "TEXT":
		qItem.Text = &datatypes_go_proto.String{Value: node.Name}
	}

	// TBD add more processing here.

	for _, c := range node.Children {
		qItem.Item = append(qItem.Item, dfs(&c))
	}
	return qItem
}
func FigmaToFhirQuestionnaireConvertor(tree FigmaDoc) *questionnaire_go_proto.Questionnaire {
	q := &questionnaire_go_proto.Questionnaire{}

	// we DFS the tree to extract and group questionnaire items
	// the key is to look for `name` fields in children that start with the prefix `fhir_`. This is
	// something the figma designer needs to keep in mind.
	// All nodes without this prefix will be dropped from processing and adding to

	//  All type FRAME will be consider and a starting point to create a new GROUP in FHIRQ
	// All types TEXT or INSTANCE is used to generate a radio button.
	for _, c := range tree.Document.Children {
		q.Item = append(q.Item, dfs(&c))
	}
	return q
}
