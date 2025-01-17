// Code generated by ent, DO NOT EDIT.

package answer

import (
	"stackoverflow-clone/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Answer {
	return predicate.Answer(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Answer {
	return predicate.Answer(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Answer {
	return predicate.Answer(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Answer {
	return predicate.Answer(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Answer {
	return predicate.Answer(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Answer {
	return predicate.Answer(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Answer {
	return predicate.Answer(sql.FieldLTE(FieldID, id))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldContent, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldUpdatedAt, v))
}

// IsAccepted applies equality check predicate on the "is_accepted" field. It's identical to IsAcceptedEQ.
func IsAccepted(v bool) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldIsAccepted, v))
}

// VoteCount applies equality check predicate on the "vote_count" field. It's identical to VoteCountEQ.
func VoteCount(v int) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldVoteCount, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Answer {
	return predicate.Answer(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Answer {
	return predicate.Answer(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Answer {
	return predicate.Answer(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Answer {
	return predicate.Answer(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Answer {
	return predicate.Answer(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Answer {
	return predicate.Answer(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Answer {
	return predicate.Answer(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Answer {
	return predicate.Answer(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Answer {
	return predicate.Answer(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Answer {
	return predicate.Answer(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Answer {
	return predicate.Answer(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Answer {
	return predicate.Answer(sql.FieldContainsFold(FieldContent, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldLTE(FieldUpdatedAt, v))
}

// IsAcceptedEQ applies the EQ predicate on the "is_accepted" field.
func IsAcceptedEQ(v bool) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldIsAccepted, v))
}

// IsAcceptedNEQ applies the NEQ predicate on the "is_accepted" field.
func IsAcceptedNEQ(v bool) predicate.Answer {
	return predicate.Answer(sql.FieldNEQ(FieldIsAccepted, v))
}

// VoteCountEQ applies the EQ predicate on the "vote_count" field.
func VoteCountEQ(v int) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldVoteCount, v))
}

// VoteCountNEQ applies the NEQ predicate on the "vote_count" field.
func VoteCountNEQ(v int) predicate.Answer {
	return predicate.Answer(sql.FieldNEQ(FieldVoteCount, v))
}

// VoteCountIn applies the In predicate on the "vote_count" field.
func VoteCountIn(vs ...int) predicate.Answer {
	return predicate.Answer(sql.FieldIn(FieldVoteCount, vs...))
}

// VoteCountNotIn applies the NotIn predicate on the "vote_count" field.
func VoteCountNotIn(vs ...int) predicate.Answer {
	return predicate.Answer(sql.FieldNotIn(FieldVoteCount, vs...))
}

// VoteCountGT applies the GT predicate on the "vote_count" field.
func VoteCountGT(v int) predicate.Answer {
	return predicate.Answer(sql.FieldGT(FieldVoteCount, v))
}

// VoteCountGTE applies the GTE predicate on the "vote_count" field.
func VoteCountGTE(v int) predicate.Answer {
	return predicate.Answer(sql.FieldGTE(FieldVoteCount, v))
}

// VoteCountLT applies the LT predicate on the "vote_count" field.
func VoteCountLT(v int) predicate.Answer {
	return predicate.Answer(sql.FieldLT(FieldVoteCount, v))
}

// VoteCountLTE applies the LTE predicate on the "vote_count" field.
func VoteCountLTE(v int) predicate.Answer {
	return predicate.Answer(sql.FieldLTE(FieldVoteCount, v))
}

// HasAuthor applies the HasEdge predicate on the "author" edge.
func HasAuthor() predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorWith applies the HasEdge predicate on the "author" edge with a given conditions (other predicates).
func HasAuthorWith(preds ...predicate.User) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		step := newAuthorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasQuestion applies the HasEdge predicate on the "question" edge.
func HasQuestion() predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, QuestionTable, QuestionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasQuestionWith applies the HasEdge predicate on the "question" edge with a given conditions (other predicates).
func HasQuestionWith(preds ...predicate.Question) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		step := newQuestionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasComments applies the HasEdge predicate on the "comments" edge.
func HasComments() predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, CommentsTable, CommentsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentsWith applies the HasEdge predicate on the "comments" edge with a given conditions (other predicates).
func HasCommentsWith(preds ...predicate.Comment) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		step := newCommentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVotes applies the HasEdge predicate on the "votes" edge.
func HasVotes() predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, VotesTable, VotesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVotesWith applies the HasEdge predicate on the "votes" edge with a given conditions (other predicates).
func HasVotesWith(preds ...predicate.Vote) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		step := newVotesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Answer) predicate.Answer {
	return predicate.Answer(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Answer) predicate.Answer {
	return predicate.Answer(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Answer) predicate.Answer {
	return predicate.Answer(sql.NotPredicates(p))
}
