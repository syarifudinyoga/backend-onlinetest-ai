package question

import (
	"encoding/json"
)

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) CreateQuestion(req CreateQuestionRequest, userID string) (string, error) {

	meta, _ := json.Marshal(req.Metadata)

	return s.Repo.InsertQuestion(
		req.QuestionType,
		req.QuestionText,
		req.Explanation,
		string(meta),
		userID,
	)
}

func (s *Service) AddOptions(questionID string, req AddOptionsRequest) error {
	return s.Repo.InsertOptions(questionID, req.Options)
}

func (s *Service) ListQuestions() ([]map[string]interface{}, error) {
	return s.Repo.GetAll()
}

func (s *Service) GetDetail(id string) (map[string]interface{}, error) {
	return s.Repo.GetByID(id)
}

func (s *Service) Delete(id string) error {
	return s.Repo.Delete(id)
}
