package services

import (
	"context"
	"fmt"
	"stackoverflow-clone/ent"
	"stackoverflow-clone/ent/question"
	"stackoverflow-clone/ent/tag"
	"time"
)

type QuestionService struct {
	client *ent.Client
}

func NewQuestionService(client *ent.Client) *QuestionService {
	return &QuestionService{client: client}
}

func (s *QuestionService) CreateQuestion(ctx context.Context, authorID int, title, content string, tags []string) (*ent.Question, error) {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	question, err := tx.Question.Create().
		SetTitle(title).
		SetContent(content).
		SetAuthorID(authorID).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	for _, tagName := range tags {
		tag, err := tx.Tag.Query().
			Where(tag.NameEQ(tagName)).
			Only(ctx)

		if ent.IsNotFound(err) {
			// If the tag does not exist, create it with a default description
			tag, err = tx.Tag.Create().
				SetName(tagName).
				SetDescription(fmt.Sprintf("Tag for %s", tagName)). // Default description
				Save(ctx)
			if err != nil {
				return nil, err
			}
		} else if err != nil {
			return nil, err
		}

		// Associate the tag with the question
		_, err = question.Update().
			AddTags(tag).
			Save(ctx)

		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return question, nil
}

func (s *QuestionService) SearchQuestions(ctx context.Context, query string, tags []string) ([]*ent.Question, error) {
	questionQuery := s.client.Question.Query()

	if query != "" {
		questionQuery = questionQuery.Where(
			question.Or(
				question.TitleContainsFold(query),   // Case-insensitive match for the title
				question.ContentContainsFold(query), // Case-insensitive match for the content
			),
		)
	}

	if len(tags) > 0 {
		questionQuery = questionQuery.Where(
			question.HasTagsWith(tag.NameIn(tags...)),
		)
	}

	return questionQuery.Order(ent.Desc(question.FieldCreatedAt)).All(ctx)
}

func (s *QuestionService) GetQuestionWithAnswers(ctx context.Context, questionID int) (map[string]interface{}, error) {
	// Query the question along with its answers
	q, err := s.client.Question.Query().
		Where(question.IDEQ(questionID)).
		WithAnswers(func(aq *ent.AnswerQuery) {
			aq.Order(ent.Desc("created_at")) // Order answers by creation date
		}).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	// Prepare the response
	response := map[string]interface{}{
		"id":       q.ID,
		"title":    q.Title,
		"content":  q.Content,
		"created":  q.CreatedAt,
		"answers":  []map[string]interface{}{},
		"accepted": []map[string]interface{}{},
	}

	for _, answer := range q.Edges.Answers {
		answerDetails := map[string]interface{}{
			"id":          answer.ID,
			"content":     answer.Content,
			"created":     answer.CreatedAt,
			"is_accepted": answer.IsAccepted,
		}

		if answer.IsAccepted {
			response["accepted"] = append(response["accepted"].([]map[string]interface{}), answerDetails)
		} else {
			response["answers"] = append(response["answers"].([]map[string]interface{}), answerDetails)
		}
	}

	return response, nil
}
