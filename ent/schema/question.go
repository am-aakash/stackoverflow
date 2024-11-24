package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Question struct {
	ent.Schema
}

func (Question) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Text("content"),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.Int("vote_count").Default(0),
		field.Int("view_count").Default(0),
	}
}

func (Question) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).
			Ref("questions").
			Unique().
			Required(),
		edge.To("answers", Answer.Type),
		edge.To("comments", Comment.Type),
		edge.To("tags", Tag.Type),
		edge.To("votes", Vote.Type),
	}
}
