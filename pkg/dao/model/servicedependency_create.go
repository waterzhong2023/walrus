// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "ent". DO NOT EDIT.

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

	"github.com/seal-io/seal/pkg/dao/model/service"
	"github.com/seal-io/seal/pkg/dao/model/servicedependency"
	"github.com/seal-io/seal/pkg/dao/types/oid"
)

// ServiceDependencyCreate is the builder for creating a ServiceDependency entity.
type ServiceDependencyCreate struct {
	config
	mutation *ServiceDependencyMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "createTime" field.
func (sdc *ServiceDependencyCreate) SetCreateTime(t time.Time) *ServiceDependencyCreate {
	sdc.mutation.SetCreateTime(t)
	return sdc
}

// SetNillableCreateTime sets the "createTime" field if the given value is not nil.
func (sdc *ServiceDependencyCreate) SetNillableCreateTime(t *time.Time) *ServiceDependencyCreate {
	if t != nil {
		sdc.SetCreateTime(*t)
	}
	return sdc
}

// SetServiceID sets the "serviceID" field.
func (sdc *ServiceDependencyCreate) SetServiceID(o oid.ID) *ServiceDependencyCreate {
	sdc.mutation.SetServiceID(o)
	return sdc
}

// SetDependentID sets the "dependentID" field.
func (sdc *ServiceDependencyCreate) SetDependentID(o oid.ID) *ServiceDependencyCreate {
	sdc.mutation.SetDependentID(o)
	return sdc
}

// SetPath sets the "path" field.
func (sdc *ServiceDependencyCreate) SetPath(o []oid.ID) *ServiceDependencyCreate {
	sdc.mutation.SetPath(o)
	return sdc
}

// SetType sets the "type" field.
func (sdc *ServiceDependencyCreate) SetType(s string) *ServiceDependencyCreate {
	sdc.mutation.SetType(s)
	return sdc
}

// SetID sets the "id" field.
func (sdc *ServiceDependencyCreate) SetID(o oid.ID) *ServiceDependencyCreate {
	sdc.mutation.SetID(o)
	return sdc
}

// SetService sets the "service" edge to the Service entity.
func (sdc *ServiceDependencyCreate) SetService(s *Service) *ServiceDependencyCreate {
	return sdc.SetServiceID(s.ID)
}

// Mutation returns the ServiceDependencyMutation object of the builder.
func (sdc *ServiceDependencyCreate) Mutation() *ServiceDependencyMutation {
	return sdc.mutation
}

// Save creates the ServiceDependency in the database.
func (sdc *ServiceDependencyCreate) Save(ctx context.Context) (*ServiceDependency, error) {
	if err := sdc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*ServiceDependency, ServiceDependencyMutation](ctx, sdc.sqlSave, sdc.mutation, sdc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sdc *ServiceDependencyCreate) SaveX(ctx context.Context) *ServiceDependency {
	v, err := sdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdc *ServiceDependencyCreate) Exec(ctx context.Context) error {
	_, err := sdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdc *ServiceDependencyCreate) ExecX(ctx context.Context) {
	if err := sdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sdc *ServiceDependencyCreate) defaults() error {
	if _, ok := sdc.mutation.CreateTime(); !ok {
		if servicedependency.DefaultCreateTime == nil {
			return fmt.Errorf("model: uninitialized servicedependency.DefaultCreateTime (forgotten import model/runtime?)")
		}
		v := servicedependency.DefaultCreateTime()
		sdc.mutation.SetCreateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sdc *ServiceDependencyCreate) check() error {
	if _, ok := sdc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "createTime", err: errors.New(`model: missing required field "ServiceDependency.createTime"`)}
	}
	if _, ok := sdc.mutation.ServiceID(); !ok {
		return &ValidationError{Name: "serviceID", err: errors.New(`model: missing required field "ServiceDependency.serviceID"`)}
	}
	if v, ok := sdc.mutation.ServiceID(); ok {
		if err := servicedependency.ServiceIDValidator(string(v)); err != nil {
			return &ValidationError{Name: "serviceID", err: fmt.Errorf(`model: validator failed for field "ServiceDependency.serviceID": %w`, err)}
		}
	}
	if _, ok := sdc.mutation.DependentID(); !ok {
		return &ValidationError{Name: "dependentID", err: errors.New(`model: missing required field "ServiceDependency.dependentID"`)}
	}
	if v, ok := sdc.mutation.DependentID(); ok {
		if err := servicedependency.DependentIDValidator(string(v)); err != nil {
			return &ValidationError{Name: "dependentID", err: fmt.Errorf(`model: validator failed for field "ServiceDependency.dependentID": %w`, err)}
		}
	}
	if _, ok := sdc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`model: missing required field "ServiceDependency.path"`)}
	}
	if _, ok := sdc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`model: missing required field "ServiceDependency.type"`)}
	}
	if v, ok := sdc.mutation.GetType(); ok {
		if err := servicedependency.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`model: validator failed for field "ServiceDependency.type": %w`, err)}
		}
	}
	if _, ok := sdc.mutation.ServiceID(); !ok {
		return &ValidationError{Name: "service", err: errors.New(`model: missing required edge "ServiceDependency.service"`)}
	}
	return nil
}

func (sdc *ServiceDependencyCreate) sqlSave(ctx context.Context) (*ServiceDependency, error) {
	if err := sdc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*oid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	sdc.mutation.id = &_node.ID
	sdc.mutation.done = true
	return _node, nil
}

func (sdc *ServiceDependencyCreate) createSpec() (*ServiceDependency, *sqlgraph.CreateSpec) {
	var (
		_node = &ServiceDependency{config: sdc.config}
		_spec = sqlgraph.NewCreateSpec(servicedependency.Table, sqlgraph.NewFieldSpec(servicedependency.FieldID, field.TypeString))
	)
	_spec.Schema = sdc.schemaConfig.ServiceDependency
	_spec.OnConflict = sdc.conflict
	if id, ok := sdc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sdc.mutation.CreateTime(); ok {
		_spec.SetField(servicedependency.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = &value
	}
	if value, ok := sdc.mutation.DependentID(); ok {
		_spec.SetField(servicedependency.FieldDependentID, field.TypeString, value)
		_node.DependentID = value
	}
	if value, ok := sdc.mutation.Path(); ok {
		_spec.SetField(servicedependency.FieldPath, field.TypeJSON, value)
		_node.Path = value
	}
	if value, ok := sdc.mutation.GetType(); ok {
		_spec.SetField(servicedependency.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if nodes := sdc.mutation.ServiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   servicedependency.ServiceTable,
			Columns: []string{servicedependency.ServiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeString),
			},
		}
		edge.Schema = sdc.schemaConfig.ServiceDependency
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ServiceID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ServiceDependency.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ServiceDependencyUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (sdc *ServiceDependencyCreate) OnConflict(opts ...sql.ConflictOption) *ServiceDependencyUpsertOne {
	sdc.conflict = opts
	return &ServiceDependencyUpsertOne{
		create: sdc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ServiceDependency.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sdc *ServiceDependencyCreate) OnConflictColumns(columns ...string) *ServiceDependencyUpsertOne {
	sdc.conflict = append(sdc.conflict, sql.ConflictColumns(columns...))
	return &ServiceDependencyUpsertOne{
		create: sdc,
	}
}

type (
	// ServiceDependencyUpsertOne is the builder for "upsert"-ing
	//  one ServiceDependency node.
	ServiceDependencyUpsertOne struct {
		create *ServiceDependencyCreate
	}

	// ServiceDependencyUpsert is the "OnConflict" setter.
	ServiceDependencyUpsert struct {
		*sql.UpdateSet
	}
)

// SetDependentID sets the "dependentID" field.
func (u *ServiceDependencyUpsert) SetDependentID(v oid.ID) *ServiceDependencyUpsert {
	u.Set(servicedependency.FieldDependentID, v)
	return u
}

// UpdateDependentID sets the "dependentID" field to the value that was provided on create.
func (u *ServiceDependencyUpsert) UpdateDependentID() *ServiceDependencyUpsert {
	u.SetExcluded(servicedependency.FieldDependentID)
	return u
}

// SetPath sets the "path" field.
func (u *ServiceDependencyUpsert) SetPath(v []oid.ID) *ServiceDependencyUpsert {
	u.Set(servicedependency.FieldPath, v)
	return u
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *ServiceDependencyUpsert) UpdatePath() *ServiceDependencyUpsert {
	u.SetExcluded(servicedependency.FieldPath)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ServiceDependency.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(servicedependency.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ServiceDependencyUpsertOne) UpdateNewValues() *ServiceDependencyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(servicedependency.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(servicedependency.FieldCreateTime)
		}
		if _, exists := u.create.mutation.ServiceID(); exists {
			s.SetIgnore(servicedependency.FieldServiceID)
		}
		if _, exists := u.create.mutation.GetType(); exists {
			s.SetIgnore(servicedependency.FieldType)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ServiceDependency.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ServiceDependencyUpsertOne) Ignore() *ServiceDependencyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ServiceDependencyUpsertOne) DoNothing() *ServiceDependencyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ServiceDependencyCreate.OnConflict
// documentation for more info.
func (u *ServiceDependencyUpsertOne) Update(set func(*ServiceDependencyUpsert)) *ServiceDependencyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ServiceDependencyUpsert{UpdateSet: update})
	}))
	return u
}

// SetDependentID sets the "dependentID" field.
func (u *ServiceDependencyUpsertOne) SetDependentID(v oid.ID) *ServiceDependencyUpsertOne {
	return u.Update(func(s *ServiceDependencyUpsert) {
		s.SetDependentID(v)
	})
}

// UpdateDependentID sets the "dependentID" field to the value that was provided on create.
func (u *ServiceDependencyUpsertOne) UpdateDependentID() *ServiceDependencyUpsertOne {
	return u.Update(func(s *ServiceDependencyUpsert) {
		s.UpdateDependentID()
	})
}

// SetPath sets the "path" field.
func (u *ServiceDependencyUpsertOne) SetPath(v []oid.ID) *ServiceDependencyUpsertOne {
	return u.Update(func(s *ServiceDependencyUpsert) {
		s.SetPath(v)
	})
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *ServiceDependencyUpsertOne) UpdatePath() *ServiceDependencyUpsertOne {
	return u.Update(func(s *ServiceDependencyUpsert) {
		s.UpdatePath()
	})
}

// Exec executes the query.
func (u *ServiceDependencyUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for ServiceDependencyCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ServiceDependencyUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ServiceDependencyUpsertOne) ID(ctx context.Context) (id oid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("model: ServiceDependencyUpsertOne.ID is not supported by MySQL driver. Use ServiceDependencyUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ServiceDependencyUpsertOne) IDX(ctx context.Context) oid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ServiceDependencyCreateBulk is the builder for creating many ServiceDependency entities in bulk.
type ServiceDependencyCreateBulk struct {
	config
	builders []*ServiceDependencyCreate
	conflict []sql.ConflictOption
}

// Save creates the ServiceDependency entities in the database.
func (sdcb *ServiceDependencyCreateBulk) Save(ctx context.Context) ([]*ServiceDependency, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sdcb.builders))
	nodes := make([]*ServiceDependency, len(sdcb.builders))
	mutators := make([]Mutator, len(sdcb.builders))
	for i := range sdcb.builders {
		func(i int, root context.Context) {
			builder := sdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ServiceDependencyMutation)
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
					_, err = mutators[i+1].Mutate(root, sdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sdcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sdcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sdcb *ServiceDependencyCreateBulk) SaveX(ctx context.Context) []*ServiceDependency {
	v, err := sdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdcb *ServiceDependencyCreateBulk) Exec(ctx context.Context) error {
	_, err := sdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdcb *ServiceDependencyCreateBulk) ExecX(ctx context.Context) {
	if err := sdcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ServiceDependency.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ServiceDependencyUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (sdcb *ServiceDependencyCreateBulk) OnConflict(opts ...sql.ConflictOption) *ServiceDependencyUpsertBulk {
	sdcb.conflict = opts
	return &ServiceDependencyUpsertBulk{
		create: sdcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ServiceDependency.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sdcb *ServiceDependencyCreateBulk) OnConflictColumns(columns ...string) *ServiceDependencyUpsertBulk {
	sdcb.conflict = append(sdcb.conflict, sql.ConflictColumns(columns...))
	return &ServiceDependencyUpsertBulk{
		create: sdcb,
	}
}

// ServiceDependencyUpsertBulk is the builder for "upsert"-ing
// a bulk of ServiceDependency nodes.
type ServiceDependencyUpsertBulk struct {
	create *ServiceDependencyCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ServiceDependency.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(servicedependency.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ServiceDependencyUpsertBulk) UpdateNewValues() *ServiceDependencyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(servicedependency.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(servicedependency.FieldCreateTime)
			}
			if _, exists := b.mutation.ServiceID(); exists {
				s.SetIgnore(servicedependency.FieldServiceID)
			}
			if _, exists := b.mutation.GetType(); exists {
				s.SetIgnore(servicedependency.FieldType)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ServiceDependency.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ServiceDependencyUpsertBulk) Ignore() *ServiceDependencyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ServiceDependencyUpsertBulk) DoNothing() *ServiceDependencyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ServiceDependencyCreateBulk.OnConflict
// documentation for more info.
func (u *ServiceDependencyUpsertBulk) Update(set func(*ServiceDependencyUpsert)) *ServiceDependencyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ServiceDependencyUpsert{UpdateSet: update})
	}))
	return u
}

// SetDependentID sets the "dependentID" field.
func (u *ServiceDependencyUpsertBulk) SetDependentID(v oid.ID) *ServiceDependencyUpsertBulk {
	return u.Update(func(s *ServiceDependencyUpsert) {
		s.SetDependentID(v)
	})
}

// UpdateDependentID sets the "dependentID" field to the value that was provided on create.
func (u *ServiceDependencyUpsertBulk) UpdateDependentID() *ServiceDependencyUpsertBulk {
	return u.Update(func(s *ServiceDependencyUpsert) {
		s.UpdateDependentID()
	})
}

// SetPath sets the "path" field.
func (u *ServiceDependencyUpsertBulk) SetPath(v []oid.ID) *ServiceDependencyUpsertBulk {
	return u.Update(func(s *ServiceDependencyUpsert) {
		s.SetPath(v)
	})
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *ServiceDependencyUpsertBulk) UpdatePath() *ServiceDependencyUpsertBulk {
	return u.Update(func(s *ServiceDependencyUpsert) {
		s.UpdatePath()
	})
}

// Exec executes the query.
func (u *ServiceDependencyUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the ServiceDependencyCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for ServiceDependencyCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ServiceDependencyUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
