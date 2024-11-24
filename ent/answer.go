// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"stackoverflow-clone/ent/answer"
	"stackoverflow-clone/ent/question"
	"stackoverflow-clone/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Answer is the model entity for the Answer schema.
type Answer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// IsAccepted holds the value of the "is_accepted" field.
	IsAccepted bool `json:"is_accepted,omitempty"`
	// VoteCount holds the value of the "vote_count" field.
	VoteCount int `json:"vote_count,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AnswerQuery when eager-loading is set.
	Edges            AnswerEdges `json:"edges"`
	question_answers *int
	user_answers     *int
	selectValues     sql.SelectValues
}

// AnswerEdges holds the relations/edges for other nodes in the graph.
type AnswerEdges struct {
	// Author holds the value of the author edge.
	Author *User `json:"author,omitempty"`
	// Question holds the value of the question edge.
	Question *Question `json:"question,omitempty"`
	// Comments holds the value of the comments edge.
	Comments []*Comment `json:"comments,omitempty"`
	// Votes holds the value of the votes edge.
	Votes []*Vote `json:"votes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AnswerEdges) AuthorOrErr() (*User, error) {
	if e.Author != nil {
		return e.Author, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "author"}
}

// QuestionOrErr returns the Question value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AnswerEdges) QuestionOrErr() (*Question, error) {
	if e.Question != nil {
		return e.Question, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: question.Label}
	}
	return nil, &NotLoadedError{edge: "question"}
}

// CommentsOrErr returns the Comments value or an error if the edge
// was not loaded in eager-loading.
func (e AnswerEdges) CommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[2] {
		return e.Comments, nil
	}
	return nil, &NotLoadedError{edge: "comments"}
}

// VotesOrErr returns the Votes value or an error if the edge
// was not loaded in eager-loading.
func (e AnswerEdges) VotesOrErr() ([]*Vote, error) {
	if e.loadedTypes[3] {
		return e.Votes, nil
	}
	return nil, &NotLoadedError{edge: "votes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Answer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case answer.FieldIsAccepted:
			values[i] = new(sql.NullBool)
		case answer.FieldID, answer.FieldVoteCount:
			values[i] = new(sql.NullInt64)
		case answer.FieldContent:
			values[i] = new(sql.NullString)
		case answer.FieldCreatedAt, answer.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case answer.ForeignKeys[0]: // question_answers
			values[i] = new(sql.NullInt64)
		case answer.ForeignKeys[1]: // user_answers
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Answer fields.
func (a *Answer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case answer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case answer.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				a.Content = value.String
			}
		case answer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case answer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case answer.FieldIsAccepted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_accepted", values[i])
			} else if value.Valid {
				a.IsAccepted = value.Bool
			}
		case answer.FieldVoteCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field vote_count", values[i])
			} else if value.Valid {
				a.VoteCount = int(value.Int64)
			}
		case answer.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field question_answers", value)
			} else if value.Valid {
				a.question_answers = new(int)
				*a.question_answers = int(value.Int64)
			}
		case answer.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_answers", value)
			} else if value.Valid {
				a.user_answers = new(int)
				*a.user_answers = int(value.Int64)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Answer.
// This includes values selected through modifiers, order, etc.
func (a *Answer) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryAuthor queries the "author" edge of the Answer entity.
func (a *Answer) QueryAuthor() *UserQuery {
	return NewAnswerClient(a.config).QueryAuthor(a)
}

// QueryQuestion queries the "question" edge of the Answer entity.
func (a *Answer) QueryQuestion() *QuestionQuery {
	return NewAnswerClient(a.config).QueryQuestion(a)
}

// QueryComments queries the "comments" edge of the Answer entity.
func (a *Answer) QueryComments() *CommentQuery {
	return NewAnswerClient(a.config).QueryComments(a)
}

// QueryVotes queries the "votes" edge of the Answer entity.
func (a *Answer) QueryVotes() *VoteQuery {
	return NewAnswerClient(a.config).QueryVotes(a)
}

// Update returns a builder for updating this Answer.
// Note that you need to call Answer.Unwrap() before calling this method if this Answer
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Answer) Update() *AnswerUpdateOne {
	return NewAnswerClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Answer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Answer) Unwrap() *Answer {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Answer is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Answer) String() string {
	var builder strings.Builder
	builder.WriteString("Answer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("content=")
	builder.WriteString(a.Content)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("is_accepted=")
	builder.WriteString(fmt.Sprintf("%v", a.IsAccepted))
	builder.WriteString(", ")
	builder.WriteString("vote_count=")
	builder.WriteString(fmt.Sprintf("%v", a.VoteCount))
	builder.WriteByte(')')
	return builder.String()
}

// Answers is a parsable slice of Answer.
type Answers []*Answer
