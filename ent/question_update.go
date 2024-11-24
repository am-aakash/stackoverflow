// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"stackoverflow-clone/ent/answer"
	"stackoverflow-clone/ent/comment"
	"stackoverflow-clone/ent/predicate"
	"stackoverflow-clone/ent/question"
	"stackoverflow-clone/ent/tag"
	"stackoverflow-clone/ent/user"
	"stackoverflow-clone/ent/vote"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// QuestionUpdate is the builder for updating Question entities.
type QuestionUpdate struct {
	config
	hooks    []Hook
	mutation *QuestionMutation
}

// Where appends a list predicates to the QuestionUpdate builder.
func (qu *QuestionUpdate) Where(ps ...predicate.Question) *QuestionUpdate {
	qu.mutation.Where(ps...)
	return qu
}

// SetTitle sets the "title" field.
func (qu *QuestionUpdate) SetTitle(s string) *QuestionUpdate {
	qu.mutation.SetTitle(s)
	return qu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (qu *QuestionUpdate) SetNillableTitle(s *string) *QuestionUpdate {
	if s != nil {
		qu.SetTitle(*s)
	}
	return qu
}

// SetContent sets the "content" field.
func (qu *QuestionUpdate) SetContent(s string) *QuestionUpdate {
	qu.mutation.SetContent(s)
	return qu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (qu *QuestionUpdate) SetNillableContent(s *string) *QuestionUpdate {
	if s != nil {
		qu.SetContent(*s)
	}
	return qu
}

// SetCreatedAt sets the "created_at" field.
func (qu *QuestionUpdate) SetCreatedAt(t time.Time) *QuestionUpdate {
	qu.mutation.SetCreatedAt(t)
	return qu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (qu *QuestionUpdate) SetNillableCreatedAt(t *time.Time) *QuestionUpdate {
	if t != nil {
		qu.SetCreatedAt(*t)
	}
	return qu
}

// SetUpdatedAt sets the "updated_at" field.
func (qu *QuestionUpdate) SetUpdatedAt(t time.Time) *QuestionUpdate {
	qu.mutation.SetUpdatedAt(t)
	return qu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (qu *QuestionUpdate) SetNillableUpdatedAt(t *time.Time) *QuestionUpdate {
	if t != nil {
		qu.SetUpdatedAt(*t)
	}
	return qu
}

// SetVoteCount sets the "vote_count" field.
func (qu *QuestionUpdate) SetVoteCount(i int) *QuestionUpdate {
	qu.mutation.ResetVoteCount()
	qu.mutation.SetVoteCount(i)
	return qu
}

// SetNillableVoteCount sets the "vote_count" field if the given value is not nil.
func (qu *QuestionUpdate) SetNillableVoteCount(i *int) *QuestionUpdate {
	if i != nil {
		qu.SetVoteCount(*i)
	}
	return qu
}

// AddVoteCount adds i to the "vote_count" field.
func (qu *QuestionUpdate) AddVoteCount(i int) *QuestionUpdate {
	qu.mutation.AddVoteCount(i)
	return qu
}

// SetViewCount sets the "view_count" field.
func (qu *QuestionUpdate) SetViewCount(i int) *QuestionUpdate {
	qu.mutation.ResetViewCount()
	qu.mutation.SetViewCount(i)
	return qu
}

// SetNillableViewCount sets the "view_count" field if the given value is not nil.
func (qu *QuestionUpdate) SetNillableViewCount(i *int) *QuestionUpdate {
	if i != nil {
		qu.SetViewCount(*i)
	}
	return qu
}

// AddViewCount adds i to the "view_count" field.
func (qu *QuestionUpdate) AddViewCount(i int) *QuestionUpdate {
	qu.mutation.AddViewCount(i)
	return qu
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (qu *QuestionUpdate) SetAuthorID(id int) *QuestionUpdate {
	qu.mutation.SetAuthorID(id)
	return qu
}

// SetNillableAuthorID sets the "author" edge to the User entity by ID if the given value is not nil.
func (qu *QuestionUpdate) SetNillableAuthorID(id *int) *QuestionUpdate {
	if id != nil {
		qu = qu.SetAuthorID(*id)
	}
	return qu
}

// SetAuthor sets the "author" edge to the User entity.
func (qu *QuestionUpdate) SetAuthor(u *User) *QuestionUpdate {
	return qu.SetAuthorID(u.ID)
}

// AddAnswerIDs adds the "answers" edge to the Answer entity by IDs.
func (qu *QuestionUpdate) AddAnswerIDs(ids ...int) *QuestionUpdate {
	qu.mutation.AddAnswerIDs(ids...)
	return qu
}

// AddAnswers adds the "answers" edges to the Answer entity.
func (qu *QuestionUpdate) AddAnswers(a ...*Answer) *QuestionUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return qu.AddAnswerIDs(ids...)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (qu *QuestionUpdate) AddCommentIDs(ids ...int) *QuestionUpdate {
	qu.mutation.AddCommentIDs(ids...)
	return qu
}

// AddComments adds the "comments" edges to the Comment entity.
func (qu *QuestionUpdate) AddComments(c ...*Comment) *QuestionUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return qu.AddCommentIDs(ids...)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (qu *QuestionUpdate) AddTagIDs(ids ...int) *QuestionUpdate {
	qu.mutation.AddTagIDs(ids...)
	return qu
}

// AddTags adds the "tags" edges to the Tag entity.
func (qu *QuestionUpdate) AddTags(t ...*Tag) *QuestionUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return qu.AddTagIDs(ids...)
}

// AddVoteIDs adds the "votes" edge to the Vote entity by IDs.
func (qu *QuestionUpdate) AddVoteIDs(ids ...int) *QuestionUpdate {
	qu.mutation.AddVoteIDs(ids...)
	return qu
}

// AddVotes adds the "votes" edges to the Vote entity.
func (qu *QuestionUpdate) AddVotes(v ...*Vote) *QuestionUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return qu.AddVoteIDs(ids...)
}

// Mutation returns the QuestionMutation object of the builder.
func (qu *QuestionUpdate) Mutation() *QuestionMutation {
	return qu.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (qu *QuestionUpdate) ClearAuthor() *QuestionUpdate {
	qu.mutation.ClearAuthor()
	return qu
}

// ClearAnswers clears all "answers" edges to the Answer entity.
func (qu *QuestionUpdate) ClearAnswers() *QuestionUpdate {
	qu.mutation.ClearAnswers()
	return qu
}

// RemoveAnswerIDs removes the "answers" edge to Answer entities by IDs.
func (qu *QuestionUpdate) RemoveAnswerIDs(ids ...int) *QuestionUpdate {
	qu.mutation.RemoveAnswerIDs(ids...)
	return qu
}

// RemoveAnswers removes "answers" edges to Answer entities.
func (qu *QuestionUpdate) RemoveAnswers(a ...*Answer) *QuestionUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return qu.RemoveAnswerIDs(ids...)
}

// ClearComments clears all "comments" edges to the Comment entity.
func (qu *QuestionUpdate) ClearComments() *QuestionUpdate {
	qu.mutation.ClearComments()
	return qu
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (qu *QuestionUpdate) RemoveCommentIDs(ids ...int) *QuestionUpdate {
	qu.mutation.RemoveCommentIDs(ids...)
	return qu
}

// RemoveComments removes "comments" edges to Comment entities.
func (qu *QuestionUpdate) RemoveComments(c ...*Comment) *QuestionUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return qu.RemoveCommentIDs(ids...)
}

// ClearTags clears all "tags" edges to the Tag entity.
func (qu *QuestionUpdate) ClearTags() *QuestionUpdate {
	qu.mutation.ClearTags()
	return qu
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (qu *QuestionUpdate) RemoveTagIDs(ids ...int) *QuestionUpdate {
	qu.mutation.RemoveTagIDs(ids...)
	return qu
}

// RemoveTags removes "tags" edges to Tag entities.
func (qu *QuestionUpdate) RemoveTags(t ...*Tag) *QuestionUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return qu.RemoveTagIDs(ids...)
}

// ClearVotes clears all "votes" edges to the Vote entity.
func (qu *QuestionUpdate) ClearVotes() *QuestionUpdate {
	qu.mutation.ClearVotes()
	return qu
}

// RemoveVoteIDs removes the "votes" edge to Vote entities by IDs.
func (qu *QuestionUpdate) RemoveVoteIDs(ids ...int) *QuestionUpdate {
	qu.mutation.RemoveVoteIDs(ids...)
	return qu
}

// RemoveVotes removes "votes" edges to Vote entities.
func (qu *QuestionUpdate) RemoveVotes(v ...*Vote) *QuestionUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return qu.RemoveVoteIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (qu *QuestionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, qu.sqlSave, qu.mutation, qu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (qu *QuestionUpdate) SaveX(ctx context.Context) int {
	affected, err := qu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (qu *QuestionUpdate) Exec(ctx context.Context) error {
	_, err := qu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qu *QuestionUpdate) ExecX(ctx context.Context) {
	if err := qu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (qu *QuestionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(question.Table, question.Columns, sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt))
	if ps := qu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := qu.mutation.Title(); ok {
		_spec.SetField(question.FieldTitle, field.TypeString, value)
	}
	if value, ok := qu.mutation.Content(); ok {
		_spec.SetField(question.FieldContent, field.TypeString, value)
	}
	if value, ok := qu.mutation.CreatedAt(); ok {
		_spec.SetField(question.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := qu.mutation.UpdatedAt(); ok {
		_spec.SetField(question.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := qu.mutation.VoteCount(); ok {
		_spec.SetField(question.FieldVoteCount, field.TypeInt, value)
	}
	if value, ok := qu.mutation.AddedVoteCount(); ok {
		_spec.AddField(question.FieldVoteCount, field.TypeInt, value)
	}
	if value, ok := qu.mutation.ViewCount(); ok {
		_spec.SetField(question.FieldViewCount, field.TypeInt, value)
	}
	if value, ok := qu.mutation.AddedViewCount(); ok {
		_spec.AddField(question.FieldViewCount, field.TypeInt, value)
	}
	if qu.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.AuthorTable,
			Columns: []string{question.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.AuthorTable,
			Columns: []string{question.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qu.mutation.AnswersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswersTable,
			Columns: []string{question.AnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.RemovedAnswersIDs(); len(nodes) > 0 && !qu.mutation.AnswersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswersTable,
			Columns: []string{question.AnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.AnswersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswersTable,
			Columns: []string{question.AnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.CommentsTable,
			Columns: question.CommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !qu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.CommentsTable,
			Columns: question.CommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.CommentsTable,
			Columns: question.CommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.TagsTable,
			Columns: question.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.RemovedTagsIDs(); len(nodes) > 0 && !qu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.TagsTable,
			Columns: question.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.TagsTable,
			Columns: question.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qu.mutation.VotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.VotesTable,
			Columns: question.VotesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.RemovedVotesIDs(); len(nodes) > 0 && !qu.mutation.VotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.VotesTable,
			Columns: question.VotesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.VotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.VotesTable,
			Columns: question.VotesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, qu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{question.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	qu.mutation.done = true
	return n, nil
}

// QuestionUpdateOne is the builder for updating a single Question entity.
type QuestionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *QuestionMutation
}

// SetTitle sets the "title" field.
func (quo *QuestionUpdateOne) SetTitle(s string) *QuestionUpdateOne {
	quo.mutation.SetTitle(s)
	return quo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (quo *QuestionUpdateOne) SetNillableTitle(s *string) *QuestionUpdateOne {
	if s != nil {
		quo.SetTitle(*s)
	}
	return quo
}

// SetContent sets the "content" field.
func (quo *QuestionUpdateOne) SetContent(s string) *QuestionUpdateOne {
	quo.mutation.SetContent(s)
	return quo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (quo *QuestionUpdateOne) SetNillableContent(s *string) *QuestionUpdateOne {
	if s != nil {
		quo.SetContent(*s)
	}
	return quo
}

// SetCreatedAt sets the "created_at" field.
func (quo *QuestionUpdateOne) SetCreatedAt(t time.Time) *QuestionUpdateOne {
	quo.mutation.SetCreatedAt(t)
	return quo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (quo *QuestionUpdateOne) SetNillableCreatedAt(t *time.Time) *QuestionUpdateOne {
	if t != nil {
		quo.SetCreatedAt(*t)
	}
	return quo
}

// SetUpdatedAt sets the "updated_at" field.
func (quo *QuestionUpdateOne) SetUpdatedAt(t time.Time) *QuestionUpdateOne {
	quo.mutation.SetUpdatedAt(t)
	return quo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (quo *QuestionUpdateOne) SetNillableUpdatedAt(t *time.Time) *QuestionUpdateOne {
	if t != nil {
		quo.SetUpdatedAt(*t)
	}
	return quo
}

// SetVoteCount sets the "vote_count" field.
func (quo *QuestionUpdateOne) SetVoteCount(i int) *QuestionUpdateOne {
	quo.mutation.ResetVoteCount()
	quo.mutation.SetVoteCount(i)
	return quo
}

// SetNillableVoteCount sets the "vote_count" field if the given value is not nil.
func (quo *QuestionUpdateOne) SetNillableVoteCount(i *int) *QuestionUpdateOne {
	if i != nil {
		quo.SetVoteCount(*i)
	}
	return quo
}

// AddVoteCount adds i to the "vote_count" field.
func (quo *QuestionUpdateOne) AddVoteCount(i int) *QuestionUpdateOne {
	quo.mutation.AddVoteCount(i)
	return quo
}

// SetViewCount sets the "view_count" field.
func (quo *QuestionUpdateOne) SetViewCount(i int) *QuestionUpdateOne {
	quo.mutation.ResetViewCount()
	quo.mutation.SetViewCount(i)
	return quo
}

// SetNillableViewCount sets the "view_count" field if the given value is not nil.
func (quo *QuestionUpdateOne) SetNillableViewCount(i *int) *QuestionUpdateOne {
	if i != nil {
		quo.SetViewCount(*i)
	}
	return quo
}

// AddViewCount adds i to the "view_count" field.
func (quo *QuestionUpdateOne) AddViewCount(i int) *QuestionUpdateOne {
	quo.mutation.AddViewCount(i)
	return quo
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (quo *QuestionUpdateOne) SetAuthorID(id int) *QuestionUpdateOne {
	quo.mutation.SetAuthorID(id)
	return quo
}

// SetNillableAuthorID sets the "author" edge to the User entity by ID if the given value is not nil.
func (quo *QuestionUpdateOne) SetNillableAuthorID(id *int) *QuestionUpdateOne {
	if id != nil {
		quo = quo.SetAuthorID(*id)
	}
	return quo
}

// SetAuthor sets the "author" edge to the User entity.
func (quo *QuestionUpdateOne) SetAuthor(u *User) *QuestionUpdateOne {
	return quo.SetAuthorID(u.ID)
}

// AddAnswerIDs adds the "answers" edge to the Answer entity by IDs.
func (quo *QuestionUpdateOne) AddAnswerIDs(ids ...int) *QuestionUpdateOne {
	quo.mutation.AddAnswerIDs(ids...)
	return quo
}

// AddAnswers adds the "answers" edges to the Answer entity.
func (quo *QuestionUpdateOne) AddAnswers(a ...*Answer) *QuestionUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return quo.AddAnswerIDs(ids...)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (quo *QuestionUpdateOne) AddCommentIDs(ids ...int) *QuestionUpdateOne {
	quo.mutation.AddCommentIDs(ids...)
	return quo
}

// AddComments adds the "comments" edges to the Comment entity.
func (quo *QuestionUpdateOne) AddComments(c ...*Comment) *QuestionUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return quo.AddCommentIDs(ids...)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (quo *QuestionUpdateOne) AddTagIDs(ids ...int) *QuestionUpdateOne {
	quo.mutation.AddTagIDs(ids...)
	return quo
}

// AddTags adds the "tags" edges to the Tag entity.
func (quo *QuestionUpdateOne) AddTags(t ...*Tag) *QuestionUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return quo.AddTagIDs(ids...)
}

// AddVoteIDs adds the "votes" edge to the Vote entity by IDs.
func (quo *QuestionUpdateOne) AddVoteIDs(ids ...int) *QuestionUpdateOne {
	quo.mutation.AddVoteIDs(ids...)
	return quo
}

// AddVotes adds the "votes" edges to the Vote entity.
func (quo *QuestionUpdateOne) AddVotes(v ...*Vote) *QuestionUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return quo.AddVoteIDs(ids...)
}

// Mutation returns the QuestionMutation object of the builder.
func (quo *QuestionUpdateOne) Mutation() *QuestionMutation {
	return quo.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (quo *QuestionUpdateOne) ClearAuthor() *QuestionUpdateOne {
	quo.mutation.ClearAuthor()
	return quo
}

// ClearAnswers clears all "answers" edges to the Answer entity.
func (quo *QuestionUpdateOne) ClearAnswers() *QuestionUpdateOne {
	quo.mutation.ClearAnswers()
	return quo
}

// RemoveAnswerIDs removes the "answers" edge to Answer entities by IDs.
func (quo *QuestionUpdateOne) RemoveAnswerIDs(ids ...int) *QuestionUpdateOne {
	quo.mutation.RemoveAnswerIDs(ids...)
	return quo
}

// RemoveAnswers removes "answers" edges to Answer entities.
func (quo *QuestionUpdateOne) RemoveAnswers(a ...*Answer) *QuestionUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return quo.RemoveAnswerIDs(ids...)
}

// ClearComments clears all "comments" edges to the Comment entity.
func (quo *QuestionUpdateOne) ClearComments() *QuestionUpdateOne {
	quo.mutation.ClearComments()
	return quo
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (quo *QuestionUpdateOne) RemoveCommentIDs(ids ...int) *QuestionUpdateOne {
	quo.mutation.RemoveCommentIDs(ids...)
	return quo
}

// RemoveComments removes "comments" edges to Comment entities.
func (quo *QuestionUpdateOne) RemoveComments(c ...*Comment) *QuestionUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return quo.RemoveCommentIDs(ids...)
}

// ClearTags clears all "tags" edges to the Tag entity.
func (quo *QuestionUpdateOne) ClearTags() *QuestionUpdateOne {
	quo.mutation.ClearTags()
	return quo
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (quo *QuestionUpdateOne) RemoveTagIDs(ids ...int) *QuestionUpdateOne {
	quo.mutation.RemoveTagIDs(ids...)
	return quo
}

// RemoveTags removes "tags" edges to Tag entities.
func (quo *QuestionUpdateOne) RemoveTags(t ...*Tag) *QuestionUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return quo.RemoveTagIDs(ids...)
}

// ClearVotes clears all "votes" edges to the Vote entity.
func (quo *QuestionUpdateOne) ClearVotes() *QuestionUpdateOne {
	quo.mutation.ClearVotes()
	return quo
}

// RemoveVoteIDs removes the "votes" edge to Vote entities by IDs.
func (quo *QuestionUpdateOne) RemoveVoteIDs(ids ...int) *QuestionUpdateOne {
	quo.mutation.RemoveVoteIDs(ids...)
	return quo
}

// RemoveVotes removes "votes" edges to Vote entities.
func (quo *QuestionUpdateOne) RemoveVotes(v ...*Vote) *QuestionUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return quo.RemoveVoteIDs(ids...)
}

// Where appends a list predicates to the QuestionUpdate builder.
func (quo *QuestionUpdateOne) Where(ps ...predicate.Question) *QuestionUpdateOne {
	quo.mutation.Where(ps...)
	return quo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (quo *QuestionUpdateOne) Select(field string, fields ...string) *QuestionUpdateOne {
	quo.fields = append([]string{field}, fields...)
	return quo
}

// Save executes the query and returns the updated Question entity.
func (quo *QuestionUpdateOne) Save(ctx context.Context) (*Question, error) {
	return withHooks(ctx, quo.sqlSave, quo.mutation, quo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (quo *QuestionUpdateOne) SaveX(ctx context.Context) *Question {
	node, err := quo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (quo *QuestionUpdateOne) Exec(ctx context.Context) error {
	_, err := quo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (quo *QuestionUpdateOne) ExecX(ctx context.Context) {
	if err := quo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (quo *QuestionUpdateOne) sqlSave(ctx context.Context) (_node *Question, err error) {
	_spec := sqlgraph.NewUpdateSpec(question.Table, question.Columns, sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt))
	id, ok := quo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Question.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := quo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, question.FieldID)
		for _, f := range fields {
			if !question.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != question.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := quo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := quo.mutation.Title(); ok {
		_spec.SetField(question.FieldTitle, field.TypeString, value)
	}
	if value, ok := quo.mutation.Content(); ok {
		_spec.SetField(question.FieldContent, field.TypeString, value)
	}
	if value, ok := quo.mutation.CreatedAt(); ok {
		_spec.SetField(question.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := quo.mutation.UpdatedAt(); ok {
		_spec.SetField(question.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := quo.mutation.VoteCount(); ok {
		_spec.SetField(question.FieldVoteCount, field.TypeInt, value)
	}
	if value, ok := quo.mutation.AddedVoteCount(); ok {
		_spec.AddField(question.FieldVoteCount, field.TypeInt, value)
	}
	if value, ok := quo.mutation.ViewCount(); ok {
		_spec.SetField(question.FieldViewCount, field.TypeInt, value)
	}
	if value, ok := quo.mutation.AddedViewCount(); ok {
		_spec.AddField(question.FieldViewCount, field.TypeInt, value)
	}
	if quo.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.AuthorTable,
			Columns: []string{question.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.AuthorTable,
			Columns: []string{question.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if quo.mutation.AnswersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswersTable,
			Columns: []string{question.AnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.RemovedAnswersIDs(); len(nodes) > 0 && !quo.mutation.AnswersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswersTable,
			Columns: []string{question.AnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.AnswersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswersTable,
			Columns: []string{question.AnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if quo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.CommentsTable,
			Columns: question.CommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !quo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.CommentsTable,
			Columns: question.CommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.CommentsTable,
			Columns: question.CommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if quo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.TagsTable,
			Columns: question.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !quo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.TagsTable,
			Columns: question.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.TagsTable,
			Columns: question.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if quo.mutation.VotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.VotesTable,
			Columns: question.VotesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.RemovedVotesIDs(); len(nodes) > 0 && !quo.mutation.VotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.VotesTable,
			Columns: question.VotesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.VotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.VotesTable,
			Columns: question.VotesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Question{config: quo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, quo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{question.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	quo.mutation.done = true
	return _node, nil
}