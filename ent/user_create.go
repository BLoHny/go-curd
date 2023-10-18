// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/blohny/ent/tourproduct"
	"github.com/blohny/ent/user"
)

// USERCreate is the builder for creating a USER entity.
type USERCreate struct {
	config
	mutation *USERMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (uc *USERCreate) SetName(s string) *USERCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetIsActivated sets the "isActivated" field.
func (uc *USERCreate) SetIsActivated(b bool) *USERCreate {
	uc.mutation.SetIsActivated(b)
	return uc
}

// SetNillableIsActivated sets the "isActivated" field if the given value is not nil.
func (uc *USERCreate) SetNillableIsActivated(b *bool) *USERCreate {
	if b != nil {
		uc.SetIsActivated(*b)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *USERCreate) SetID(s string) *USERCreate {
	uc.mutation.SetID(s)
	return uc
}

// AddProductIDs adds the "products" edge to the TourProduct entity by IDs.
func (uc *USERCreate) AddProductIDs(ids ...int) *USERCreate {
	uc.mutation.AddProductIDs(ids...)
	return uc
}

// AddProducts adds the "products" edges to the TourProduct entity.
func (uc *USERCreate) AddProducts(t ...*TourProduct) *USERCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uc.AddProductIDs(ids...)
}

// Mutation returns the USERMutation object of the builder.
func (uc *USERCreate) Mutation() *USERMutation {
	return uc.mutation
}

// Save creates the USER in the database.
func (uc *USERCreate) Save(ctx context.Context) (*USER, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *USERCreate) SaveX(ctx context.Context) *USER {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *USERCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *USERCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *USERCreate) defaults() {
	if _, ok := uc.mutation.IsActivated(); !ok {
		v := user.DefaultIsActivated
		uc.mutation.SetIsActivated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *USERCreate) check() error {
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "USER.name"`)}
	}
	if _, ok := uc.mutation.IsActivated(); !ok {
		return &ValidationError{Name: "isActivated", err: errors.New(`ent: missing required field "USER.isActivated"`)}
	}
	return nil
}

func (uc *USERCreate) sqlSave(ctx context.Context) (*USER, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected USER.ID type: %T", _spec.ID.Value)
		}
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *USERCreate) createSpec() (*USER, *sqlgraph.CreateSpec) {
	var (
		_node = &USER{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.IsActivated(); ok {
		_spec.SetField(user.FieldIsActivated, field.TypeBool, value)
		_node.IsActivated = value
	}
	if nodes := uc.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProductsTable,
			Columns: []string{user.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tourproduct.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// USERCreateBulk is the builder for creating many USER entities in bulk.
type USERCreateBulk struct {
	config
	err      error
	builders []*USERCreate
}

// Save creates the USER entities in the database.
func (ucb *USERCreateBulk) Save(ctx context.Context) ([]*USER, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*USER, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*USERMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *USERCreateBulk) SaveX(ctx context.Context) []*USER {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *USERCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *USERCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
