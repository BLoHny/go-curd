package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// USER holds the schema definition for the USER entity.
type USER struct {
	ent.Schema
}

// Fields of the USER.
func (USER) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
		field.Bool("isActivated").Default(true),
	}

}

// Edges of the USER.
func (USER) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products", TourProduct.Type),
	}

}
