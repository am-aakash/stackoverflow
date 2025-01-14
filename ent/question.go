// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"stackoverflow-clone/ent/question"
	"stackoverflow-clone/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Question is the model entity for the Question schema.
type Question struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// VoteCount holds the value of the "vote_count" field.
	VoteCount int `json:"vote_count,omitempty"`
	// ViewCount holds the value of the "view_count" field.
	ViewCount int `json:"view_count,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the QuestionQuery when eager-loading is set.
	Edges          QuestionEdges `json:"edges"`
	user_questions *int
	selectValues   sql.SelectValues
}

// QuestionEdges holds the relations/edges for other nodes in the graph.
type QuestionEdges struct {
	// Author holds the value of the author edge.
	Author *User `json:"author,omitempty"`
	// Answers holds the value of the answers edge.
	Answers []*Answer `json:"answers,omitempty"`
	// Comments holds the value of the comments edge.
	Comments []*Comment `json:"comments,omitempty"`
	// Tags holds the value of the tags edge.
	Tags []*Tag `json:"tags,omitempty"`
	// Votes holds the value of the votes edge.
	Votes []*Vote `json:"votes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e QuestionEdges) AuthorOrErr() (*User, error) {
	if e.Author != nil {
		return e.Author, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "author"}
}

// AnswersOrErr returns the Answers value or an error if the edge
// was not loaded in eager-loading.
func (e QuestionEdges) AnswersOrErr() ([]*Answer, error) {
	if e.loadedTypes[1] {
		return e.Answers, nil
	}
	return nil, &NotLoadedError{edge: "answers"}
}

// CommentsOrErr returns the Comments value or an error if the edge
// was not loaded in eager-loading.
func (e QuestionEdges) CommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[2] {
		return e.Comments, nil
	}
	return nil, &NotLoadedError{edge: "comments"}
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e QuestionEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[3] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// VotesOrErr returns the Votes value or an error if the edge
// was not loaded in eager-loading.
func (e QuestionEdges) VotesOrErr() ([]*Vote, error) {
	if e.loadedTypes[4] {
		return e.Votes, nil
	}
	return nil, &NotLoadedError{edge: "votes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Question) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case question.FieldID, question.FieldVoteCount, question.FieldViewCount:
			values[i] = new(sql.NullInt64)
		case question.FieldTitle, question.FieldContent:
			values[i] = new(sql.NullString)
		case question.FieldCreatedAt, question.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case question.ForeignKeys[0]: // user_questions
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Question fields.
func (q *Question) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case question.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			q.ID = int(value.Int64)
		case question.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				q.Title = value.String
			}
		case question.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				q.Content = value.String
			}
		case question.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				q.CreatedAt = value.Time
			}
		case question.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				q.UpdatedAt = value.Time
			}
		case question.FieldVoteCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field vote_count", values[i])
			} else if value.Valid {
				q.VoteCount = int(value.Int64)
			}
		case question.FieldViewCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field view_count", values[i])
			} else if value.Valid {
				q.ViewCount = int(value.Int64)
			}
		case question.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_questions", value)
			} else if value.Valid {
				q.user_questions = new(int)
				*q.user_questions = int(value.Int64)
			}
		default:
			q.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Question.
// This includes values selected through modifiers, order, etc.
func (q *Question) Value(name string) (ent.Value, error) {
	return q.selectValues.Get(name)
}

// QueryAuthor queries the "author" edge of the Question entity.
func (q *Question) QueryAuthor() *UserQuery {
	return NewQuestionClient(q.config).QueryAuthor(q)
}

// QueryAnswers queries the "answers" edge of the Question entity.
func (q *Question) QueryAnswers() *AnswerQuery {
	return NewQuestionClient(q.config).QueryAnswers(q)
}

// QueryComments queries the "comments" edge of the Question entity.
func (q *Question) QueryComments() *CommentQuery {
	return NewQuestionClient(q.config).QueryComments(q)
}

// QueryTags queries the "tags" edge of the Question entity.
func (q *Question) QueryTags() *TagQuery {
	return NewQuestionClient(q.config).QueryTags(q)
}

// QueryVotes queries the "votes" edge of the Question entity.
func (q *Question) QueryVotes() *VoteQuery {
	return NewQuestionClient(q.config).QueryVotes(q)
}

// Update returns a builder for updating this Question.
// Note that you need to call Question.Unwrap() before calling this method if this Question
// was returned from a transaction, and the transaction was committed or rolled back.
func (q *Question) Update() *QuestionUpdateOne {
	return NewQuestionClient(q.config).UpdateOne(q)
}

// Unwrap unwraps the Question entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (q *Question) Unwrap() *Question {
	_tx, ok := q.config.driver.(*txDriver)
	if !ok {
		panic("ent: Question is not a transactional entity")
	}
	q.config.driver = _tx.drv
	return q
}

// String implements the fmt.Stringer.
func (q *Question) String() string {
	var builder strings.Builder
	builder.WriteString("Question(")
	builder.WriteString(fmt.Sprintf("id=%v, ", q.ID))
	builder.WriteString("title=")
	builder.WriteString(q.Title)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(q.Content)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(q.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(q.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("vote_count=")
	builder.WriteString(fmt.Sprintf("%v", q.VoteCount))
	builder.WriteString(", ")
	builder.WriteString("view_count=")
	builder.WriteString(fmt.Sprintf("%v", q.ViewCount))
	builder.WriteByte(')')
	return builder.String()
}

// Questions is a parsable slice of Question.
type Questions []*Question
