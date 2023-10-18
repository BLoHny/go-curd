// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/blohny/ent/tourproduct"
)

// TourProductCreate is the builder for creating a TourProduct entity.
type TourProductCreate struct {
	config
	mutation *TourProductMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tpc *TourProductCreate) SetName(s string) *TourProductCreate {
	tpc.mutation.SetName(s)
	return tpc
}

// Mutation returns the TourProductMutation object of the builder.
func (tpc *TourProductCreate) Mutation() *TourProductMutation {
	return tpc.mutation
}

// Save creates the TourProduct in the database.
func (tpc *TourProductCreate) Save(ctx context.Context) (*TourProduct, error) {
	return withHooks(ctx, tpc.sqlSave, tpc.mutation, tpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tpc *TourProductCreate) SaveX(ctx context.Context) *TourProduct {
	v, err := tpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tpc *TourProductCreate) Exec(ctx context.Context) error {
	_, err := tpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpc *TourProductCreate) ExecX(ctx context.Context) {
	if err := tpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tpc *TourProductCreate) check() error {
	if _, ok := tpc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "TourProduct.name"`)}
	}
	return nil
}

func (tpc *TourProductCreate) sqlSave(ctx context.Context) (*TourProduct, error) {
	if err := tpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tpc.mutation.id = &_node.ID
	tpc.mutation.done = true
	return _node, nil
}

func (tpc *TourProductCreate) createSpec() (*TourProduct, *sqlgraph.CreateSpec) {
	var (
		_node = &TourProduct{config: tpc.config}
		_spec = sqlgraph.NewCreateSpec(tourproduct.Table, sqlgraph.NewFieldSpec(tourproduct.FieldID, field.TypeInt))
	)
	if value, ok := tpc.mutation.Name(); ok {
		_spec.SetField(tourproduct.FieldName, field.TypeString, value)
		_node.Name = value
	}
	return _node, _spec
}

// TourProductCreateBulk is the builder for creating many TourProduct entities in bulk.
type TourProductCreateBulk struct {
	config
	err      error
	builders []*TourProductCreate
}

// Save creates the TourProduct entities in the database.
func (tpcb *TourProductCreateBulk) Save(ctx context.Context) ([]*TourProduct, error) {
	if tpcb.err != nil {
		return nil, tpcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tpcb.builders))
	nodes := make([]*TourProduct, len(tpcb.builders))
	mutators := make([]Mutator, len(tpcb.builders))
	for i := range tpcb.builders {
		func(i int, root context.Context) {
			builder := tpcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TourProductMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tpcb *TourProductCreateBulk) SaveX(ctx context.Context) []*TourProduct {
	v, err := tpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tpcb *TourProductCreateBulk) Exec(ctx context.Context) error {
	_, err := tpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpcb *TourProductCreateBulk) ExecX(ctx context.Context) {
	if err := tpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
