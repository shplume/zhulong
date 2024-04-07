package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Review holds the schema definition for the Review entity.
type Review struct {
	ent.Schema
}

// Fields of the Review.
func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("status"),
		field.String("remark").Nillable(),
	}
}

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("thesis", File.Type).Required().Unique(),
		edge.To("student_id", Student.Type).Required().Unique(),
		edge.To("teacher_id", Teacher.Type).Unique(),
		edge.To("book_id", File.Type).Unique(),
	}
}
