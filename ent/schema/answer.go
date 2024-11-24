package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Answer struct {
	ent.Schema
}

func (Answer) Fields() []ent.Field {
	return []ent.Field{
			field.Text("content"),
			field.Time("created_at"),
			field.Time("updated_at"),
			field.Bool("is_accepted").Default(false),
			field.Int("vote_count").Default(0),
	}
}

func (Answer) Edges() []ent.Edge {
	return []ent.Edge{
			edge.From("author", User.Type).
					Ref("answers").
					Unique(),
			edge.From("question", Question.Type).
					Ref("answers").
					Unique(),
			edge.To("comments", Comment.Type),
			edge.To("votes", Vote.Type),
	}
}