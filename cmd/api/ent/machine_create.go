// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/crowdsecurity/crowdsec/cmd/api/ent/alert"
	"github.com/crowdsecurity/crowdsec/cmd/api/ent/machine"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// MachineCreate is the builder for creating a Machine entity.
type MachineCreate struct {
	config
	mutation *MachineMutation
	hooks    []Hook
}

// SetCreatedAt sets the created_at field.
func (mc *MachineCreate) SetCreatedAt(t time.Time) *MachineCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (mc *MachineCreate) SetNillableCreatedAt(t *time.Time) *MachineCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the updated_at field.
func (mc *MachineCreate) SetUpdatedAt(t time.Time) *MachineCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (mc *MachineCreate) SetNillableUpdatedAt(t *time.Time) *MachineCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetMachineId sets the machineId field.
func (mc *MachineCreate) SetMachineId(s string) *MachineCreate {
	mc.mutation.SetMachineId(s)
	return mc
}

// SetPassword sets the password field.
func (mc *MachineCreate) SetPassword(s string) *MachineCreate {
	mc.mutation.SetPassword(s)
	return mc
}

// SetIpAddress sets the ipAddress field.
func (mc *MachineCreate) SetIpAddress(s string) *MachineCreate {
	mc.mutation.SetIpAddress(s)
	return mc
}

// SetIsValidated sets the isValidated field.
func (mc *MachineCreate) SetIsValidated(b bool) *MachineCreate {
	mc.mutation.SetIsValidated(b)
	return mc
}

// SetNillableIsValidated sets the isValidated field if the given value is not nil.
func (mc *MachineCreate) SetNillableIsValidated(b *bool) *MachineCreate {
	if b != nil {
		mc.SetIsValidated(*b)
	}
	return mc
}

// SetStatus sets the status field.
func (mc *MachineCreate) SetStatus(s string) *MachineCreate {
	mc.mutation.SetStatus(s)
	return mc
}

// SetNillableStatus sets the status field if the given value is not nil.
func (mc *MachineCreate) SetNillableStatus(s *string) *MachineCreate {
	if s != nil {
		mc.SetStatus(*s)
	}
	return mc
}

// AddSignalIDs adds the signals edge to Alert by ids.
func (mc *MachineCreate) AddSignalIDs(ids ...int) *MachineCreate {
	mc.mutation.AddSignalIDs(ids...)
	return mc
}

// AddSignals adds the signals edges to Alert.
func (mc *MachineCreate) AddSignals(a ...*Alert) *MachineCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return mc.AddSignalIDs(ids...)
}

// Mutation returns the MachineMutation object of the builder.
func (mc *MachineCreate) Mutation() *MachineMutation {
	return mc.mutation
}

// Save creates the Machine in the database.
func (mc *MachineCreate) Save(ctx context.Context) (*Machine, error) {
	if err := mc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Machine
	)
	if len(mc.hooks) == 0 {
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MachineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mc.mutation = mutation
			node, err = mc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MachineCreate) SaveX(ctx context.Context) *Machine {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mc *MachineCreate) preSave() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := machine.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := machine.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.MachineId(); !ok {
		return &ValidationError{Name: "machineId", err: errors.New("ent: missing required field \"machineId\"")}
	}
	if _, ok := mc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New("ent: missing required field \"password\"")}
	}
	if _, ok := mc.mutation.IpAddress(); !ok {
		return &ValidationError{Name: "ipAddress", err: errors.New("ent: missing required field \"ipAddress\"")}
	}
	if _, ok := mc.mutation.IsValidated(); !ok {
		v := machine.DefaultIsValidated
		mc.mutation.SetIsValidated(v)
	}
	return nil
}

func (mc *MachineCreate) sqlSave(ctx context.Context) (*Machine, error) {
	m, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	m.ID = int(id)
	return m, nil
}

func (mc *MachineCreate) createSpec() (*Machine, *sqlgraph.CreateSpec) {
	var (
		m     = &Machine{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: machine.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: machine.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: machine.FieldCreatedAt,
		})
		m.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: machine.FieldUpdatedAt,
		})
		m.UpdatedAt = value
	}
	if value, ok := mc.mutation.MachineId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldMachineId,
		})
		m.MachineId = value
	}
	if value, ok := mc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldPassword,
		})
		m.Password = value
	}
	if value, ok := mc.mutation.IpAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldIpAddress,
		})
		m.IpAddress = value
	}
	if value, ok := mc.mutation.IsValidated(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: machine.FieldIsValidated,
		})
		m.IsValidated = value
	}
	if value, ok := mc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldStatus,
		})
		m.Status = value
	}
	if nodes := mc.mutation.SignalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   machine.SignalsTable,
			Columns: []string{machine.SignalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: alert.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return m, _spec
}

// MachineCreateBulk is the builder for creating a bulk of Machine entities.
type MachineCreateBulk struct {
	config
	builders []*MachineCreate
}

// Save creates the Machine entities in the database.
func (mcb *MachineCreateBulk) Save(ctx context.Context) ([]*Machine, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Machine, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*MachineMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (mcb *MachineCreateBulk) SaveX(ctx context.Context) []*Machine {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
