// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"context"
	"enttry/entgen/profile"
	"enttry/entgen/tenant"
	"enttry/entgen/user"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProfileCreate is the builder for creating a Profile entity.
type ProfileCreate struct {
	config
	mutation *ProfileMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pc *ProfileCreate) SetCreateTime(t time.Time) *ProfileCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableCreateTime(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *ProfileCreate) SetUpdateTime(t time.Time) *ProfileCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableUpdateTime(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetNickName sets the "nick_name" field.
func (pc *ProfileCreate) SetNickName(s string) *ProfileCreate {
	pc.mutation.SetNickName(s)
	return pc
}

// SetOwnerID sets the "owner_id" field.
func (pc *ProfileCreate) SetOwnerID(s string) *ProfileCreate {
	pc.mutation.SetOwnerID(s)
	return pc
}

// SetTenantID sets the "tenant_id" field.
func (pc *ProfileCreate) SetTenantID(s string) *ProfileCreate {
	pc.mutation.SetTenantID(s)
	return pc
}

// SetID sets the "id" field.
func (pc *ProfileCreate) SetID(s string) *ProfileCreate {
	pc.mutation.SetID(s)
	return pc
}

// SetOwner sets the "owner" edge to the User entity.
func (pc *ProfileCreate) SetOwner(u *User) *ProfileCreate {
	return pc.SetOwnerID(u.ID)
}

// SetTenant sets the "tenant" edge to the Tenant entity.
func (pc *ProfileCreate) SetTenant(t *Tenant) *ProfileCreate {
	return pc.SetTenantID(t.ID)
}

// Mutation returns the ProfileMutation object of the builder.
func (pc *ProfileCreate) Mutation() *ProfileMutation {
	return pc.mutation
}

// Save creates the Profile in the database.
func (pc *ProfileCreate) Save(ctx context.Context) (*Profile, error) {
	var (
		err  error
		node *Profile
	)
	if err := pc.defaults(); err != nil {
		return nil, err
	}
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("entgen: uninitialized hook (forgotten import entgen/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Profile)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ProfileMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProfileCreate) SaveX(ctx context.Context) *Profile {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProfileCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProfileCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProfileCreate) defaults() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		if profile.DefaultCreateTime == nil {
			return fmt.Errorf("entgen: uninitialized profile.DefaultCreateTime (forgotten import entgen/runtime?)")
		}
		v := profile.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		if profile.DefaultUpdateTime == nil {
			return fmt.Errorf("entgen: uninitialized profile.DefaultUpdateTime (forgotten import entgen/runtime?)")
		}
		v := profile.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProfileCreate) check() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`entgen: missing required field "Profile.create_time"`)}
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`entgen: missing required field "Profile.update_time"`)}
	}
	if _, ok := pc.mutation.NickName(); !ok {
		return &ValidationError{Name: "nick_name", err: errors.New(`entgen: missing required field "Profile.nick_name"`)}
	}
	if v, ok := pc.mutation.NickName(); ok {
		if err := profile.NickNameValidator(v); err != nil {
			return &ValidationError{Name: "nick_name", err: fmt.Errorf(`entgen: validator failed for field "Profile.nick_name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`entgen: missing required field "Profile.owner_id"`)}
	}
	if _, ok := pc.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant_id", err: errors.New(`entgen: missing required field "Profile.tenant_id"`)}
	}
	if _, ok := pc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`entgen: missing required edge "Profile.owner"`)}
	}
	if _, ok := pc.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant", err: errors.New(`entgen: missing required edge "Profile.tenant"`)}
	}
	return nil
}

func (pc *ProfileCreate) sqlSave(ctx context.Context) (*Profile, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Profile.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (pc *ProfileCreate) createSpec() (*Profile, *sqlgraph.CreateSpec) {
	var (
		_node = &Profile{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: profile.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: profile.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := pc.mutation.NickName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldNickName,
		})
		_node.NickName = value
	}
	if nodes := pc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   profile.OwnerTable,
			Columns: []string{profile.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.TenantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   profile.TenantTable,
			Columns: []string{profile.TenantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tenant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TenantID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProfileCreateBulk is the builder for creating many Profile entities in bulk.
type ProfileCreateBulk struct {
	config
	builders []*ProfileCreate
}

// Save creates the Profile entities in the database.
func (pcb *ProfileCreateBulk) Save(ctx context.Context) ([]*Profile, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Profile, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProfileMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProfileCreateBulk) SaveX(ctx context.Context) []*Profile {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProfileCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProfileCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
