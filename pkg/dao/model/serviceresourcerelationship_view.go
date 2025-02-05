// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/serviceresourcerelationship"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// ServiceResourceRelationshipCreateInput holds the creation input of the ServiceResourceRelationship entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type ServiceResourceRelationshipCreateInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Type of the relationship.
	Type string `path:"-" query:"-" json:"type"`

	// Dependency specifies full inserting the new ServiceResource entity of the ServiceResourceRelationship entity.
	Dependency *ServiceResourceQueryInput `uri:"-" query:"-" json:"dependency"`
}

// Model returns the ServiceResourceRelationship entity for creating,
// after validating.
func (srrci *ServiceResourceRelationshipCreateInput) Model() *ServiceResourceRelationship {
	if srrci == nil {
		return nil
	}

	_srr := &ServiceResourceRelationship{
		Type: srrci.Type,
	}

	if srrci.Dependency != nil {
		_srr.DependencyID = srrci.Dependency.ID
	}
	return _srr
}

// Validate checks the ServiceResourceRelationshipCreateInput entity.
func (srrci *ServiceResourceRelationshipCreateInput) Validate() error {
	if srrci == nil {
		return errors.New("nil receiver")
	}

	return srrci.ValidateWith(srrci.inputConfig.Context, srrci.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceResourceRelationshipCreateInput entity with the given context and client set.
func (srrci *ServiceResourceRelationshipCreateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srrci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	if srrci.Dependency != nil {
		if err := srrci.Dependency.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srrci.Dependency = nil
			}
		}
	}

	return nil
}

// ServiceResourceRelationshipCreateInputs holds the creation input item of the ServiceResourceRelationship entities.
type ServiceResourceRelationshipCreateInputsItem struct {
	// Type of the relationship.
	Type string `path:"-" query:"-" json:"type"`

	// Dependency specifies full inserting the new ServiceResource entity.
	Dependency *ServiceResourceQueryInput `uri:"-" query:"-" json:"dependency"`
}

// ValidateWith checks the ServiceResourceRelationshipCreateInputsItem entity with the given context and client set.
func (srrci *ServiceResourceRelationshipCreateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srrci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	if srrci.Dependency != nil {
		if err := srrci.Dependency.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srrci.Dependency = nil
			}
		}
	}

	return nil
}

// ServiceResourceRelationshipCreateInputs holds the creation input of the ServiceResourceRelationship entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ServiceResourceRelationshipCreateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ServiceResourceRelationshipCreateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the ServiceResourceRelationship entities for creating,
// after validating.
func (srrci *ServiceResourceRelationshipCreateInputs) Model() []*ServiceResourceRelationship {
	if srrci == nil || len(srrci.Items) == 0 {
		return nil
	}

	_srrs := make([]*ServiceResourceRelationship, len(srrci.Items))

	for i := range srrci.Items {
		_srr := &ServiceResourceRelationship{
			Type: srrci.Items[i].Type,
		}

		if srrci.Items[i].Dependency != nil {
			_srr.DependencyID = srrci.Items[i].Dependency.ID
		}

		_srrs[i] = _srr
	}

	return _srrs
}

// Validate checks the ServiceResourceRelationshipCreateInputs entity .
func (srrci *ServiceResourceRelationshipCreateInputs) Validate() error {
	if srrci == nil {
		return errors.New("nil receiver")
	}

	return srrci.ValidateWith(srrci.inputConfig.Context, srrci.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceResourceRelationshipCreateInputs entity with the given context and client set.
func (srrci *ServiceResourceRelationshipCreateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srrci == nil {
		return errors.New("nil receiver")
	}

	if len(srrci.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	for i := range srrci.Items {
		if srrci.Items[i] == nil {
			continue
		}

		if err := srrci.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// ServiceResourceRelationshipDeleteInput holds the deletion input of the ServiceResourceRelationship entity,
// please tags with `path:",inline"` if embedding.
type ServiceResourceRelationshipDeleteInput struct {
	ServiceResourceRelationshipQueryInput `path:",inline"`
}

// ServiceResourceRelationshipDeleteInputs holds the deletion input item of the ServiceResourceRelationship entities.
type ServiceResourceRelationshipDeleteInputsItem struct {
	// ID of the ServiceResourceRelationship entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Type of the ServiceResourceRelationship entity, a part of the unique index.
	Type string `path:"-" query:"-" json:"type,omitempty"`
}

// ServiceResourceRelationshipDeleteInputs holds the deletion input of the ServiceResourceRelationship entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ServiceResourceRelationshipDeleteInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ServiceResourceRelationshipDeleteInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the ServiceResourceRelationship entities for deleting,
// after validating.
func (srrdi *ServiceResourceRelationshipDeleteInputs) Model() []*ServiceResourceRelationship {
	if srrdi == nil || len(srrdi.Items) == 0 {
		return nil
	}

	_srrs := make([]*ServiceResourceRelationship, len(srrdi.Items))
	for i := range srrdi.Items {
		_srrs[i] = &ServiceResourceRelationship{
			ID: srrdi.Items[i].ID,
		}
	}
	return _srrs
}

// IDs returns the ID list of the ServiceResourceRelationship entities for deleting,
// after validating.
func (srrdi *ServiceResourceRelationshipDeleteInputs) IDs() []object.ID {
	if srrdi == nil || len(srrdi.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(srrdi.Items))
	for i := range srrdi.Items {
		ids[i] = srrdi.Items[i].ID
	}
	return ids
}

// Validate checks the ServiceResourceRelationshipDeleteInputs entity.
func (srrdi *ServiceResourceRelationshipDeleteInputs) Validate() error {
	if srrdi == nil {
		return errors.New("nil receiver")
	}

	return srrdi.ValidateWith(srrdi.inputConfig.Context, srrdi.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceResourceRelationshipDeleteInputs entity with the given context and client set.
func (srrdi *ServiceResourceRelationshipDeleteInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srrdi == nil {
		return errors.New("nil receiver")
	}

	if len(srrdi.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.ServiceResourceRelationships().Query()

	ids := make([]object.ID, 0, len(srrdi.Items))
	ors := make([]predicate.ServiceResourceRelationship, 0, len(srrdi.Items))
	indexers := make(map[any][]int)

	for i := range srrdi.Items {
		if srrdi.Items[i] == nil {
			return errors.New("nil item")
		}

		if srrdi.Items[i].ID != "" {
			ids = append(ids, srrdi.Items[i].ID)
			ors = append(ors, serviceresourcerelationship.ID(srrdi.Items[i].ID))
			indexers[srrdi.Items[i].ID] = append(indexers[srrdi.Items[i].ID], i)
		} else if srrdi.Items[i].Type != "" {
			ors = append(ors, serviceresourcerelationship.And(
				serviceresourcerelationship.Type(srrdi.Items[i].Type)))
			indexerKey := fmt.Sprint("/", srrdi.Items[i].Type)
			indexers[indexerKey] = append(indexers[indexerKey], i)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	p := serviceresourcerelationship.IDIn(ids...)
	if len(ids) != cap(ids) {
		p = serviceresourcerelationship.Or(ors...)
	}

	es, err := q.
		Where(p).
		Select(
			serviceresourcerelationship.FieldID,
			serviceresourcerelationship.FieldType,
		).
		All(ctx)
	if err != nil {
		return err
	}

	if len(es) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range es {
		indexer := indexers[es[i].ID]
		if indexer == nil {
			indexerKey := fmt.Sprint("/", es[i].Type)
			indexer = indexers[indexerKey]
		}
		for _, j := range indexer {
			srrdi.Items[j].ID = es[i].ID
			srrdi.Items[j].Type = es[i].Type
		}
	}

	return nil
}

// ServiceResourceRelationshipQueryInput holds the query input of the ServiceResourceRelationship entity,
// please tags with `path:",inline"` if embedding.
type ServiceResourceRelationshipQueryInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Refer holds the route path reference of the ServiceResourceRelationship entity.
	Refer *object.Refer `path:"serviceresourcerelationship,default=" query:"-" json:"-"`
	// ID of the ServiceResourceRelationship entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Type of the ServiceResourceRelationship entity, a part of the unique index.
	Type string `path:"-" query:"-" json:"type,omitempty"`
}

// Model returns the ServiceResourceRelationship entity for querying,
// after validating.
func (srrqi *ServiceResourceRelationshipQueryInput) Model() *ServiceResourceRelationship {
	if srrqi == nil {
		return nil
	}

	return &ServiceResourceRelationship{
		ID:   srrqi.ID,
		Type: srrqi.Type,
	}
}

// Validate checks the ServiceResourceRelationshipQueryInput entity.
func (srrqi *ServiceResourceRelationshipQueryInput) Validate() error {
	if srrqi == nil {
		return errors.New("nil receiver")
	}

	return srrqi.ValidateWith(srrqi.inputConfig.Context, srrqi.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceResourceRelationshipQueryInput entity with the given context and client set.
func (srrqi *ServiceResourceRelationshipQueryInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srrqi == nil {
		return errors.New("nil receiver")
	}

	if srrqi.Refer != nil && *srrqi.Refer == "" {
		return fmt.Errorf("model: %s : %w", serviceresourcerelationship.Label, ErrBlankResourceRefer)
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.ServiceResourceRelationships().Query()

	if srrqi.Refer != nil {
		if srrqi.Refer.IsID() {
			q.Where(
				serviceresourcerelationship.ID(srrqi.Refer.ID()))
		} else if refers := srrqi.Refer.Split(1); len(refers) == 1 {
			q.Where(
				serviceresourcerelationship.Type(refers[0].String()))
		} else {
			return errors.New("invalid identify refer of serviceresourcerelationship")
		}
	} else if srrqi.ID != "" {
		q.Where(
			serviceresourcerelationship.ID(srrqi.ID))
	} else if srrqi.Type != "" {
		q.Where(
			serviceresourcerelationship.Type(srrqi.Type))
	} else {
		return errors.New("invalid identify of serviceresourcerelationship")
	}

	q.Select(
		serviceresourcerelationship.FieldID,
		serviceresourcerelationship.FieldType,
	)

	var e *ServiceResourceRelationship
	{
		// Get cache from previous validation.
		queryStmt, queryArgs := q.sqlQuery(setContextOp(ctx, q.ctx, "cache")).Query()
		ck := fmt.Sprintf("stmt=%v, args=%v", queryStmt, queryArgs)
		if cv, existed := cache[ck]; !existed {
			var err error
			e, err = q.Only(ctx)
			if err != nil {
				return err
			}

			// Set cache for other validation.
			cache[ck] = e
		} else {
			e = cv.(*ServiceResourceRelationship)
		}
	}

	srrqi.ID = e.ID
	srrqi.Type = e.Type
	return nil
}

// ServiceResourceRelationshipQueryInputs holds the query input of the ServiceResourceRelationship entities,
// please tags with `path:",inline" query:",inline"` if embedding.
type ServiceResourceRelationshipQueryInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`
}

// Validate checks the ServiceResourceRelationshipQueryInputs entity.
func (srrqi *ServiceResourceRelationshipQueryInputs) Validate() error {
	if srrqi == nil {
		return errors.New("nil receiver")
	}

	return srrqi.ValidateWith(srrqi.inputConfig.Context, srrqi.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceResourceRelationshipQueryInputs entity with the given context and client set.
func (srrqi *ServiceResourceRelationshipQueryInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srrqi == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// ServiceResourceRelationshipUpdateInput holds the modification input of the ServiceResourceRelationship entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type ServiceResourceRelationshipUpdateInput struct {
	ServiceResourceRelationshipQueryInput `path:",inline" query:"-" json:"-"`

	// Dependency indicates replacing the stale ServiceResource entity.
	Dependency *ServiceResourceQueryInput `uri:"-" query:"-" json:"dependency"`
}

// Model returns the ServiceResourceRelationship entity for modifying,
// after validating.
func (srrui *ServiceResourceRelationshipUpdateInput) Model() *ServiceResourceRelationship {
	if srrui == nil {
		return nil
	}

	_srr := &ServiceResourceRelationship{
		ID:   srrui.ID,
		Type: srrui.Type,
	}

	if srrui.Dependency != nil {
		_srr.DependencyID = srrui.Dependency.ID
	}
	return _srr
}

// Validate checks the ServiceResourceRelationshipUpdateInput entity.
func (srrui *ServiceResourceRelationshipUpdateInput) Validate() error {
	if srrui == nil {
		return errors.New("nil receiver")
	}

	return srrui.ValidateWith(srrui.inputConfig.Context, srrui.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceResourceRelationshipUpdateInput entity with the given context and client set.
func (srrui *ServiceResourceRelationshipUpdateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := srrui.ServiceResourceRelationshipQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	if srrui.Dependency != nil {
		if err := srrui.Dependency.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srrui.Dependency = nil
			}
		}
	}

	return nil
}

// ServiceResourceRelationshipUpdateInputs holds the modification input item of the ServiceResourceRelationship entities.
type ServiceResourceRelationshipUpdateInputsItem struct {
	// ID of the ServiceResourceRelationship entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Type of the ServiceResourceRelationship entity, a part of the unique index.
	Type string `path:"-" query:"-" json:"type,omitempty"`

	// Dependency indicates replacing the stale ServiceResource entity.
	Dependency *ServiceResourceQueryInput `uri:"-" query:"-" json:"dependency"`
}

// ValidateWith checks the ServiceResourceRelationshipUpdateInputsItem entity with the given context and client set.
func (srrui *ServiceResourceRelationshipUpdateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srrui == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	if srrui.Dependency != nil {
		if err := srrui.Dependency.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srrui.Dependency = nil
			}
		}
	}

	return nil
}

// ServiceResourceRelationshipUpdateInputs holds the modification input of the ServiceResourceRelationship entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ServiceResourceRelationshipUpdateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ServiceResourceRelationshipUpdateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the ServiceResourceRelationship entities for modifying,
// after validating.
func (srrui *ServiceResourceRelationshipUpdateInputs) Model() []*ServiceResourceRelationship {
	if srrui == nil || len(srrui.Items) == 0 {
		return nil
	}

	_srrs := make([]*ServiceResourceRelationship, len(srrui.Items))

	for i := range srrui.Items {
		_srr := &ServiceResourceRelationship{
			ID:   srrui.Items[i].ID,
			Type: srrui.Items[i].Type,
		}

		if srrui.Items[i].Dependency != nil {
			_srr.DependencyID = srrui.Items[i].Dependency.ID
		}

		_srrs[i] = _srr
	}

	return _srrs
}

// IDs returns the ID list of the ServiceResourceRelationship entities for modifying,
// after validating.
func (srrui *ServiceResourceRelationshipUpdateInputs) IDs() []object.ID {
	if srrui == nil || len(srrui.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(srrui.Items))
	for i := range srrui.Items {
		ids[i] = srrui.Items[i].ID
	}
	return ids
}

// Validate checks the ServiceResourceRelationshipUpdateInputs entity.
func (srrui *ServiceResourceRelationshipUpdateInputs) Validate() error {
	if srrui == nil {
		return errors.New("nil receiver")
	}

	return srrui.ValidateWith(srrui.inputConfig.Context, srrui.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceResourceRelationshipUpdateInputs entity with the given context and client set.
func (srrui *ServiceResourceRelationshipUpdateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srrui == nil {
		return errors.New("nil receiver")
	}

	if len(srrui.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.ServiceResourceRelationships().Query()

	ids := make([]object.ID, 0, len(srrui.Items))
	ors := make([]predicate.ServiceResourceRelationship, 0, len(srrui.Items))
	indexers := make(map[any][]int)

	for i := range srrui.Items {
		if srrui.Items[i] == nil {
			return errors.New("nil item")
		}

		if srrui.Items[i].ID != "" {
			ids = append(ids, srrui.Items[i].ID)
			ors = append(ors, serviceresourcerelationship.ID(srrui.Items[i].ID))
			indexers[srrui.Items[i].ID] = append(indexers[srrui.Items[i].ID], i)
		} else if srrui.Items[i].Type != "" {
			ors = append(ors, serviceresourcerelationship.And(
				serviceresourcerelationship.Type(srrui.Items[i].Type)))
			indexerKey := fmt.Sprint("/", srrui.Items[i].Type)
			indexers[indexerKey] = append(indexers[indexerKey], i)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	p := serviceresourcerelationship.IDIn(ids...)
	if len(ids) != cap(ids) {
		p = serviceresourcerelationship.Or(ors...)
	}

	es, err := q.
		Where(p).
		Select(
			serviceresourcerelationship.FieldID,
			serviceresourcerelationship.FieldType,
		).
		All(ctx)
	if err != nil {
		return err
	}

	if len(es) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range es {
		indexer := indexers[es[i].ID]
		if indexer == nil {
			indexerKey := fmt.Sprint("/", es[i].Type)
			indexer = indexers[indexerKey]
		}
		for _, j := range indexer {
			srrui.Items[j].ID = es[i].ID
			srrui.Items[j].Type = es[i].Type
		}
	}

	for i := range srrui.Items {
		if err := srrui.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// ServiceResourceRelationshipOutput holds the output of the ServiceResourceRelationship entity.
type ServiceResourceRelationshipOutput struct {
	ID         object.ID  `json:"id,omitempty"`
	CreateTime *time.Time `json:"createTime,omitempty"`
	Type       string     `json:"type,omitempty"`

	Dependency *ServiceResourceOutput `json:"dependency,omitempty"`
}

// View returns the output of ServiceResourceRelationship entity.
func (_srr *ServiceResourceRelationship) View() *ServiceResourceRelationshipOutput {
	return ExposeServiceResourceRelationship(_srr)
}

// View returns the output of ServiceResourceRelationship entities.
func (_srrs ServiceResourceRelationships) View() []*ServiceResourceRelationshipOutput {
	return ExposeServiceResourceRelationships(_srrs)
}

// ExposeServiceResourceRelationship converts the ServiceResourceRelationship to ServiceResourceRelationshipOutput.
func ExposeServiceResourceRelationship(_srr *ServiceResourceRelationship) *ServiceResourceRelationshipOutput {
	if _srr == nil {
		return nil
	}

	srro := &ServiceResourceRelationshipOutput{
		ID:         _srr.ID,
		CreateTime: _srr.CreateTime,
		Type:       _srr.Type,
	}

	if _srr.Edges.Dependency != nil {
		srro.Dependency = ExposeServiceResource(_srr.Edges.Dependency)
	} else if _srr.DependencyID != "" {
		srro.Dependency = &ServiceResourceOutput{
			ID: _srr.DependencyID,
		}
	}
	return srro
}

// ExposeServiceResourceRelationships converts the ServiceResourceRelationship slice to ServiceResourceRelationshipOutput pointer slice.
func ExposeServiceResourceRelationships(_srrs []*ServiceResourceRelationship) []*ServiceResourceRelationshipOutput {
	if len(_srrs) == 0 {
		return nil
	}

	srros := make([]*ServiceResourceRelationshipOutput, len(_srrs))
	for i := range _srrs {
		srros[i] = ExposeServiceResourceRelationship(_srrs[i])
	}
	return srros
}
