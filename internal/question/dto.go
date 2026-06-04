package question

type CreateQuestionRequest struct {
	QuestionType string                 `json:"question_type" binding:"required"`
	QuestionText string                 `json:"question_text" binding:"required"`
	Explanation  string                 `json:"explanation"`
	Metadata     map[string]interface{} `json:"metadata"`
}

type AddOptionsRequest struct {
	Options []OptionRequest `json:"options" binding:"required"`
}

type OptionRequest struct {
	OptionText string `json:"option_text" binding:"required"`
	IsCorrect  bool   `json:"is_correct"`
}
