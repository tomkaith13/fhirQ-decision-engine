package questionnaire_resp_model

type QRBody struct {
	QuestionnaireResponse any      `json:"questionnaire_resp,omitempty"`
	QuestionnaireIds      []string `json:"q_ids,omitempty"`
}
