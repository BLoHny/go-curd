package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TourProduct holds the schema definition for the TourProduct entity.
type TourProduct struct {
	ent.Schema
}

// Fields of the TourProduct.
func (TourProduct) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the TourProduct.
func (TourProduct) Edges() []ent.Edge {
	return nil
}
