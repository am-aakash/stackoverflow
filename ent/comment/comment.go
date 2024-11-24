// Code generated by ent, DO NOT EDIT.

package comment

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the comment type in the database.
	Label = "comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// EdgeQuestion holds the string denoting the question edge name in mutations.
	EdgeQuestion = "question"
	// EdgeAnswer holds the string denoting the answer edge name in mutations.
	EdgeAnswer = "answer"
	// Table holds the table name of the comment in the database.
	Table = "comments"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "comments"
	// AuthorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AuthorInverseTable = "users"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "user_comments"
	// QuestionTable is the table that holds the question relation/edge. The primary key declared below.
	QuestionTable = "question_comments"
	// QuestionInverseTable is the table name for the Question entity.
	// It exists in this package in order to avoid circular dependency with the "question" package.
	QuestionInverseTable = "questions"
	// AnswerTable is the table that holds the answer relation/edge. The primary key declared below.
	AnswerTable = "answer_comments"
	// AnswerInverseTable is the table name for the Answer entity.
	// It exists in this package in order to avoid circular dependency with the "answer" package.
	AnswerInverseTable = "answers"
)

// Columns holds all SQL columns for comment fields.
var Columns = []string{
	FieldID,
	FieldContent,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "comments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_comments",
}

var (
	// QuestionPrimaryKey and QuestionColumn2 are the table columns denoting the
	// primary key for the question relation (M2M).
	QuestionPrimaryKey = []string{"question_id", "comment_id"}
	// AnswerPrimaryKey and AnswerColumn2 are the table columns denoting the
	// primary key for the answer relation (M2M).
	AnswerPrimaryKey = []string{"answer_id", "comment_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Comment queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByAuthorField orders the results by author field.
func ByAuthorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthorStep(), sql.OrderByField(field, opts...))
	}
}

// ByQuestionCount orders the results by question count.
func ByQuestionCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newQuestionStep(), opts...)
	}
}

// ByQuestion orders the results by question terms.
func ByQuestion(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newQuestionStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAnswerCount orders the results by answer count.
func ByAnswerCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAnswerStep(), opts...)
	}
}

// ByAnswer orders the results by answer terms.
func ByAnswer(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAnswerStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAuthorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
	)
}
func newQuestionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(QuestionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, QuestionTable, QuestionPrimaryKey...),
	)
}
func newAnswerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AnswerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, AnswerTable, AnswerPrimaryKey...),
	)
}