package view

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/seal-io/seal/pkg/apis/runtime"
	"github.com/seal-io/seal/pkg/dao/model"
	"github.com/seal-io/seal/pkg/dao/model/connector"
	"github.com/seal-io/seal/pkg/dao/model/environment"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
	"github.com/seal-io/seal/pkg/dao/model/project"
	"github.com/seal-io/seal/pkg/dao/model/templateversion"
	"github.com/seal-io/seal/pkg/dao/types/object"
	"github.com/seal-io/seal/utils/strs"
	"github.com/seal-io/seal/utils/validation"
)

// Basic APIs.

type CreateRequest struct {
	model.EnvironmentCreateInput `json:",inline"`

	//nolint:staticcheck
	Services []model.ServiceCreateInput `json:"services,cli-ignore"`

	// Project id and project name are used in permission checking.
	ProjectID   object.ID `query:"projectID,omitempty"`
	ProjectName string    `query:"projectName,omitempty"`
}

func (r *CreateRequest) ValidateWith(ctx context.Context, input any) error {
	if err := validation.IsDNSLabel(r.Name); err != nil {
		return fmt.Errorf("invalid name: %w", err)
	}

	modelClient := input.(model.ClientSet)

	switch {
	case r.ProjectID != "":
		if !r.ProjectID.Valid() {
			return errors.New("invalid project id: blank")
		}

		r.Project = &model.ProjectQueryInput{
			ID: r.ProjectID,
		}
	case r.ProjectName != "":
		projectID, err := modelClient.Projects().Query().
			Where(project.Name(r.ProjectName)).
			OnlyID(ctx)
		if err != nil {
			return runtime.Errorw(err, "failed to get project")
		}

		r.ProjectID = projectID
		r.Project = &model.ProjectQueryInput{
			ID: projectID,
		}
	default:
		return errors.New("both project id and project name are blank")
	}

	// Get template versions.
	templateVersionKeys := sets.NewString()
	templateVersionPredicates := make([]predicate.TemplateVersion, 0)

	for _, s := range r.Services {
		key := strs.Join("/", s.Template.Name, s.Template.Version)
		if templateVersionKeys.Has(key) {
			continue
		}

		templateVersionKeys.Insert(key)

		templateVersionPredicates = append(templateVersionPredicates, templateversion.And(
			templateversion.TemplateName(s.Template.Name),
			templateversion.Version(s.Template.Version),
		))
	}

	templateVersions, err := modelClient.TemplateVersions().Query().
		Select(
			templateversion.FieldTemplateID,
			templateversion.FieldVersion,
			templateversion.FieldSchema,
		).
		Where(templateversion.Or(
			templateVersionPredicates...,
		)).
		All(ctx)
	if err != nil {
		return runtime.Errorw(err, "failed to get template version")
	}
	templateVersionMap := make(map[string]*model.TemplateVersion, len(templateVersions))

	for _, tv := range templateVersions {
		key := strs.Join("/", tv.TemplateName, tv.Version)
		if _, ok := templateVersionMap[key]; !ok {
			templateVersionMap[key] = tv
		}
	}

	for _, s := range r.Services {
		if s.Name == "" {
			return errors.New("invalid service name: blank")
		}

		if err := validation.IsDNSLabel(s.Name); err != nil {
			return fmt.Errorf("invalid name: %w", err)
		}

		// Verify template version.
		key := strs.Join("/", s.Template.Name, s.Template.Version)

		templateVersion, ok := templateVersionMap[key]
		if !ok {
			return runtime.Errorw(err, "failed to get template version")
		}

		// Verify variables with variables schema that defined on the template version.
		err = s.Attributes.ValidateWith(templateVersion.Schema.Variables)
		if err != nil {
			return fmt.Errorf("invalid variables: %w", err)
		}
	}

	connectorIDs := make([]object.ID, len(r.Connectors))
	for i, c := range r.Connectors {
		connectorIDs[i] = c.Connector.ID
	}

	if err = validateConnectors(ctx, modelClient, connectorIDs); err != nil {
		return err
	}

	return nil
}

type CreateResponse = *model.EnvironmentOutput

type DeleteRequest = GetRequest

type UpdateRequest struct {
	model.EnvironmentUpdateInput `uri:",inline" json:",inline"`
}

func (r *UpdateRequest) ValidateWith(ctx context.Context, input any) error {
	if !r.ID.Valid() {
		return errors.New("invalid id: blank")
	}

	if err := validation.IsDNSLabel(r.Name); err != nil {
		return fmt.Errorf("invalid name: %w", err)
	}

	modelClient := input.(model.ClientSet)

	exist, err := modelClient.Environments().Query().
		Where(environment.ID(r.ID)).
		Exist(ctx)
	if err != nil {
		return runtime.Errorw(err, "failed to get environment")
	}

	if !exist {
		return runtime.Error(http.StatusNotFound, "environment is not found")
	}

	connectorIDs := make([]object.ID, len(r.Connectors))
	for i, c := range r.Connectors {
		connectorIDs[i] = c.Connector.ID
	}

	if err = validateConnectors(ctx, modelClient, connectorIDs); err != nil {
		return err
	}

	return nil
}

// validateConnectors checks if given connector IDs are valid within the same project or globally.
func validateConnectors(ctx context.Context, client model.ClientSet, connectorIDs []object.ID) error {
	if len(connectorIDs) == 0 {
		return nil
	}

	for _, id := range connectorIDs {
		if !id.Valid() {
			return errors.New("invalid connector id: blank")
		}
	}

	var typeCount []struct {
		Type  string `json:"type"`
		Count int    `json:"count"`
	}

	err := client.Connectors().Query().
		Where(connector.IDIn(connectorIDs...)).
		GroupBy(connector.FieldType).
		Aggregate(model.Count()).
		Scan(ctx, &typeCount)
	if err != nil {
		return runtime.Errorw(err, "failed to get connector type count")
	}

	// Validate connector type is duplicated, only one connector type is allowed in one environment.
	for _, c := range typeCount {
		if c.Count > 1 {
			return fmt.Errorf("invalid connectors: duplicated connector type %s", c.Type)
		}
	}

	validCount, err := client.Connectors().Query().
		Where(connector.IDIn(connectorIDs...)).
		Count(ctx)
	if err != nil {
		return runtime.Errorw(err, "failed to get connectors")
	}

	if validCount != len(connectorIDs) {
		return fmt.Errorf("invalid connector id")
	}

	return nil
}

type GetRequest struct {
	model.EnvironmentQueryInput `uri:",inline"`
}

func (r *GetRequest) Validate() error {
	if !r.ID.Valid() {
		return errors.New("invalid id: blank")
	}

	return nil
}

type GetResponse = *model.EnvironmentOutput

// Batch APIs.

type CollectionDeleteRequest []*model.EnvironmentQueryInput

func (r CollectionDeleteRequest) Validate() error {
	if len(r) == 0 {
		return errors.New("invalid input: empty")
	}

	for _, i := range r {
		if !i.ID.Valid() {
			return errors.New("invalid id: blank")
		}
	}

	return nil
}

type CollectionGetRequest struct {
	runtime.RequestCollection[predicate.Environment, environment.OrderOption] `query:",inline"`
	ProjectID                                                                 object.ID `query:"projectID,omitempty"`
	ProjectName                                                               string    `query:"projectName,omitempty"`
}

func (r *CollectionGetRequest) ValidateWith(ctx context.Context, input any) error {
	modelClient := input.(model.ClientSet)

	switch {
	case r.ProjectID != "":
		if !r.ProjectID.Valid() {
			return errors.New("invalid project id: blank")
		}
	case r.ProjectName != "":
		projectID, err := modelClient.Projects().Query().
			Where(project.Name(r.ProjectName)).
			OnlyID(ctx)
		if err != nil {
			return runtime.Errorw(err, "failed to get project")
		}

		r.ProjectID = projectID
	}

	return nil
}

type CollectionGetResponse = []*model.EnvironmentOutput

// Extensional APIs.
