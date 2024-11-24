package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Vote struct {
	ent.Schema
}

func (Vote) Fields() []ent.Field {
	return []ent.Field{
			field.Int("value").Comment("1 for upvote, -1 for downvote"),
			field.Time("created_at"),
	}
}

func (Vote) Edges() []ent.Edge {
	return []ent.Edge{
			edge.From("user", User.Type).
					Ref("votes").
					Unique(),
			edge.From("question", Question.Type).
					Ref("votes"),
			edge.From("answer", Answer.Type).
					Ref("votes"),
	}
}