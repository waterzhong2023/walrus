// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/perspective"
	"github.com/seal-io/seal/pkg/dao/types"
)

// PerspectiveCreate is the builder for creating a Perspective entity.
type PerspectiveCreate struct {
	config
	mutation *PerspectiveMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "createTime" field.
func (pc *PerspectiveCreate) SetCreateTime(t time.Time) *PerspectiveCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "createTime" field if the given value is not nil.
func (pc *PerspectiveCreate) SetNillableCreateTime(t *time.Time) *PerspectiveCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the "updateTime" field.
func (pc *PerspectiveCreate) SetUpdateTime(t time.Time) *PerspectiveCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "updateTime" field if the given value is not nil.
func (pc *PerspectiveCreate) SetNillableUpdateTime(t *time.Time) *PerspectiveCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *PerspectiveCreate) SetName(s string) *PerspectiveCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetStartTime sets the "startTime" field.
func (pc *PerspectiveCreate) SetStartTime(s string) *PerspectiveCreate {
	pc.mutation.SetStartTime(s)
	return pc
}

// SetEndTime sets the "endTime" field.
func (pc *PerspectiveCreate) SetEndTime(s string) *PerspectiveCreate {
	pc.mutation.SetEndTime(s)
	return pc
}

// SetBuiltin sets the "builtin" field.
func (pc *PerspectiveCreate) SetBuiltin(b bool) *PerspectiveCreate {
	pc.mutation.SetBuiltin(b)
	return pc
}

// SetNillableBuiltin sets the "builtin" field if the given value is not nil.
func (pc *PerspectiveCreate) SetNillableBuiltin(b *bool) *PerspectiveCreate {
	if b != nil {
		pc.SetBuiltin(*b)
	}
	return pc
}

// SetAllocationQueries sets the "allocationQueries" field.
func (pc *PerspectiveCreate) SetAllocationQueries(tc []types.QueryCondition) *PerspectiveCreate {
	pc.mutation.SetAllocationQueries(tc)
	return pc
}

// SetID sets the "id" field.
func (pc *PerspectiveCreate) SetID(t types.ID) *PerspectiveCreate {
	pc.mutation.SetID(t)
	return pc
}

// Mutation returns the PerspectiveMutation object of the builder.
func (pc *PerspectiveCreate) Mutation() *PerspectiveMutation {
	return pc.mutation
}

// Save creates the Perspective in the database.
func (pc *PerspectiveCreate) Save(ctx context.Context) (*Perspective, error) {
	if err := pc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*Perspective, PerspectiveMutation](ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PerspectiveCreate) SaveX(ctx context.Context) *Perspective {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PerspectiveCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PerspectiveCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PerspectiveCreate) defaults() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		if perspective.DefaultCreateTime == nil {
			return fmt.Errorf("model: uninitialized perspective.DefaultCreateTime (forgotten import model/runtime?)")
		}
		v := perspective.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		if perspective.DefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized perspective.DefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := perspective.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
	if _, ok := pc.mutation.Builtin(); !ok {
		v := perspective.DefaultBuiltin
		pc.mutation.SetBuiltin(v)
	}
	if _, ok := pc.mutation.AllocationQueries(); !ok {
		v := perspective.DefaultAllocationQueries
		pc.mutation.SetAllocationQueries(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pc *PerspectiveCreate) check() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "createTime", err: errors.New(`model: missing required field "Perspective.createTime"`)}
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "updateTime", err: errors.New(`model: missing required field "Perspective.updateTime"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`model: missing required field "Perspective.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := perspective.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`model: validator failed for field "Perspective.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.StartTime(); !ok {
		return &ValidationError{Name: "startTime", err: errors.New(`model: missing required field "Perspective.startTime"`)}
	}
	if v, ok := pc.mutation.StartTime(); ok {
		if err := perspective.StartTimeValidator(v); err != nil {
			return &ValidationError{Name: "startTime", err: fmt.Errorf(`model: validator failed for field "Perspective.startTime": %w`, err)}
		}
	}
	if _, ok := pc.mutation.EndTime(); !ok {
		return &ValidationError{Name: "endTime", err: errors.New(`model: missing required field "Perspective.endTime"`)}
	}
	if v, ok := pc.mutation.EndTime(); ok {
		if err := perspective.EndTimeValidator(v); err != nil {
			return &ValidationError{Name: "endTime", err: fmt.Errorf(`model: validator failed for field "Perspective.endTime": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Builtin(); !ok {
		return &ValidationError{Name: "builtin", err: errors.New(`model: missing required field "Perspective.builtin"`)}
	}
	if _, ok := pc.mutation.AllocationQueries(); !ok {
		return &ValidationError{Name: "allocationQueries", err: errors.New(`model: missing required field "Perspective.allocationQueries"`)}
	}
	return nil
}

func (pc *PerspectiveCreate) sqlSave(ctx context.Context) (*Perspective, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*types.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PerspectiveCreate) createSpec() (*Perspective, *sqlgraph.CreateSpec) {
	var (
		_node = &Perspective{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: perspective.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: perspective.FieldID,
			},
		}
	)
	_spec.Schema = pc.schemaConfig.Perspective
	_spec.OnConflict = pc.conflict
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.SetField(perspective.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = &value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.SetField(perspective.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = &value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(perspective.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.StartTime(); ok {
		_spec.SetField(perspective.FieldStartTime, field.TypeString, value)
		_node.StartTime = value
	}
	if value, ok := pc.mutation.EndTime(); ok {
		_spec.SetField(perspective.FieldEndTime, field.TypeString, value)
		_node.EndTime = value
	}
	if value, ok := pc.mutation.Builtin(); ok {
		_spec.SetField(perspective.FieldBuiltin, field.TypeBool, value)
		_node.Builtin = value
	}
	if value, ok := pc.mutation.AllocationQueries(); ok {
		_spec.SetField(perspective.FieldAllocationQueries, field.TypeJSON, value)
		_node.AllocationQueries = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Perspective.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PerspectiveUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (pc *PerspectiveCreate) OnConflict(opts ...sql.ConflictOption) *PerspectiveUpsertOne {
	pc.conflict = opts
	return &PerspectiveUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Perspective.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PerspectiveCreate) OnConflictColumns(columns ...string) *PerspectiveUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PerspectiveUpsertOne{
		create: pc,
	}
}

type (
	// PerspectiveUpsertOne is the builder for "upsert"-ing
	//  one Perspective node.
	PerspectiveUpsertOne struct {
		create *PerspectiveCreate
	}

	// PerspectiveUpsert is the "OnConflict" setter.
	PerspectiveUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "updateTime" field.
func (u *PerspectiveUpsert) SetUpdateTime(v time.Time) *PerspectiveUpsert {
	u.Set(perspective.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "updateTime" field to the value that was provided on create.
func (u *PerspectiveUpsert) UpdateUpdateTime() *PerspectiveUpsert {
	u.SetExcluded(perspective.FieldUpdateTime)
	return u
}

// SetStartTime sets the "startTime" field.
func (u *PerspectiveUpsert) SetStartTime(v string) *PerspectiveUpsert {
	u.Set(perspective.FieldStartTime, v)
	return u
}

// UpdateStartTime sets the "startTime" field to the value that was provided on create.
func (u *PerspectiveUpsert) UpdateStartTime() *PerspectiveUpsert {
	u.SetExcluded(perspective.FieldStartTime)
	return u
}

// SetEndTime sets the "endTime" field.
func (u *PerspectiveUpsert) SetEndTime(v string) *PerspectiveUpsert {
	u.Set(perspective.FieldEndTime, v)
	return u
}

// UpdateEndTime sets the "endTime" field to the value that was provided on create.
func (u *PerspectiveUpsert) UpdateEndTime() *PerspectiveUpsert {
	u.SetExcluded(perspective.FieldEndTime)
	return u
}

// SetBuiltin sets the "builtin" field.
func (u *PerspectiveUpsert) SetBuiltin(v bool) *PerspectiveUpsert {
	u.Set(perspective.FieldBuiltin, v)
	return u
}

// UpdateBuiltin sets the "builtin" field to the value that was provided on create.
func (u *PerspectiveUpsert) UpdateBuiltin() *PerspectiveUpsert {
	u.SetExcluded(perspective.FieldBuiltin)
	return u
}

// SetAllocationQueries sets the "allocationQueries" field.
func (u *PerspectiveUpsert) SetAllocationQueries(v []types.QueryCondition) *PerspectiveUpsert {
	u.Set(perspective.FieldAllocationQueries, v)
	return u
}

// UpdateAllocationQueries sets the "allocationQueries" field to the value that was provided on create.
func (u *PerspectiveUpsert) UpdateAllocationQueries() *PerspectiveUpsert {
	u.SetExcluded(perspective.FieldAllocationQueries)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Perspective.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(perspective.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PerspectiveUpsertOne) UpdateNewValues() *PerspectiveUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(perspective.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(perspective.FieldCreateTime)
		}
		if _, exists := u.create.mutation.Name(); exists {
			s.SetIgnore(perspective.FieldName)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Perspective.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PerspectiveUpsertOne) Ignore() *PerspectiveUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PerspectiveUpsertOne) DoNothing() *PerspectiveUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PerspectiveCreate.OnConflict
// documentation for more info.
func (u *PerspectiveUpsertOne) Update(set func(*PerspectiveUpsert)) *PerspectiveUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PerspectiveUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "updateTime" field.
func (u *PerspectiveUpsertOne) SetUpdateTime(v time.Time) *PerspectiveUpsertOne {
	return u.Update(func(s *PerspectiveUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "updateTime" field to the value that was provided on create.
func (u *PerspectiveUpsertOne) UpdateUpdateTime() *PerspectiveUpsertOne {
	return u.Update(func(s *PerspectiveUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetStartTime sets the "startTime" field.
func (u *PerspectiveUpsertOne) SetStartTime(v string) *PerspectiveUpsertOne {
	return u.Update(func(s *PerspectiveUpsert) {
		s.SetStartTime(v)
	})
}

// UpdateStartTime sets the "startTime" field to the value that was provided on create.
func (u *PerspectiveUpsertOne) UpdateStartTime() *PerspectiveUpsertOne {
	return u.Update(func(s *PerspectiveUpsert) {
		s.UpdateStartTime()
	})
}

// SetEndTime sets the "endTime" field.
func (u *PerspectiveUpsertOne) SetEndTime(v string) *PerspectiveUpsertOne {
	return u.Update(func(s *PerspectiveUpsert) {
		s.SetEndTime(v)
	})
}

// UpdateEndTime sets the "endTime" field to the value that was provided on create.
func (u *PerspectiveUpsertOne) UpdateEndTime() *PerspectiveUpsertOne {
	return u.Update(func(s *PerspectiveUpsert) {
		s.UpdateEndTime()
	})
}

// SetBuiltin sets the "builtin" field.
func (u *PerspectiveUpsertOne) SetBuiltin(v bool) *PerspectiveUpsertOne {
	return u.Update(func(s *PerspectiveUpsert) {
		s.SetBuiltin(v)
	})
}

// UpdateBuiltin sets the "builtin" field to the value that was provided on create.
func (u *PerspectiveUpsertOne) UpdateBuiltin() *PerspectiveUpsertOne {
	return u.Update(func(s *PerspectiveUpsert) {
		s.UpdateBuiltin()
	})
}

// SetAllocationQueries sets the "allocationQueries" field.
func (u *PerspectiveUpsertOne) SetAllocationQueries(v []types.QueryCondition) *PerspectiveUpsertOne {
	return u.Update(func(s *PerspectiveUpsert) {
		s.SetAllocationQueries(v)
	})
}

// UpdateAllocationQueries sets the "allocationQueries" field to the value that was provided on create.
func (u *PerspectiveUpsertOne) UpdateAllocationQueries() *PerspectiveUpsertOne {
	return u.Update(func(s *PerspectiveUpsert) {
		s.UpdateAllocationQueries()
	})
}

// Exec executes the query.
func (u *PerspectiveUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for PerspectiveCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PerspectiveUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PerspectiveUpsertOne) ID(ctx context.Context) (id types.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("model: PerspectiveUpsertOne.ID is not supported by MySQL driver. Use PerspectiveUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PerspectiveUpsertOne) IDX(ctx context.Context) types.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PerspectiveCreateBulk is the builder for creating many Perspective entities in bulk.
type PerspectiveCreateBulk struct {
	config
	builders []*PerspectiveCreate
	conflict []sql.ConflictOption
}

// Save creates the Perspective entities in the database.
func (pcb *PerspectiveCreateBulk) Save(ctx context.Context) ([]*Perspective, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Perspective, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PerspectiveMutation)
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
					spec.OnConflict = pcb.conflict
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
func (pcb *PerspectiveCreateBulk) SaveX(ctx context.Context) []*Perspective {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PerspectiveCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PerspectiveCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Perspective.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PerspectiveUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (pcb *PerspectiveCreateBulk) OnConflict(opts ...sql.ConflictOption) *PerspectiveUpsertBulk {
	pcb.conflict = opts
	return &PerspectiveUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Perspective.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PerspectiveCreateBulk) OnConflictColumns(columns ...string) *PerspectiveUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PerspectiveUpsertBulk{
		create: pcb,
	}
}

// PerspectiveUpsertBulk is the builder for "upsert"-ing
// a bulk of Perspective nodes.
type PerspectiveUpsertBulk struct {
	create *PerspectiveCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Perspective.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(perspective.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PerspectiveUpsertBulk) UpdateNewValues() *PerspectiveUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(perspective.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(perspective.FieldCreateTime)
			}
			if _, exists := b.mutation.Name(); exists {
				s.SetIgnore(perspective.FieldName)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Perspective.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PerspectiveUpsertBulk) Ignore() *PerspectiveUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PerspectiveUpsertBulk) DoNothing() *PerspectiveUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PerspectiveCreateBulk.OnConflict
// documentation for more info.
func (u *PerspectiveUpsertBulk) Update(set func(*PerspectiveUpsert)) *PerspectiveUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PerspectiveUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "updateTime" field.
func (u *PerspectiveUpsertBulk) SetUpdateTime(v time.Time) *PerspectiveUpsertBulk {
	return u.Update(func(s *PerspectiveUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "updateTime" field to the value that was provided on create.
func (u *PerspectiveUpsertBulk) UpdateUpdateTime() *PerspectiveUpsertBulk {
	return u.Update(func(s *PerspectiveUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetStartTime sets the "startTime" field.
func (u *PerspectiveUpsertBulk) SetStartTime(v string) *PerspectiveUpsertBulk {
	return u.Update(func(s *PerspectiveUpsert) {
		s.SetStartTime(v)
	})
}

// UpdateStartTime sets the "startTime" field to the value that was provided on create.
func (u *PerspectiveUpsertBulk) UpdateStartTime() *PerspectiveUpsertBulk {
	return u.Update(func(s *PerspectiveUpsert) {
		s.UpdateStartTime()
	})
}

// SetEndTime sets the "endTime" field.
func (u *PerspectiveUpsertBulk) SetEndTime(v string) *PerspectiveUpsertBulk {
	return u.Update(func(s *PerspectiveUpsert) {
		s.SetEndTime(v)
	})
}

// UpdateEndTime sets the "endTime" field to the value that was provided on create.
func (u *PerspectiveUpsertBulk) UpdateEndTime() *PerspectiveUpsertBulk {
	return u.Update(func(s *PerspectiveUpsert) {
		s.UpdateEndTime()
	})
}

// SetBuiltin sets the "builtin" field.
func (u *PerspectiveUpsertBulk) SetBuiltin(v bool) *PerspectiveUpsertBulk {
	return u.Update(func(s *PerspectiveUpsert) {
		s.SetBuiltin(v)
	})
}

// UpdateBuiltin sets the "builtin" field to the value that was provided on create.
func (u *PerspectiveUpsertBulk) UpdateBuiltin() *PerspectiveUpsertBulk {
	return u.Update(func(s *PerspectiveUpsert) {
		s.UpdateBuiltin()
	})
}

// SetAllocationQueries sets the "allocationQueries" field.
func (u *PerspectiveUpsertBulk) SetAllocationQueries(v []types.QueryCondition) *PerspectiveUpsertBulk {
	return u.Update(func(s *PerspectiveUpsert) {
		s.SetAllocationQueries(v)
	})
}

// UpdateAllocationQueries sets the "allocationQueries" field to the value that was provided on create.
func (u *PerspectiveUpsertBulk) UpdateAllocationQueries() *PerspectiveUpsertBulk {
	return u.Update(func(s *PerspectiveUpsert) {
		s.UpdateAllocationQueries()
	})
}

// Exec executes the query.
func (u *PerspectiveUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the PerspectiveCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for PerspectiveCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PerspectiveUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
