package question

import (
	"context"
	"fmt"

	"online-test/config"
)

type Repository struct{}

// =========================
// INSERT QUESTION
// =========================
func (r *Repository) InsertQuestion(
	qType string,
	text string,
	explanation string,
	meta string,
	userID string,
) (string, error) {

	var id string

	err := config.DB.QueryRow(
		context.Background(),
		`INSERT INTO question_bank
			(question_type, question_text, explanation, metadata, created_by)
		VALUES ($1,$2,$3,$4,$5)
		RETURNING id`,
		qType, text, explanation, meta, userID,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

// =========================
// INSERT OPTIONS
// =========================
func (r *Repository) InsertOptions(questionID string, options []OptionRequest) error {

	for _, opt := range options {

		_, err := config.DB.Exec(
			context.Background(),
			`INSERT INTO question_options
				(question_id, option_text, is_correct)
			VALUES ($1,$2,$3)`,
			questionID,
			opt.OptionText,
			opt.IsCorrect,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

// =========================
// GET ALL QUESTIONS
// =========================
func (r *Repository) GetAll() ([]map[string]interface{}, error) {

	rows, err := config.DB.Query(
		context.Background(),
		`SELECT id, question_type, question_text, metadata
		 FROM question_bank
		 ORDER BY created_at DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []map[string]interface{}

	for rows.Next() {

		var id, qtype, text, meta string

		if err := rows.Scan(&id, &qtype, &text, &meta); err != nil {
			return nil, err
		}

		result = append(result, map[string]interface{}{
			"id":            id,
			"type":          qtype,
			"question_text": text,
			"metadata":      meta,
		})
	}

	return result, nil
}

// =========================
// GET BY ID (WITH OPTIONS)
// =========================
func (r *Repository) GetByID(id string) (map[string]interface{}, error) {

	var qID, qType, text, meta string

	err := config.DB.QueryRow(
		context.Background(),
		`SELECT id, question_type, question_text, metadata
		 FROM question_bank
		 WHERE id=$1`,
		id,
	).Scan(&qID, &qType, &text, &meta)

	if err != nil {
		return nil, err
	}

	rows, err := config.DB.Query(
		context.Background(),
		`SELECT option_text, is_correct
		 FROM question_options
		 WHERE question_id=$1`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var options []map[string]interface{}

	for rows.Next() {

		var optText string
		var correct bool

		if err := rows.Scan(&optText, &correct); err != nil {
			return nil, err
		}

		options = append(options, map[string]interface{}{
			"text":       optText,
			"is_correct": correct,
		})
	}

	question := map[string]interface{}{
		"id":       qID,
		"type":     qType,
		"text":     text,
		"metadata": meta,
		"options":  options,
	}

	return question, nil
}

// =========================
// DELETE QUESTION
// =========================
func (r *Repository) Delete(id string) error {

	res, err := config.DB.Exec(
		context.Background(),
		`DELETE FROM question_bank WHERE id=$1`,
		id,
	)
	if err != nil {
		return err
	}

	affected := res.RowsAffected()
	if affected == 0 {
		return fmt.Errorf("question not found")
	}

	return nil
}
