package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("account").Unique(),
		field.String("password"),
		field.Int("role"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("administrators", Administrators.Type).
			Unique(),
		edge.To("students", Students.Type).
			Unique(),
		edge.To("teachers", Teachers.Type).
			Unique(),
		edge.To("thesis", Thesis.Type),
		edge.To("reviews", Reviews.Type),
		edge.From("examineThesis", Thesis.Type).
			Ref("examine"),
	}
}
