package figma

import (
	"github.com/google/fhir/go/proto/google/fhir/proto/r4/core/codes_go_proto"
	"github.com/google/fhir/go/proto/google/fhir/proto/r4/core/datatypes_go_proto"
	"github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/questionnaire_go_proto"
)

// func dfs(node *FigmaNode, qItems []*questionnaire_go_proto.Questionnaire_Item, qItem *questionnaire_go_proto.Questionnaire_Item) []*questionnaire_go_proto.Questionnaire_Item {
// 	if node == nil {
// 		return nil
// 	}
// 	fmt.Printf("\n\nnode:: %+v\n", node)
// 	if qItem == nil {
// 		qItem = &questionnaire_go_proto.Questionnaire_Item{}
// 	}
// 	qItem.Id = &datatypes_go_proto.String{Value: node.Id}
// switch node.Type {
// case "FRAME":
// 	qItem.Type = &questionnaire_go_proto.Questionnaire_Item_TypeCode{}
// 	qItem.Type.Value = codes_go_proto.QuestionnaireItemTypeCode_GROUP

// case "TEXT":
// 	qItem.Text = &datatypes_go_proto.String{Value: node.Name}
// }

// 	for _, c := range node.Children {
// 		// fmt.Printf("index: %v and child: %+v\n", i, c)
// 		fmt.Println(qItem)
// 		qItems = append(qItems, dfs(&c, qItems, qItem)...)

// 	}
// 	fmt.Printf("qItems: %+v", qItems)
// 	return qItems
// }

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
		// dfs(&c, q.Item)
	}
	// fmt.Printf("\n\n\n q::::%+v\n\n", q)
	return q
}
