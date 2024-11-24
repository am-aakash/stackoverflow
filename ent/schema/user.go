package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique(),
		field.String("email").Unique(),
		field.String("password_hash"),
		field.String("display_name"),
		field.Time("created_at"),
		field.Int("reputation").Default(0),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("questions", Question.Type),
		edge.To("answers", Answer.Type),
		edge.To("comments", Comment.Type),
		edge.To("votes", Vote.Type),
	}
}
