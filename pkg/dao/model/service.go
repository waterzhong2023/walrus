// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"

	"github.com/seal-io/walrus/pkg/dao/model/environment"
	"github.com/seal-io/walrus/pkg/dao/model/project"
	"github.com/seal-io/walrus/pkg/dao/model/service"
	"github.com/seal-io/walrus/pkg/dao/model/templateversion"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/property"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/utils/json"
)

// Service is the model entity for the Service schema.
type Service struct {
	config `json:"-"`
	// ID of the ent.
	ID object.ID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `json:"labels,omitempty"`
	// Annotations holds the value of the "annotations" field.
	Annotations map[string]string `json:"annotations,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// Status holds the value of the "status" field.
	Status status.Status `json:"status,omitempty"`
	// ID of the project to belong.
	ProjectID object.ID `json:"project_id,omitempty"`
	// ID of the environment to which the service deploys.
	EnvironmentID object.ID `json:"environment_id,omitempty"`
	// ID of the template version to which the service belong.
	TemplateID object.ID `json:"template_id,omitempty"`
	// Attributes to configure the template.
	Attributes property.Values `json:"attributes,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ServiceQuery when eager-loading is set.
	Edges        ServiceEdges `json:"edges,omitempty"`
	selectValues sql.SelectValues
}

// ServiceEdges holds the relations/edges for other nodes in the graph.
type ServiceEdges struct {
	// Project to which the service belongs.
	Project *Project `json:"project,omitempty"`
	// Environment to which the service belongs.
	Environment *Environment `json:"environment,omitempty"`
	// Template to which the service belongs.
	Template *TemplateVersion `json:"template,omitempty"`
	// Revisions that belong to the service.
	Revisions []*ServiceRevision `json:"revisions,omitempty"`
	// Resources that belong to the service.
	Resources []*ServiceResource `json:"resources,omitempty"`
	// Dependencies holds the value of the dependencies edge.
	Dependencies []*ServiceRelationship `json:"dependencies,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ServiceEdges) ProjectOrErr() (*Project, error) {
	if e.loadedTypes[0] {
		if e.Project == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// EnvironmentOrErr returns the Environment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ServiceEdges) EnvironmentOrErr() (*Environment, error) {
	if e.loadedTypes[1] {
		if e.Environment == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: environment.Label}
		}
		return e.Environment, nil
	}
	return nil, &NotLoadedError{edge: "environment"}
}

// TemplateOrErr returns the Template value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ServiceEdges) TemplateOrErr() (*TemplateVersion, error) {
	if e.loadedTypes[2] {
		if e.Template == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: templateversion.Label}
		}
		return e.Template, nil
	}
	return nil, &NotLoadedError{edge: "template"}
}

// RevisionsOrErr returns the Revisions value or an error if the edge
// was not loaded in eager-loading.
func (e ServiceEdges) RevisionsOrErr() ([]*ServiceRevision, error) {
	if e.loadedTypes[3] {
		return e.Revisions, nil
	}
	return nil, &NotLoadedError{edge: "revisions"}
}

// ResourcesOrErr returns the Resources value or an error if the edge
// was not loaded in eager-loading.
func (e ServiceEdges) ResourcesOrErr() ([]*ServiceResource, error) {
	if e.loadedTypes[4] {
		return e.Resources, nil
	}
	return nil, &NotLoadedError{edge: "resources"}
}

// DependenciesOrErr returns the Dependencies value or an error if the edge
// was not loaded in eager-loading.
func (e ServiceEdges) DependenciesOrErr() ([]*ServiceRelationship, error) {
	if e.loadedTypes[5] {
		return e.Dependencies, nil
	}
	return nil, &NotLoadedError{edge: "dependencies"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Service) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case service.FieldLabels, service.FieldAnnotations, service.FieldStatus:
			values[i] = new([]byte)
		case service.FieldID, service.FieldProjectID, service.FieldEnvironmentID, service.FieldTemplateID:
			values[i] = new(object.ID)
		case service.FieldAttributes:
			values[i] = new(property.Values)
		case service.FieldName, service.FieldDescription:
			values[i] = new(sql.NullString)
		case service.FieldCreateTime, service.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Service fields.
func (s *Service) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case service.FieldID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case service.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case service.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		case service.FieldLabels:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field labels", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.Labels); err != nil {
					return fmt.Errorf("unmarshal field labels: %w", err)
				}
			}
		case service.FieldAnnotations:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field annotations", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.Annotations); err != nil {
					return fmt.Errorf("unmarshal field annotations: %w", err)
				}
			}
		case service.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				s.CreateTime = new(time.Time)
				*s.CreateTime = value.Time
			}
		case service.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				s.UpdateTime = new(time.Time)
				*s.UpdateTime = value.Time
			}
		case service.FieldStatus:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.Status); err != nil {
					return fmt.Errorf("unmarshal field status: %w", err)
				}
			}
		case service.FieldProjectID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value != nil {
				s.ProjectID = *value
			}
		case service.FieldEnvironmentID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field environment_id", values[i])
			} else if value != nil {
				s.EnvironmentID = *value
			}
		case service.FieldTemplateID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field template_id", values[i])
			} else if value != nil {
				s.TemplateID = *value
			}
		case service.FieldAttributes:
			if value, ok := values[i].(*property.Values); !ok {
				return fmt.Errorf("unexpected type %T for field attributes", values[i])
			} else if value != nil {
				s.Attributes = *value
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Service.
// This includes values selected through modifiers, order, etc.
func (s *Service) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryProject queries the "project" edge of the Service entity.
func (s *Service) QueryProject() *ProjectQuery {
	return NewServiceClient(s.config).QueryProject(s)
}

// QueryEnvironment queries the "environment" edge of the Service entity.
func (s *Service) QueryEnvironment() *EnvironmentQuery {
	return NewServiceClient(s.config).QueryEnvironment(s)
}

// QueryTemplate queries the "template" edge of the Service entity.
func (s *Service) QueryTemplate() *TemplateVersionQuery {
	return NewServiceClient(s.config).QueryTemplate(s)
}

// QueryRevisions queries the "revisions" edge of the Service entity.
func (s *Service) QueryRevisions() *ServiceRevisionQuery {
	return NewServiceClient(s.config).QueryRevisions(s)
}

// QueryResources queries the "resources" edge of the Service entity.
func (s *Service) QueryResources() *ServiceResourceQuery {
	return NewServiceClient(s.config).QueryResources(s)
}

// QueryDependencies queries the "dependencies" edge of the Service entity.
func (s *Service) QueryDependencies() *ServiceRelationshipQuery {
	return NewServiceClient(s.config).QueryDependencies(s)
}

// Update returns a builder for updating this Service.
// Note that you need to call Service.Unwrap() before calling this method if this Service
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Service) Update() *ServiceUpdateOne {
	return NewServiceClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Service entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Service) Unwrap() *Service {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("model: Service is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Service) String() string {
	var builder strings.Builder
	builder.WriteString("Service(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(s.Description)
	builder.WriteString(", ")
	builder.WriteString("labels=")
	builder.WriteString(fmt.Sprintf("%v", s.Labels))
	builder.WriteString(", ")
	builder.WriteString("annotations=")
	builder.WriteString(fmt.Sprintf("%v", s.Annotations))
	builder.WriteString(", ")
	if v := s.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := s.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", s.Status))
	builder.WriteString(", ")
	builder.WriteString("project_id=")
	builder.WriteString(fmt.Sprintf("%v", s.ProjectID))
	builder.WriteString(", ")
	builder.WriteString("environment_id=")
	builder.WriteString(fmt.Sprintf("%v", s.EnvironmentID))
	builder.WriteString(", ")
	builder.WriteString("template_id=")
	builder.WriteString(fmt.Sprintf("%v", s.TemplateID))
	builder.WriteString(", ")
	builder.WriteString("attributes=")
	builder.WriteString(fmt.Sprintf("%v", s.Attributes))
	builder.WriteByte(')')
	return builder.String()
}

// Services is a parsable slice of Service.
type Services []*Service
