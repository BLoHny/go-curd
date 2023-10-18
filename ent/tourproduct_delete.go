// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/blohny/ent/predicate"
	"github.com/blohny/ent/tourproduct"
)

// TourProductDelete is the builder for deleting a TourProduct entity.
type TourProductDelete struct {
	config
	hooks    []Hook
	mutation *TourProductMutation
}

// Where appends a list predicates to the TourProductDelete builder.
func (tpd *TourProductDelete) Where(ps ...predicate.TourProduct) *TourProductDelete {
	tpd.mutation.Where(ps...)
	return tpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tpd *TourProductDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tpd.sqlExec, tpd.mutation, tpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tpd *TourProductDelete) ExecX(ctx context.Context) int {
	n, err := tpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tpd *TourProductDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(tourproduct.Table, sqlgraph.NewFieldSpec(tourproduct.FieldID, field.TypeInt))
	if ps := tpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tpd.mutation.done = true
	return affected, err
}

// TourProductDeleteOne is the builder for deleting a single TourProduct entity.
type TourProductDeleteOne struct {
	tpd *TourProductDelete
}

// Where appends a list predicates to the TourProductDelete builder.
func (tpdo *TourProductDeleteOne) Where(ps ...predicate.TourProduct) *TourProductDeleteOne {
	tpdo.tpd.mutation.Where(ps...)
	return tpdo
}

// Exec executes the deletion query.
func (tpdo *TourProductDeleteOne) Exec(ctx context.Context) error {
	n, err := tpdo.tpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tourproduct.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tpdo *TourProductDeleteOne) ExecX(ctx context.Context) {
	if err := tpdo.Exec(ctx); err != nil {
		panic(err)
	}
}