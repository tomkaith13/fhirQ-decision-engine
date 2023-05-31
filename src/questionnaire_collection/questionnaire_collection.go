package questionnaire_collection

import "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/questionnaire_go_proto"

var QMap = make(map[string]*questionnaire_go_proto.Questionnaire)

type DaAnswer struct {
	Id      int    `json:"da_workflow_question_answer_id"`
	AnsText string `json:"workflow_question_answer_text"`
	AnsVal  string `json:"workflow_question_answer_value"`
}
type DaWorflowQuestion struct {
	Id            int        `json:"da_workflow_question_id"`
	QText         string     `json:"workflow_question_text"`
	QUICode       string     `json:"workflow_question_ui"`
	QTagCode      string     `json:"workflow_question_tag"`
	QAnswerOption []DaAnswer `json:"workflow_question_answers"`
}

type DaQuestionnaireNode struct {
	Id         int                 `json:"da_workflow_node_id"`
	QQuestions []DaWorflowQuestion `json:"workflow_questions"`
}

type DaQuesionnaire struct {
	QNode DaQuestionnaireNode `json:"workflow_node"`
}

type FamilyMember struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type CareAdvisor struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type FamilyAccount struct {
	FamilyMembers []FamilyMember `json:"familyMembers"`
	CareAdvisor   CareAdvisor    `json:"careAdvisor"`
}
