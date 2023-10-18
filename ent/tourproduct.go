// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/blohny/ent/tourproduct"
)

// TourProduct is the model entity for the TourProduct schema.
type TourProduct struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name          string `json:"name,omitempty"`
	user_products *string
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TourProduct) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tourproduct.FieldID:
			values[i] = new(sql.NullInt64)
		case tourproduct.FieldName:
			values[i] = new(sql.NullString)
		case tourproduct.ForeignKeys[0]: // user_products
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TourProduct fields.
func (tp *TourProduct) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tourproduct.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tp.ID = int(value.Int64)
		case tourproduct.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				tp.Name = value.String
			}
		case tourproduct.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_products", values[i])
			} else if value.Valid {
				tp.user_products = new(string)
				*tp.user_products = value.String
			}
		default:
			tp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TourProduct.
// This includes values selected through modifiers, order, etc.
func (tp *TourProduct) Value(name string) (ent.Value, error) {
	return tp.selectValues.Get(name)
}

// Update returns a builder for updating this TourProduct.
// Note that you need to call TourProduct.Unwrap() before calling this method if this TourProduct
// was returned from a transaction, and the transaction was committed or rolled back.
func (tp *TourProduct) Update() *TourProductUpdateOne {
	return NewTourProductClient(tp.config).UpdateOne(tp)
}

// Unwrap unwraps the TourProduct entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tp *TourProduct) Unwrap() *TourProduct {
	_tx, ok := tp.config.driver.(*txDriver)
	if !ok {
		panic("ent: TourProduct is not a transactional entity")
	}
	tp.config.driver = _tx.drv
	return tp
}

// String implements the fmt.Stringer.
func (tp *TourProduct) String() string {
	var builder strings.Builder
	builder.WriteString("TourProduct(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tp.ID))
	builder.WriteString("name=")
	builder.WriteString(tp.Name)
	builder.WriteByte(')')
	return builder.String()
}

// TourProducts is a parsable slice of TourProduct.
type TourProducts []*TourProduct
