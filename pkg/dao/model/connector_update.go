// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/allocationcost"
	"github.com/seal-io/seal/pkg/dao/model/applicationresource"
	"github.com/seal-io/seal/pkg/dao/model/clustercost"
	"github.com/seal-io/seal/pkg/dao/model/connector"
	"github.com/seal-io/seal/pkg/dao/model/environment"
	"github.com/seal-io/seal/pkg/dao/model/environmentconnectorrelationship"
	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
	"github.com/seal-io/seal/pkg/dao/types"
)

// ConnectorUpdate is the builder for updating Connector entities.
type ConnectorUpdate struct {
	config
	hooks     []Hook
	mutation  *ConnectorMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ConnectorUpdate builder.
func (cu *ConnectorUpdate) Where(ps ...predicate.Connector) *ConnectorUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *ConnectorUpdate) SetName(s string) *ConnectorUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetDescription sets the "description" field.
func (cu *ConnectorUpdate) SetDescription(s string) *ConnectorUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cu *ConnectorUpdate) SetNillableDescription(s *string) *ConnectorUpdate {
	if s != nil {
		cu.SetDescription(*s)
	}
	return cu
}

// ClearDescription clears the value of the "description" field.
func (cu *ConnectorUpdate) ClearDescription() *ConnectorUpdate {
	cu.mutation.ClearDescription()
	return cu
}

// SetLabels sets the "labels" field.
func (cu *ConnectorUpdate) SetLabels(m map[string]string) *ConnectorUpdate {
	cu.mutation.SetLabels(m)
	return cu
}

// SetStatus sets the "status" field.
func (cu *ConnectorUpdate) SetStatus(s string) *ConnectorUpdate {
	cu.mutation.SetStatus(s)
	return cu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cu *ConnectorUpdate) SetNillableStatus(s *string) *ConnectorUpdate {
	if s != nil {
		cu.SetStatus(*s)
	}
	return cu
}

// ClearStatus clears the value of the "status" field.
func (cu *ConnectorUpdate) ClearStatus() *ConnectorUpdate {
	cu.mutation.ClearStatus()
	return cu
}

// SetStatusMessage sets the "statusMessage" field.
func (cu *ConnectorUpdate) SetStatusMessage(s string) *ConnectorUpdate {
	cu.mutation.SetStatusMessage(s)
	return cu
}

// SetNillableStatusMessage sets the "statusMessage" field if the given value is not nil.
func (cu *ConnectorUpdate) SetNillableStatusMessage(s *string) *ConnectorUpdate {
	if s != nil {
		cu.SetStatusMessage(*s)
	}
	return cu
}

// ClearStatusMessage clears the value of the "statusMessage" field.
func (cu *ConnectorUpdate) ClearStatusMessage() *ConnectorUpdate {
	cu.mutation.ClearStatusMessage()
	return cu
}

// SetUpdateTime sets the "updateTime" field.
func (cu *ConnectorUpdate) SetUpdateTime(t time.Time) *ConnectorUpdate {
	cu.mutation.SetUpdateTime(t)
	return cu
}

// SetConfigVersion sets the "configVersion" field.
func (cu *ConnectorUpdate) SetConfigVersion(s string) *ConnectorUpdate {
	cu.mutation.SetConfigVersion(s)
	return cu
}

// SetConfigData sets the "configData" field.
func (cu *ConnectorUpdate) SetConfigData(m map[string]interface{}) *ConnectorUpdate {
	cu.mutation.SetConfigData(m)
	return cu
}

// SetEnableFinOps sets the "enableFinOps" field.
func (cu *ConnectorUpdate) SetEnableFinOps(b bool) *ConnectorUpdate {
	cu.mutation.SetEnableFinOps(b)
	return cu
}

// SetFinOpsStatus sets the "finOpsStatus" field.
func (cu *ConnectorUpdate) SetFinOpsStatus(s string) *ConnectorUpdate {
	cu.mutation.SetFinOpsStatus(s)
	return cu
}

// SetNillableFinOpsStatus sets the "finOpsStatus" field if the given value is not nil.
func (cu *ConnectorUpdate) SetNillableFinOpsStatus(s *string) *ConnectorUpdate {
	if s != nil {
		cu.SetFinOpsStatus(*s)
	}
	return cu
}

// ClearFinOpsStatus clears the value of the "finOpsStatus" field.
func (cu *ConnectorUpdate) ClearFinOpsStatus() *ConnectorUpdate {
	cu.mutation.ClearFinOpsStatus()
	return cu
}

// SetFinOpsStatusMessage sets the "finOpsStatusMessage" field.
func (cu *ConnectorUpdate) SetFinOpsStatusMessage(s string) *ConnectorUpdate {
	cu.mutation.SetFinOpsStatusMessage(s)
	return cu
}

// SetNillableFinOpsStatusMessage sets the "finOpsStatusMessage" field if the given value is not nil.
func (cu *ConnectorUpdate) SetNillableFinOpsStatusMessage(s *string) *ConnectorUpdate {
	if s != nil {
		cu.SetFinOpsStatusMessage(*s)
	}
	return cu
}

// ClearFinOpsStatusMessage clears the value of the "finOpsStatusMessage" field.
func (cu *ConnectorUpdate) ClearFinOpsStatusMessage() *ConnectorUpdate {
	cu.mutation.ClearFinOpsStatusMessage()
	return cu
}

// AddEnvironmentIDs adds the "environments" edge to the Environment entity by IDs.
func (cu *ConnectorUpdate) AddEnvironmentIDs(ids ...types.ID) *ConnectorUpdate {
	cu.mutation.AddEnvironmentIDs(ids...)
	return cu
}

// AddEnvironments adds the "environments" edges to the Environment entity.
func (cu *ConnectorUpdate) AddEnvironments(e ...*Environment) *ConnectorUpdate {
	ids := make([]types.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cu.AddEnvironmentIDs(ids...)
}

// AddResourceIDs adds the "resources" edge to the ApplicationResource entity by IDs.
func (cu *ConnectorUpdate) AddResourceIDs(ids ...types.ID) *ConnectorUpdate {
	cu.mutation.AddResourceIDs(ids...)
	return cu
}

// AddResources adds the "resources" edges to the ApplicationResource entity.
func (cu *ConnectorUpdate) AddResources(a ...*ApplicationResource) *ConnectorUpdate {
	ids := make([]types.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cu.AddResourceIDs(ids...)
}

// AddClusterCostIDs adds the "clusterCosts" edge to the ClusterCost entity by IDs.
func (cu *ConnectorUpdate) AddClusterCostIDs(ids ...int) *ConnectorUpdate {
	cu.mutation.AddClusterCostIDs(ids...)
	return cu
}

// AddClusterCosts adds the "clusterCosts" edges to the ClusterCost entity.
func (cu *ConnectorUpdate) AddClusterCosts(c ...*ClusterCost) *ConnectorUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddClusterCostIDs(ids...)
}

// AddAllocationCostIDs adds the "allocationCosts" edge to the AllocationCost entity by IDs.
func (cu *ConnectorUpdate) AddAllocationCostIDs(ids ...int) *ConnectorUpdate {
	cu.mutation.AddAllocationCostIDs(ids...)
	return cu
}

// AddAllocationCosts adds the "allocationCosts" edges to the AllocationCost entity.
func (cu *ConnectorUpdate) AddAllocationCosts(a ...*AllocationCost) *ConnectorUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cu.AddAllocationCostIDs(ids...)
}

// AddEnvironmentConnectorRelationshipIDs adds the "environmentConnectorRelationships" edge to the EnvironmentConnectorRelationship entity by IDs.
func (cu *ConnectorUpdate) AddEnvironmentConnectorRelationshipIDs(ids ...int) *ConnectorUpdate {
	cu.mutation.AddEnvironmentConnectorRelationshipIDs(ids...)
	return cu
}

// AddEnvironmentConnectorRelationships adds the "environmentConnectorRelationships" edges to the EnvironmentConnectorRelationship entity.
func (cu *ConnectorUpdate) AddEnvironmentConnectorRelationships(e ...*EnvironmentConnectorRelationship) *ConnectorUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cu.AddEnvironmentConnectorRelationshipIDs(ids...)
}

// Mutation returns the ConnectorMutation object of the builder.
func (cu *ConnectorUpdate) Mutation() *ConnectorMutation {
	return cu.mutation
}

// ClearEnvironments clears all "environments" edges to the Environment entity.
func (cu *ConnectorUpdate) ClearEnvironments() *ConnectorUpdate {
	cu.mutation.ClearEnvironments()
	return cu
}

// RemoveEnvironmentIDs removes the "environments" edge to Environment entities by IDs.
func (cu *ConnectorUpdate) RemoveEnvironmentIDs(ids ...types.ID) *ConnectorUpdate {
	cu.mutation.RemoveEnvironmentIDs(ids...)
	return cu
}

// RemoveEnvironments removes "environments" edges to Environment entities.
func (cu *ConnectorUpdate) RemoveEnvironments(e ...*Environment) *ConnectorUpdate {
	ids := make([]types.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cu.RemoveEnvironmentIDs(ids...)
}

// ClearResources clears all "resources" edges to the ApplicationResource entity.
func (cu *ConnectorUpdate) ClearResources() *ConnectorUpdate {
	cu.mutation.ClearResources()
	return cu
}

// RemoveResourceIDs removes the "resources" edge to ApplicationResource entities by IDs.
func (cu *ConnectorUpdate) RemoveResourceIDs(ids ...types.ID) *ConnectorUpdate {
	cu.mutation.RemoveResourceIDs(ids...)
	return cu
}

// RemoveResources removes "resources" edges to ApplicationResource entities.
func (cu *ConnectorUpdate) RemoveResources(a ...*ApplicationResource) *ConnectorUpdate {
	ids := make([]types.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cu.RemoveResourceIDs(ids...)
}

// ClearClusterCosts clears all "clusterCosts" edges to the ClusterCost entity.
func (cu *ConnectorUpdate) ClearClusterCosts() *ConnectorUpdate {
	cu.mutation.ClearClusterCosts()
	return cu
}

// RemoveClusterCostIDs removes the "clusterCosts" edge to ClusterCost entities by IDs.
func (cu *ConnectorUpdate) RemoveClusterCostIDs(ids ...int) *ConnectorUpdate {
	cu.mutation.RemoveClusterCostIDs(ids...)
	return cu
}

// RemoveClusterCosts removes "clusterCosts" edges to ClusterCost entities.
func (cu *ConnectorUpdate) RemoveClusterCosts(c ...*ClusterCost) *ConnectorUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveClusterCostIDs(ids...)
}

// ClearAllocationCosts clears all "allocationCosts" edges to the AllocationCost entity.
func (cu *ConnectorUpdate) ClearAllocationCosts() *ConnectorUpdate {
	cu.mutation.ClearAllocationCosts()
	return cu
}

// RemoveAllocationCostIDs removes the "allocationCosts" edge to AllocationCost entities by IDs.
func (cu *ConnectorUpdate) RemoveAllocationCostIDs(ids ...int) *ConnectorUpdate {
	cu.mutation.RemoveAllocationCostIDs(ids...)
	return cu
}

// RemoveAllocationCosts removes "allocationCosts" edges to AllocationCost entities.
func (cu *ConnectorUpdate) RemoveAllocationCosts(a ...*AllocationCost) *ConnectorUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cu.RemoveAllocationCostIDs(ids...)
}

// ClearEnvironmentConnectorRelationships clears all "environmentConnectorRelationships" edges to the EnvironmentConnectorRelationship entity.
func (cu *ConnectorUpdate) ClearEnvironmentConnectorRelationships() *ConnectorUpdate {
	cu.mutation.ClearEnvironmentConnectorRelationships()
	return cu
}

// RemoveEnvironmentConnectorRelationshipIDs removes the "environmentConnectorRelationships" edge to EnvironmentConnectorRelationship entities by IDs.
func (cu *ConnectorUpdate) RemoveEnvironmentConnectorRelationshipIDs(ids ...int) *ConnectorUpdate {
	cu.mutation.RemoveEnvironmentConnectorRelationshipIDs(ids...)
	return cu
}

// RemoveEnvironmentConnectorRelationships removes "environmentConnectorRelationships" edges to EnvironmentConnectorRelationship entities.
func (cu *ConnectorUpdate) RemoveEnvironmentConnectorRelationships(e ...*EnvironmentConnectorRelationship) *ConnectorUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cu.RemoveEnvironmentConnectorRelationshipIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ConnectorUpdate) Save(ctx context.Context) (int, error) {
	if err := cu.defaults(); err != nil {
		return 0, err
	}
	return withHooks[int, ConnectorMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ConnectorUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ConnectorUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ConnectorUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ConnectorUpdate) defaults() error {
	if _, ok := cu.mutation.UpdateTime(); !ok {
		if connector.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized connector.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := connector.UpdateDefaultUpdateTime()
		cu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cu *ConnectorUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := connector.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`model: validator failed for field "Connector.name": %w`, err)}
		}
	}
	if v, ok := cu.mutation.ConfigVersion(); ok {
		if err := connector.ConfigVersionValidator(v); err != nil {
			return &ValidationError{Name: "configVersion", err: fmt.Errorf(`model: validator failed for field "Connector.configVersion": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *ConnectorUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ConnectorUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *ConnectorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   connector.Table,
			Columns: connector.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: connector.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(connector.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(connector.FieldDescription, field.TypeString, value)
	}
	if cu.mutation.DescriptionCleared() {
		_spec.ClearField(connector.FieldDescription, field.TypeString)
	}
	if value, ok := cu.mutation.Labels(); ok {
		_spec.SetField(connector.FieldLabels, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.SetField(connector.FieldStatus, field.TypeString, value)
	}
	if cu.mutation.StatusCleared() {
		_spec.ClearField(connector.FieldStatus, field.TypeString)
	}
	if value, ok := cu.mutation.StatusMessage(); ok {
		_spec.SetField(connector.FieldStatusMessage, field.TypeString, value)
	}
	if cu.mutation.StatusMessageCleared() {
		_spec.ClearField(connector.FieldStatusMessage, field.TypeString)
	}
	if value, ok := cu.mutation.UpdateTime(); ok {
		_spec.SetField(connector.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := cu.mutation.ConfigVersion(); ok {
		_spec.SetField(connector.FieldConfigVersion, field.TypeString, value)
	}
	if value, ok := cu.mutation.ConfigData(); ok {
		_spec.SetField(connector.FieldConfigData, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.EnableFinOps(); ok {
		_spec.SetField(connector.FieldEnableFinOps, field.TypeBool, value)
	}
	if value, ok := cu.mutation.FinOpsStatus(); ok {
		_spec.SetField(connector.FieldFinOpsStatus, field.TypeString, value)
	}
	if cu.mutation.FinOpsStatusCleared() {
		_spec.ClearField(connector.FieldFinOpsStatus, field.TypeString)
	}
	if value, ok := cu.mutation.FinOpsStatusMessage(); ok {
		_spec.SetField(connector.FieldFinOpsStatusMessage, field.TypeString, value)
	}
	if cu.mutation.FinOpsStatusMessageCleared() {
		_spec.ClearField(connector.FieldFinOpsStatusMessage, field.TypeString)
	}
	if cu.mutation.EnvironmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   connector.EnvironmentsTable,
			Columns: connector.EnvironmentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: environment.FieldID,
				},
			},
		}
		edge.Schema = cu.schemaConfig.EnvironmentConnectorRelationship
		createE := &EnvironmentConnectorRelationshipCreate{config: cu.config, mutation: newEnvironmentConnectorRelationshipMutation(cu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedEnvironmentsIDs(); len(nodes) > 0 && !cu.mutation.EnvironmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   connector.EnvironmentsTable,
			Columns: connector.EnvironmentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: environment.FieldID,
				},
			},
		}
		edge.Schema = cu.schemaConfig.EnvironmentConnectorRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &EnvironmentConnectorRelationshipCreate{config: cu.config, mutation: newEnvironmentConnectorRelationshipMutation(cu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.EnvironmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   connector.EnvironmentsTable,
			Columns: connector.EnvironmentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: environment.FieldID,
				},
			},
		}
		edge.Schema = cu.schemaConfig.EnvironmentConnectorRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &EnvironmentConnectorRelationshipCreate{config: cu.config, mutation: newEnvironmentConnectorRelationshipMutation(cu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ResourcesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.ResourcesTable,
			Columns: []string{connector.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationresource.FieldID,
				},
			},
		}
		edge.Schema = cu.schemaConfig.ApplicationResource
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedResourcesIDs(); len(nodes) > 0 && !cu.mutation.ResourcesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.ResourcesTable,
			Columns: []string{connector.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationresource.FieldID,
				},
			},
		}
		edge.Schema = cu.schemaConfig.ApplicationResource
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ResourcesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.ResourcesTable,
			Columns: []string{connector.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationresource.FieldID,
				},
			},
		}
		edge.Schema = cu.schemaConfig.ApplicationResource
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ClusterCostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.ClusterCostsTable,
			Columns: []string{connector.ClusterCostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clustercost.FieldID,
				},
			},
		}
		edge.Schema = cu.schemaConfig.ClusterCost
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedClusterCostsIDs(); len(nodes) > 0 && !cu.mutation.ClusterCostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.ClusterCostsTable,
			Columns: []string{connector.ClusterCostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clustercost.FieldID,
				},
			},
		}
		edge.Schema = cu.schemaConfig.ClusterCost
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ClusterCostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.ClusterCostsTable,
			Columns: []string{connector.ClusterCostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clustercost.FieldID,
				},
			},
		}
		edge.Schema = cu.schemaConfig.ClusterCost
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.AllocationCostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.AllocationCostsTable,
			Columns: []string{connector.AllocationCostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: allocationcost.FieldID,
				},
			},
		}
		edge.Schema = cu.schemaConfig.AllocationCost
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedAllocationCostsIDs(); len(nodes) > 0 && !cu.mutation.AllocationCostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.AllocationCostsTable,
			Columns: []string{connector.AllocationCostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: allocationcost.FieldID,
				},
			},
		}
		edge.Schema = cu.schemaConfig.AllocationCost
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.AllocationCostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.AllocationCostsTable,
			Columns: []string{connector.AllocationCostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: allocationcost.FieldID,
				},
			},
		}
		edge.Schema = cu.schemaConfig.AllocationCost
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.EnvironmentConnectorRelationshipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   connector.EnvironmentConnectorRelationshipsTable,
			Columns: []string{connector.EnvironmentConnectorRelationshipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: environmentconnectorrelationship.FieldID,
				},
			},
		}
		edge.Schema = cu.schemaConfig.EnvironmentConnectorRelationship
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedEnvironmentConnectorRelationshipsIDs(); len(nodes) > 0 && !cu.mutation.EnvironmentConnectorRelationshipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   connector.EnvironmentConnectorRelationshipsTable,
			Columns: []string{connector.EnvironmentConnectorRelationshipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: environmentconnectorrelationship.FieldID,
				},
			},
		}
		edge.Schema = cu.schemaConfig.EnvironmentConnectorRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.EnvironmentConnectorRelationshipsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   connector.EnvironmentConnectorRelationshipsTable,
			Columns: []string{connector.EnvironmentConnectorRelationshipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: environmentconnectorrelationship.FieldID,
				},
			},
		}
		edge.Schema = cu.schemaConfig.EnvironmentConnectorRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = cu.schemaConfig.Connector
	ctx = internal.NewSchemaConfigContext(ctx, cu.schemaConfig)
	_spec.AddModifiers(cu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{connector.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ConnectorUpdateOne is the builder for updating a single Connector entity.
type ConnectorUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ConnectorMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetName sets the "name" field.
func (cuo *ConnectorUpdateOne) SetName(s string) *ConnectorUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *ConnectorUpdateOne) SetDescription(s string) *ConnectorUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cuo *ConnectorUpdateOne) SetNillableDescription(s *string) *ConnectorUpdateOne {
	if s != nil {
		cuo.SetDescription(*s)
	}
	return cuo
}

// ClearDescription clears the value of the "description" field.
func (cuo *ConnectorUpdateOne) ClearDescription() *ConnectorUpdateOne {
	cuo.mutation.ClearDescription()
	return cuo
}

// SetLabels sets the "labels" field.
func (cuo *ConnectorUpdateOne) SetLabels(m map[string]string) *ConnectorUpdateOne {
	cuo.mutation.SetLabels(m)
	return cuo
}

// SetStatus sets the "status" field.
func (cuo *ConnectorUpdateOne) SetStatus(s string) *ConnectorUpdateOne {
	cuo.mutation.SetStatus(s)
	return cuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cuo *ConnectorUpdateOne) SetNillableStatus(s *string) *ConnectorUpdateOne {
	if s != nil {
		cuo.SetStatus(*s)
	}
	return cuo
}

// ClearStatus clears the value of the "status" field.
func (cuo *ConnectorUpdateOne) ClearStatus() *ConnectorUpdateOne {
	cuo.mutation.ClearStatus()
	return cuo
}

// SetStatusMessage sets the "statusMessage" field.
func (cuo *ConnectorUpdateOne) SetStatusMessage(s string) *ConnectorUpdateOne {
	cuo.mutation.SetStatusMessage(s)
	return cuo
}

// SetNillableStatusMessage sets the "statusMessage" field if the given value is not nil.
func (cuo *ConnectorUpdateOne) SetNillableStatusMessage(s *string) *ConnectorUpdateOne {
	if s != nil {
		cuo.SetStatusMessage(*s)
	}
	return cuo
}

// ClearStatusMessage clears the value of the "statusMessage" field.
func (cuo *ConnectorUpdateOne) ClearStatusMessage() *ConnectorUpdateOne {
	cuo.mutation.ClearStatusMessage()
	return cuo
}

// SetUpdateTime sets the "updateTime" field.
func (cuo *ConnectorUpdateOne) SetUpdateTime(t time.Time) *ConnectorUpdateOne {
	cuo.mutation.SetUpdateTime(t)
	return cuo
}

// SetConfigVersion sets the "configVersion" field.
func (cuo *ConnectorUpdateOne) SetConfigVersion(s string) *ConnectorUpdateOne {
	cuo.mutation.SetConfigVersion(s)
	return cuo
}

// SetConfigData sets the "configData" field.
func (cuo *ConnectorUpdateOne) SetConfigData(m map[string]interface{}) *ConnectorUpdateOne {
	cuo.mutation.SetConfigData(m)
	return cuo
}

// SetEnableFinOps sets the "enableFinOps" field.
func (cuo *ConnectorUpdateOne) SetEnableFinOps(b bool) *ConnectorUpdateOne {
	cuo.mutation.SetEnableFinOps(b)
	return cuo
}

// SetFinOpsStatus sets the "finOpsStatus" field.
func (cuo *ConnectorUpdateOne) SetFinOpsStatus(s string) *ConnectorUpdateOne {
	cuo.mutation.SetFinOpsStatus(s)
	return cuo
}

// SetNillableFinOpsStatus sets the "finOpsStatus" field if the given value is not nil.
func (cuo *ConnectorUpdateOne) SetNillableFinOpsStatus(s *string) *ConnectorUpdateOne {
	if s != nil {
		cuo.SetFinOpsStatus(*s)
	}
	return cuo
}

// ClearFinOpsStatus clears the value of the "finOpsStatus" field.
func (cuo *ConnectorUpdateOne) ClearFinOpsStatus() *ConnectorUpdateOne {
	cuo.mutation.ClearFinOpsStatus()
	return cuo
}

// SetFinOpsStatusMessage sets the "finOpsStatusMessage" field.
func (cuo *ConnectorUpdateOne) SetFinOpsStatusMessage(s string) *ConnectorUpdateOne {
	cuo.mutation.SetFinOpsStatusMessage(s)
	return cuo
}

// SetNillableFinOpsStatusMessage sets the "finOpsStatusMessage" field if the given value is not nil.
func (cuo *ConnectorUpdateOne) SetNillableFinOpsStatusMessage(s *string) *ConnectorUpdateOne {
	if s != nil {
		cuo.SetFinOpsStatusMessage(*s)
	}
	return cuo
}

// ClearFinOpsStatusMessage clears the value of the "finOpsStatusMessage" field.
func (cuo *ConnectorUpdateOne) ClearFinOpsStatusMessage() *ConnectorUpdateOne {
	cuo.mutation.ClearFinOpsStatusMessage()
	return cuo
}

// AddEnvironmentIDs adds the "environments" edge to the Environment entity by IDs.
func (cuo *ConnectorUpdateOne) AddEnvironmentIDs(ids ...types.ID) *ConnectorUpdateOne {
	cuo.mutation.AddEnvironmentIDs(ids...)
	return cuo
}

// AddEnvironments adds the "environments" edges to the Environment entity.
func (cuo *ConnectorUpdateOne) AddEnvironments(e ...*Environment) *ConnectorUpdateOne {
	ids := make([]types.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cuo.AddEnvironmentIDs(ids...)
}

// AddResourceIDs adds the "resources" edge to the ApplicationResource entity by IDs.
func (cuo *ConnectorUpdateOne) AddResourceIDs(ids ...types.ID) *ConnectorUpdateOne {
	cuo.mutation.AddResourceIDs(ids...)
	return cuo
}

// AddResources adds the "resources" edges to the ApplicationResource entity.
func (cuo *ConnectorUpdateOne) AddResources(a ...*ApplicationResource) *ConnectorUpdateOne {
	ids := make([]types.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cuo.AddResourceIDs(ids...)
}

// AddClusterCostIDs adds the "clusterCosts" edge to the ClusterCost entity by IDs.
func (cuo *ConnectorUpdateOne) AddClusterCostIDs(ids ...int) *ConnectorUpdateOne {
	cuo.mutation.AddClusterCostIDs(ids...)
	return cuo
}

// AddClusterCosts adds the "clusterCosts" edges to the ClusterCost entity.
func (cuo *ConnectorUpdateOne) AddClusterCosts(c ...*ClusterCost) *ConnectorUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddClusterCostIDs(ids...)
}

// AddAllocationCostIDs adds the "allocationCosts" edge to the AllocationCost entity by IDs.
func (cuo *ConnectorUpdateOne) AddAllocationCostIDs(ids ...int) *ConnectorUpdateOne {
	cuo.mutation.AddAllocationCostIDs(ids...)
	return cuo
}

// AddAllocationCosts adds the "allocationCosts" edges to the AllocationCost entity.
func (cuo *ConnectorUpdateOne) AddAllocationCosts(a ...*AllocationCost) *ConnectorUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cuo.AddAllocationCostIDs(ids...)
}

// AddEnvironmentConnectorRelationshipIDs adds the "environmentConnectorRelationships" edge to the EnvironmentConnectorRelationship entity by IDs.
func (cuo *ConnectorUpdateOne) AddEnvironmentConnectorRelationshipIDs(ids ...int) *ConnectorUpdateOne {
	cuo.mutation.AddEnvironmentConnectorRelationshipIDs(ids...)
	return cuo
}

// AddEnvironmentConnectorRelationships adds the "environmentConnectorRelationships" edges to the EnvironmentConnectorRelationship entity.
func (cuo *ConnectorUpdateOne) AddEnvironmentConnectorRelationships(e ...*EnvironmentConnectorRelationship) *ConnectorUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cuo.AddEnvironmentConnectorRelationshipIDs(ids...)
}

// Mutation returns the ConnectorMutation object of the builder.
func (cuo *ConnectorUpdateOne) Mutation() *ConnectorMutation {
	return cuo.mutation
}

// ClearEnvironments clears all "environments" edges to the Environment entity.
func (cuo *ConnectorUpdateOne) ClearEnvironments() *ConnectorUpdateOne {
	cuo.mutation.ClearEnvironments()
	return cuo
}

// RemoveEnvironmentIDs removes the "environments" edge to Environment entities by IDs.
func (cuo *ConnectorUpdateOne) RemoveEnvironmentIDs(ids ...types.ID) *ConnectorUpdateOne {
	cuo.mutation.RemoveEnvironmentIDs(ids...)
	return cuo
}

// RemoveEnvironments removes "environments" edges to Environment entities.
func (cuo *ConnectorUpdateOne) RemoveEnvironments(e ...*Environment) *ConnectorUpdateOne {
	ids := make([]types.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cuo.RemoveEnvironmentIDs(ids...)
}

// ClearResources clears all "resources" edges to the ApplicationResource entity.
func (cuo *ConnectorUpdateOne) ClearResources() *ConnectorUpdateOne {
	cuo.mutation.ClearResources()
	return cuo
}

// RemoveResourceIDs removes the "resources" edge to ApplicationResource entities by IDs.
func (cuo *ConnectorUpdateOne) RemoveResourceIDs(ids ...types.ID) *ConnectorUpdateOne {
	cuo.mutation.RemoveResourceIDs(ids...)
	return cuo
}

// RemoveResources removes "resources" edges to ApplicationResource entities.
func (cuo *ConnectorUpdateOne) RemoveResources(a ...*ApplicationResource) *ConnectorUpdateOne {
	ids := make([]types.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cuo.RemoveResourceIDs(ids...)
}

// ClearClusterCosts clears all "clusterCosts" edges to the ClusterCost entity.
func (cuo *ConnectorUpdateOne) ClearClusterCosts() *ConnectorUpdateOne {
	cuo.mutation.ClearClusterCosts()
	return cuo
}

// RemoveClusterCostIDs removes the "clusterCosts" edge to ClusterCost entities by IDs.
func (cuo *ConnectorUpdateOne) RemoveClusterCostIDs(ids ...int) *ConnectorUpdateOne {
	cuo.mutation.RemoveClusterCostIDs(ids...)
	return cuo
}

// RemoveClusterCosts removes "clusterCosts" edges to ClusterCost entities.
func (cuo *ConnectorUpdateOne) RemoveClusterCosts(c ...*ClusterCost) *ConnectorUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveClusterCostIDs(ids...)
}

// ClearAllocationCosts clears all "allocationCosts" edges to the AllocationCost entity.
func (cuo *ConnectorUpdateOne) ClearAllocationCosts() *ConnectorUpdateOne {
	cuo.mutation.ClearAllocationCosts()
	return cuo
}

// RemoveAllocationCostIDs removes the "allocationCosts" edge to AllocationCost entities by IDs.
func (cuo *ConnectorUpdateOne) RemoveAllocationCostIDs(ids ...int) *ConnectorUpdateOne {
	cuo.mutation.RemoveAllocationCostIDs(ids...)
	return cuo
}

// RemoveAllocationCosts removes "allocationCosts" edges to AllocationCost entities.
func (cuo *ConnectorUpdateOne) RemoveAllocationCosts(a ...*AllocationCost) *ConnectorUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cuo.RemoveAllocationCostIDs(ids...)
}

// ClearEnvironmentConnectorRelationships clears all "environmentConnectorRelationships" edges to the EnvironmentConnectorRelationship entity.
func (cuo *ConnectorUpdateOne) ClearEnvironmentConnectorRelationships() *ConnectorUpdateOne {
	cuo.mutation.ClearEnvironmentConnectorRelationships()
	return cuo
}

// RemoveEnvironmentConnectorRelationshipIDs removes the "environmentConnectorRelationships" edge to EnvironmentConnectorRelationship entities by IDs.
func (cuo *ConnectorUpdateOne) RemoveEnvironmentConnectorRelationshipIDs(ids ...int) *ConnectorUpdateOne {
	cuo.mutation.RemoveEnvironmentConnectorRelationshipIDs(ids...)
	return cuo
}

// RemoveEnvironmentConnectorRelationships removes "environmentConnectorRelationships" edges to EnvironmentConnectorRelationship entities.
func (cuo *ConnectorUpdateOne) RemoveEnvironmentConnectorRelationships(e ...*EnvironmentConnectorRelationship) *ConnectorUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cuo.RemoveEnvironmentConnectorRelationshipIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ConnectorUpdateOne) Select(field string, fields ...string) *ConnectorUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Connector entity.
func (cuo *ConnectorUpdateOne) Save(ctx context.Context) (*Connector, error) {
	if err := cuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*Connector, ConnectorMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ConnectorUpdateOne) SaveX(ctx context.Context) *Connector {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ConnectorUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ConnectorUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ConnectorUpdateOne) defaults() error {
	if _, ok := cuo.mutation.UpdateTime(); !ok {
		if connector.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized connector.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := connector.UpdateDefaultUpdateTime()
		cuo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ConnectorUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := connector.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`model: validator failed for field "Connector.name": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.ConfigVersion(); ok {
		if err := connector.ConfigVersionValidator(v); err != nil {
			return &ValidationError{Name: "configVersion", err: fmt.Errorf(`model: validator failed for field "Connector.configVersion": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *ConnectorUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ConnectorUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *ConnectorUpdateOne) sqlSave(ctx context.Context) (_node *Connector, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   connector.Table,
			Columns: connector.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: connector.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "Connector.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, connector.FieldID)
		for _, f := range fields {
			if !connector.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != connector.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(connector.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(connector.FieldDescription, field.TypeString, value)
	}
	if cuo.mutation.DescriptionCleared() {
		_spec.ClearField(connector.FieldDescription, field.TypeString)
	}
	if value, ok := cuo.mutation.Labels(); ok {
		_spec.SetField(connector.FieldLabels, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.SetField(connector.FieldStatus, field.TypeString, value)
	}
	if cuo.mutation.StatusCleared() {
		_spec.ClearField(connector.FieldStatus, field.TypeString)
	}
	if value, ok := cuo.mutation.StatusMessage(); ok {
		_spec.SetField(connector.FieldStatusMessage, field.TypeString, value)
	}
	if cuo.mutation.StatusMessageCleared() {
		_spec.ClearField(connector.FieldStatusMessage, field.TypeString)
	}
	if value, ok := cuo.mutation.UpdateTime(); ok {
		_spec.SetField(connector.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.ConfigVersion(); ok {
		_spec.SetField(connector.FieldConfigVersion, field.TypeString, value)
	}
	if value, ok := cuo.mutation.ConfigData(); ok {
		_spec.SetField(connector.FieldConfigData, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.EnableFinOps(); ok {
		_spec.SetField(connector.FieldEnableFinOps, field.TypeBool, value)
	}
	if value, ok := cuo.mutation.FinOpsStatus(); ok {
		_spec.SetField(connector.FieldFinOpsStatus, field.TypeString, value)
	}
	if cuo.mutation.FinOpsStatusCleared() {
		_spec.ClearField(connector.FieldFinOpsStatus, field.TypeString)
	}
	if value, ok := cuo.mutation.FinOpsStatusMessage(); ok {
		_spec.SetField(connector.FieldFinOpsStatusMessage, field.TypeString, value)
	}
	if cuo.mutation.FinOpsStatusMessageCleared() {
		_spec.ClearField(connector.FieldFinOpsStatusMessage, field.TypeString)
	}
	if cuo.mutation.EnvironmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   connector.EnvironmentsTable,
			Columns: connector.EnvironmentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: environment.FieldID,
				},
			},
		}
		edge.Schema = cuo.schemaConfig.EnvironmentConnectorRelationship
		createE := &EnvironmentConnectorRelationshipCreate{config: cuo.config, mutation: newEnvironmentConnectorRelationshipMutation(cuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedEnvironmentsIDs(); len(nodes) > 0 && !cuo.mutation.EnvironmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   connector.EnvironmentsTable,
			Columns: connector.EnvironmentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: environment.FieldID,
				},
			},
		}
		edge.Schema = cuo.schemaConfig.EnvironmentConnectorRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &EnvironmentConnectorRelationshipCreate{config: cuo.config, mutation: newEnvironmentConnectorRelationshipMutation(cuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.EnvironmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   connector.EnvironmentsTable,
			Columns: connector.EnvironmentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: environment.FieldID,
				},
			},
		}
		edge.Schema = cuo.schemaConfig.EnvironmentConnectorRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &EnvironmentConnectorRelationshipCreate{config: cuo.config, mutation: newEnvironmentConnectorRelationshipMutation(cuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ResourcesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.ResourcesTable,
			Columns: []string{connector.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationresource.FieldID,
				},
			},
		}
		edge.Schema = cuo.schemaConfig.ApplicationResource
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedResourcesIDs(); len(nodes) > 0 && !cuo.mutation.ResourcesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.ResourcesTable,
			Columns: []string{connector.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationresource.FieldID,
				},
			},
		}
		edge.Schema = cuo.schemaConfig.ApplicationResource
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ResourcesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.ResourcesTable,
			Columns: []string{connector.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationresource.FieldID,
				},
			},
		}
		edge.Schema = cuo.schemaConfig.ApplicationResource
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ClusterCostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.ClusterCostsTable,
			Columns: []string{connector.ClusterCostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clustercost.FieldID,
				},
			},
		}
		edge.Schema = cuo.schemaConfig.ClusterCost
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedClusterCostsIDs(); len(nodes) > 0 && !cuo.mutation.ClusterCostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.ClusterCostsTable,
			Columns: []string{connector.ClusterCostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clustercost.FieldID,
				},
			},
		}
		edge.Schema = cuo.schemaConfig.ClusterCost
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ClusterCostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.ClusterCostsTable,
			Columns: []string{connector.ClusterCostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clustercost.FieldID,
				},
			},
		}
		edge.Schema = cuo.schemaConfig.ClusterCost
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.AllocationCostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.AllocationCostsTable,
			Columns: []string{connector.AllocationCostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: allocationcost.FieldID,
				},
			},
		}
		edge.Schema = cuo.schemaConfig.AllocationCost
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedAllocationCostsIDs(); len(nodes) > 0 && !cuo.mutation.AllocationCostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.AllocationCostsTable,
			Columns: []string{connector.AllocationCostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: allocationcost.FieldID,
				},
			},
		}
		edge.Schema = cuo.schemaConfig.AllocationCost
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.AllocationCostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.AllocationCostsTable,
			Columns: []string{connector.AllocationCostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: allocationcost.FieldID,
				},
			},
		}
		edge.Schema = cuo.schemaConfig.AllocationCost
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.EnvironmentConnectorRelationshipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   connector.EnvironmentConnectorRelationshipsTable,
			Columns: []string{connector.EnvironmentConnectorRelationshipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: environmentconnectorrelationship.FieldID,
				},
			},
		}
		edge.Schema = cuo.schemaConfig.EnvironmentConnectorRelationship
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedEnvironmentConnectorRelationshipsIDs(); len(nodes) > 0 && !cuo.mutation.EnvironmentConnectorRelationshipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   connector.EnvironmentConnectorRelationshipsTable,
			Columns: []string{connector.EnvironmentConnectorRelationshipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: environmentconnectorrelationship.FieldID,
				},
			},
		}
		edge.Schema = cuo.schemaConfig.EnvironmentConnectorRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.EnvironmentConnectorRelationshipsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   connector.EnvironmentConnectorRelationshipsTable,
			Columns: []string{connector.EnvironmentConnectorRelationshipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: environmentconnectorrelationship.FieldID,
				},
			},
		}
		edge.Schema = cuo.schemaConfig.EnvironmentConnectorRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = cuo.schemaConfig.Connector
	ctx = internal.NewSchemaConfigContext(ctx, cuo.schemaConfig)
	_spec.AddModifiers(cuo.modifiers...)
	_node = &Connector{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{connector.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
