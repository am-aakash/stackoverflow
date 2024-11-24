package services

import (
	"context"
	"errors"
	"stackoverflow-clone/ent"
	"stackoverflow-clone/ent/question"
	"stackoverflow-clone/ent/user"
	"stackoverflow-clone/ent/vote"
	"time"
)

type VoteService struct {
	client *ent.Client
}

func NewVoteService(client *ent.Client) *VoteService {
	return &VoteService{client: client}
}

func (s *VoteService) VoteQuestion(ctx context.Context, userID, questionID int, value int) error {
	tx, err := s.client.Tx(ctx)
	if err != nil {
			return err
	}
	defer tx.Rollback()

	// Check if user has already voted
	existingVote, err := tx.Vote.Query().
			Where(
					vote.HasUserWith(user.ID(userID)),
					vote.HasQuestionWith(question.ID(questionID)),
			).Only(ctx)

	if err != nil && !ent.IsNotFound(err) {
			return err
	}

	if existingVote != nil {
			if existingVote.Value == value {
					return errors.New("already voted")
			}

			// Update existing vote
			_, err = existingVote.Update().
					SetValue(value).
					Save(ctx)
	} else {
			// Create new vote
			_, err = tx.Vote.Create().
					SetUserID(userID).
					// SetQuestionID(questionID).
					SetValue(value).
					SetCreatedAt(time.Now()).
					Save(ctx)
	}

	if err != nil {
			return err
	}

	// Update question vote count
	count, err := tx.Vote.Query().
			Where(vote.HasQuestionWith(question.ID(questionID))).
			Aggregate(ent.Sum(vote.FieldValue)).
			Int(ctx)

	if err != nil {
			return err
	}

	_, err = tx.Question.UpdateOneID(questionID).
			SetVoteCount(count).
			Save(ctx)

	if err != nil {
			return err
	}

	return tx.Commit()
}