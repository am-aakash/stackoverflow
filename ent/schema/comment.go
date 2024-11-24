package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Comment struct {
	ent.Schema
}

func (Comment) Fields() []ent.Field {
	return []ent.Field{
			field.String("content"),
			field.Time("created_at"),
	}
}

func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
			edge.From("author", User.Type).
					Ref("comments").
					Unique(),
			edge.From("question", Question.Type).
					Ref("comments"),
			edge.From("answer", Answer.Type).
					Ref("comments"),
	}
}
