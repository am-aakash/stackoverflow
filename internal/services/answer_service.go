package services

import (
	"context"
	"errors"
	"stackoverflow-clone/ent"
	"stackoverflow-clone/ent/answer"
	"time"
)

type AnswerService struct {
	client *ent.Client
}

func NewAnswerService(client *ent.Client) *AnswerService {
	return &AnswerService{client: client}
}

func (s *AnswerService) CreateAnswer(ctx context.Context, userID, questionID int, content string) (*ent.Answer, error) {
	return s.client.Answer.Create().
		SetContent(content).
		SetAuthorID(userID).
		SetQuestionID(questionID).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
}

func (s *AnswerService) AcceptAnswer(ctx context.Context, userID, answerID int) error {
	// Fetch the answer along with its associated question and the question's author
	ans, err := s.client.Answer.Query().
		Where(answer.IDEQ(answerID)).
		WithQuestion(func(q *ent.QuestionQuery) {
			q.WithAuthor() // Load the nested Author edge
		}).
		Only(ctx)
	if err != nil {
		return errors.New("answer not found")
	}

	if ans.Edges.Question == nil || ans.Edges.Question.Edges.Author == nil {
		return errors.New("associated question or author not found")
	}

	// Check if the user is the creator of the question
	if ans.Edges.Question.Edges.Author.ID != userID {
		return errors.New("you are not the creator of this question")
	}

	// Update the answer to mark it as accepted
	_, err = s.client.Answer.UpdateOneID(answerID).
		SetIsAccepted(true).
		Save(ctx)
	if err != nil {
		return errors.New("failed to accept answer")
	}

	return nil
}
