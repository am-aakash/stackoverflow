// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"stackoverflow-clone/ent/answer"
	"stackoverflow-clone/ent/predicate"
	"stackoverflow-clone/ent/question"
	"stackoverflow-clone/ent/user"
	"stackoverflow-clone/ent/vote"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VoteUpdate is the builder for updating Vote entities.
type VoteUpdate struct {
	config
	hooks    []Hook
	mutation *VoteMutation
}

// Where appends a list predicates to the VoteUpdate builder.
func (vu *VoteUpdate) Where(ps ...predicate.Vote) *VoteUpdate {
	vu.mutation.Where(ps...)
	return vu
}

// SetValue sets the "value" field.
func (vu *VoteUpdate) SetValue(i int) *VoteUpdate {
	vu.mutation.ResetValue()
	vu.mutation.SetValue(i)
	return vu
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (vu *VoteUpdate) SetNillableValue(i *int) *VoteUpdate {
	if i != nil {
		vu.SetValue(*i)
	}
	return vu
}

// AddValue adds i to the "value" field.
func (vu *VoteUpdate) AddValue(i int) *VoteUpdate {
	vu.mutation.AddValue(i)
	return vu
}

// SetCreatedAt sets the "created_at" field.
func (vu *VoteUpdate) SetCreatedAt(t time.Time) *VoteUpdate {
	vu.mutation.SetCreatedAt(t)
	return vu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vu *VoteUpdate) SetNillableCreatedAt(t *time.Time) *VoteUpdate {
	if t != nil {
		vu.SetCreatedAt(*t)
	}
	return vu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (vu *VoteUpdate) SetUserID(id int) *VoteUpdate {
	vu.mutation.SetUserID(id)
	return vu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (vu *VoteUpdate) SetNillableUserID(id *int) *VoteUpdate {
	if id != nil {
		vu = vu.SetUserID(*id)
	}
	return vu
}

// SetUser sets the "user" edge to the User entity.
func (vu *VoteUpdate) SetUser(u *User) *VoteUpdate {
	return vu.SetUserID(u.ID)
}

// AddQuestionIDs adds the "question" edge to the Question entity by IDs.
func (vu *VoteUpdate) AddQuestionIDs(ids ...int) *VoteUpdate {
	vu.mutation.AddQuestionIDs(ids...)
	return vu
}

// AddQuestion adds the "question" edges to the Question entity.
func (vu *VoteUpdate) AddQuestion(q ...*Question) *VoteUpdate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return vu.AddQuestionIDs(ids...)
}

// AddAnswerIDs adds the "answer" edge to the Answer entity by IDs.
func (vu *VoteUpdate) AddAnswerIDs(ids ...int) *VoteUpdate {
	vu.mutation.AddAnswerIDs(ids...)
	return vu
}

// AddAnswer adds the "answer" edges to the Answer entity.
func (vu *VoteUpdate) AddAnswer(a ...*Answer) *VoteUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return vu.AddAnswerIDs(ids...)
}

// Mutation returns the VoteMutation object of the builder.
func (vu *VoteUpdate) Mutation() *VoteMutation {
	return vu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (vu *VoteUpdate) ClearUser() *VoteUpdate {
	vu.mutation.ClearUser()
	return vu
}

// ClearQuestion clears all "question" edges to the Question entity.
func (vu *VoteUpdate) ClearQuestion() *VoteUpdate {
	vu.mutation.ClearQuestion()
	return vu
}

// RemoveQuestionIDs removes the "question" edge to Question entities by IDs.
func (vu *VoteUpdate) RemoveQuestionIDs(ids ...int) *VoteUpdate {
	vu.mutation.RemoveQuestionIDs(ids...)
	return vu
}

// RemoveQuestion removes "question" edges to Question entities.
func (vu *VoteUpdate) RemoveQuestion(q ...*Question) *VoteUpdate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return vu.RemoveQuestionIDs(ids...)
}

// ClearAnswer clears all "answer" edges to the Answer entity.
func (vu *VoteUpdate) ClearAnswer() *VoteUpdate {
	vu.mutation.ClearAnswer()
	return vu
}

// RemoveAnswerIDs removes the "answer" edge to Answer entities by IDs.
func (vu *VoteUpdate) RemoveAnswerIDs(ids ...int) *VoteUpdate {
	vu.mutation.RemoveAnswerIDs(ids...)
	return vu
}

// RemoveAnswer removes "answer" edges to Answer entities.
func (vu *VoteUpdate) RemoveAnswer(a ...*Answer) *VoteUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return vu.RemoveAnswerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *VoteUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, vu.sqlSave, vu.mutation, vu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vu *VoteUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VoteUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VoteUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vu *VoteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(vote.Table, vote.Columns, sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt))
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.Value(); ok {
		_spec.SetField(vote.FieldValue, field.TypeInt, value)
	}
	if value, ok := vu.mutation.AddedValue(); ok {
		_spec.AddField(vote.FieldValue, field.TypeInt, value)
	}
	if value, ok := vu.mutation.CreatedAt(); ok {
		_spec.SetField(vote.FieldCreatedAt, field.TypeTime, value)
	}
	if vu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vote.UserTable,
			Columns: []string{vote.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vote.UserTable,
			Columns: []string{vote.UserColumn},
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
	if vu.mutation.QuestionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vote.QuestionTable,
			Columns: vote.QuestionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.RemovedQuestionIDs(); len(nodes) > 0 && !vu.mutation.QuestionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vote.QuestionTable,
			Columns: vote.QuestionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vote.QuestionTable,
			Columns: vote.QuestionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vu.mutation.AnswerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vote.AnswerTable,
			Columns: vote.AnswerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.RemovedAnswerIDs(); len(nodes) > 0 && !vu.mutation.AnswerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vote.AnswerTable,
			Columns: vote.AnswerPrimaryKey,
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
	if nodes := vu.mutation.AnswerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vote.AnswerTable,
			Columns: vote.AnswerPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vote.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vu.mutation.done = true
	return n, nil
}

// VoteUpdateOne is the builder for updating a single Vote entity.
type VoteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VoteMutation
}

// SetValue sets the "value" field.
func (vuo *VoteUpdateOne) SetValue(i int) *VoteUpdateOne {
	vuo.mutation.ResetValue()
	vuo.mutation.SetValue(i)
	return vuo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (vuo *VoteUpdateOne) SetNillableValue(i *int) *VoteUpdateOne {
	if i != nil {
		vuo.SetValue(*i)
	}
	return vuo
}

// AddValue adds i to the "value" field.
func (vuo *VoteUpdateOne) AddValue(i int) *VoteUpdateOne {
	vuo.mutation.AddValue(i)
	return vuo
}

// SetCreatedAt sets the "created_at" field.
func (vuo *VoteUpdateOne) SetCreatedAt(t time.Time) *VoteUpdateOne {
	vuo.mutation.SetCreatedAt(t)
	return vuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vuo *VoteUpdateOne) SetNillableCreatedAt(t *time.Time) *VoteUpdateOne {
	if t != nil {
		vuo.SetCreatedAt(*t)
	}
	return vuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (vuo *VoteUpdateOne) SetUserID(id int) *VoteUpdateOne {
	vuo.mutation.SetUserID(id)
	return vuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (vuo *VoteUpdateOne) SetNillableUserID(id *int) *VoteUpdateOne {
	if id != nil {
		vuo = vuo.SetUserID(*id)
	}
	return vuo
}

// SetUser sets the "user" edge to the User entity.
func (vuo *VoteUpdateOne) SetUser(u *User) *VoteUpdateOne {
	return vuo.SetUserID(u.ID)
}

// AddQuestionIDs adds the "question" edge to the Question entity by IDs.
func (vuo *VoteUpdateOne) AddQuestionIDs(ids ...int) *VoteUpdateOne {
	vuo.mutation.AddQuestionIDs(ids...)
	return vuo
}

// AddQuestion adds the "question" edges to the Question entity.
func (vuo *VoteUpdateOne) AddQuestion(q ...*Question) *VoteUpdateOne {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return vuo.AddQuestionIDs(ids...)
}

// AddAnswerIDs adds the "answer" edge to the Answer entity by IDs.
func (vuo *VoteUpdateOne) AddAnswerIDs(ids ...int) *VoteUpdateOne {
	vuo.mutation.AddAnswerIDs(ids...)
	return vuo
}

// AddAnswer adds the "answer" edges to the Answer entity.
func (vuo *VoteUpdateOne) AddAnswer(a ...*Answer) *VoteUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return vuo.AddAnswerIDs(ids...)
}

// Mutation returns the VoteMutation object of the builder.
func (vuo *VoteUpdateOne) Mutation() *VoteMutation {
	return vuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (vuo *VoteUpdateOne) ClearUser() *VoteUpdateOne {
	vuo.mutation.ClearUser()
	return vuo
}

// ClearQuestion clears all "question" edges to the Question entity.
func (vuo *VoteUpdateOne) ClearQuestion() *VoteUpdateOne {
	vuo.mutation.ClearQuestion()
	return vuo
}

// RemoveQuestionIDs removes the "question" edge to Question entities by IDs.
func (vuo *VoteUpdateOne) RemoveQuestionIDs(ids ...int) *VoteUpdateOne {
	vuo.mutation.RemoveQuestionIDs(ids...)
	return vuo
}

// RemoveQuestion removes "question" edges to Question entities.
func (vuo *VoteUpdateOne) RemoveQuestion(q ...*Question) *VoteUpdateOne {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return vuo.RemoveQuestionIDs(ids...)
}

// ClearAnswer clears all "answer" edges to the Answer entity.
func (vuo *VoteUpdateOne) ClearAnswer() *VoteUpdateOne {
	vuo.mutation.ClearAnswer()
	return vuo
}

// RemoveAnswerIDs removes the "answer" edge to Answer entities by IDs.
func (vuo *VoteUpdateOne) RemoveAnswerIDs(ids ...int) *VoteUpdateOne {
	vuo.mutation.RemoveAnswerIDs(ids...)
	return vuo
}

// RemoveAnswer removes "answer" edges to Answer entities.
func (vuo *VoteUpdateOne) RemoveAnswer(a ...*Answer) *VoteUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return vuo.RemoveAnswerIDs(ids...)
}

// Where appends a list predicates to the VoteUpdate builder.
func (vuo *VoteUpdateOne) Where(ps ...predicate.Vote) *VoteUpdateOne {
	vuo.mutation.Where(ps...)
	return vuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vuo *VoteUpdateOne) Select(field string, fields ...string) *VoteUpdateOne {
	vuo.fields = append([]string{field}, fields...)
	return vuo
}

// Save executes the query and returns the updated Vote entity.
func (vuo *VoteUpdateOne) Save(ctx context.Context) (*Vote, error) {
	return withHooks(ctx, vuo.sqlSave, vuo.mutation, vuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *VoteUpdateOne) SaveX(ctx context.Context) *Vote {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VoteUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VoteUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vuo *VoteUpdateOne) sqlSave(ctx context.Context) (_node *Vote, err error) {
	_spec := sqlgraph.NewUpdateSpec(vote.Table, vote.Columns, sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt))
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Vote.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vote.FieldID)
		for _, f := range fields {
			if !vote.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != vote.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vuo.mutation.Value(); ok {
		_spec.SetField(vote.FieldValue, field.TypeInt, value)
	}
	if value, ok := vuo.mutation.AddedValue(); ok {
		_spec.AddField(vote.FieldValue, field.TypeInt, value)
	}
	if value, ok := vuo.mutation.CreatedAt(); ok {
		_spec.SetField(vote.FieldCreatedAt, field.TypeTime, value)
	}
	if vuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vote.UserTable,
			Columns: []string{vote.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vote.UserTable,
			Columns: []string{vote.UserColumn},
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
	if vuo.mutation.QuestionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vote.QuestionTable,
			Columns: vote.QuestionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.RemovedQuestionIDs(); len(nodes) > 0 && !vuo.mutation.QuestionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vote.QuestionTable,
			Columns: vote.QuestionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vote.QuestionTable,
			Columns: vote.QuestionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vuo.mutation.AnswerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vote.AnswerTable,
			Columns: vote.AnswerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.RemovedAnswerIDs(); len(nodes) > 0 && !vuo.mutation.AnswerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vote.AnswerTable,
			Columns: vote.AnswerPrimaryKey,
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
	if nodes := vuo.mutation.AnswerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vote.AnswerTable,
			Columns: vote.AnswerPrimaryKey,
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
	_node = &Vote{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vote.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vuo.mutation.done = true
	return _node, nil
}