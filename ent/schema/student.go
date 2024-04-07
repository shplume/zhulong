package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("name").Nillable(),
		field.String("college").Nillable(),
		field.String("subject").Nillable(),
		field.String("class").Nillable(),
		field.String("identity").Unique(),
	}
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_id", User.Type).Required().Unique(),
	}
}
