package models

type Question struct {
	ID          string
	Type        string
	Text        string
	Explanation string
	Metadata    string // JSON string
}

type QuestionOption struct {
	ID         string
	QuestionID string
	Text       string
	IsCorrect  bool
}
