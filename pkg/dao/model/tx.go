// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	stdsql "database/sql"
	"fmt"
	"sync"

	"entgo.io/ent/dialect"
)

// Tx is a transactional client that is created by calling Client.Tx().
type Tx struct {
	config
	// Catalog is the client for interacting with the Catalog builders.
	Catalog *CatalogClient
	// Connector is the client for interacting with the Connector builders.
	Connector *ConnectorClient
	// CostReport is the client for interacting with the CostReport builders.
	CostReport *CostReportClient
	// DistributeLock is the client for interacting with the DistributeLock builders.
	DistributeLock *DistributeLockClient
	// Environment is the client for interacting with the Environment builders.
	Environment *EnvironmentClient
	// EnvironmentConnectorRelationship is the client for interacting with the EnvironmentConnectorRelationship builders.
	EnvironmentConnectorRelationship *EnvironmentConnectorRelationshipClient
	// Perspective is the client for interacting with the Perspective builders.
	Perspective *PerspectiveClient
	// Project is the client for interacting with the Project builders.
	Project *ProjectClient
	// Role is the client for interacting with the Role builders.
	Role *RoleClient
	// Service is the client for interacting with the Service builders.
	Service *ServiceClient
	// ServiceRelationship is the client for interacting with the ServiceRelationship builders.
	ServiceRelationship *ServiceRelationshipClient
	// ServiceResource is the client for interacting with the ServiceResource builders.
	ServiceResource *ServiceResourceClient
	// ServiceResourceRelationship is the client for interacting with the ServiceResourceRelationship builders.
	ServiceResourceRelationship *ServiceResourceRelationshipClient
	// ServiceRevision is the client for interacting with the ServiceRevision builders.
	ServiceRevision *ServiceRevisionClient
	// Setting is the client for interacting with the Setting builders.
	Setting *SettingClient
	// Subject is the client for interacting with the Subject builders.
	Subject *SubjectClient
	// SubjectRoleRelationship is the client for interacting with the SubjectRoleRelationship builders.
	SubjectRoleRelationship *SubjectRoleRelationshipClient
	// Template is the client for interacting with the Template builders.
	Template *TemplateClient
	// TemplateVersion is the client for interacting with the TemplateVersion builders.
	TemplateVersion *TemplateVersionClient
	// Token is the client for interacting with the Token builders.
	Token *TokenClient
	// Variable is the client for interacting with the Variable builders.
	Variable *VariableClient

	// lazily loaded.
	client     *Client
	clientOnce sync.Once
	// ctx lives for the life of the transaction. It is
	// the same context used by the underlying connection.
	ctx context.Context
}

type (
	// Committer is the interface that wraps the Commit method.
	Committer interface {
		Commit(context.Context, *Tx) error
	}

	// The CommitFunc type is an adapter to allow the use of ordinary
	// function as a Committer. If f is a function with the appropriate
	// signature, CommitFunc(f) is a Committer that calls f.
	CommitFunc func(context.Context, *Tx) error

	// CommitHook defines the "commit middleware". A function that gets a Committer
	// and returns a Committer. For example:
	//
	//	hook := func(next ent.Committer) ent.Committer {
	//		return ent.CommitFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Commit(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	CommitHook func(Committer) Committer
)

// Commit calls f(ctx, m).
func (f CommitFunc) Commit(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Commit commits the transaction.
func (tx *Tx) Commit() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Committer = CommitFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Commit()
	})
	txDriver.mu.Lock()
	hooks := append([]CommitHook(nil), txDriver.onCommit...)
	txDriver.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Commit(tx.ctx, tx)
}

// OnCommit adds a hook to call on commit.
func (tx *Tx) OnCommit(f CommitHook) {
	txDriver := tx.config.driver.(*txDriver)
	txDriver.mu.Lock()
	txDriver.onCommit = append(txDriver.onCommit, f)
	txDriver.mu.Unlock()
}

type (
	// Rollbacker is the interface that wraps the Rollback method.
	Rollbacker interface {
		Rollback(context.Context, *Tx) error
	}

	// The RollbackFunc type is an adapter to allow the use of ordinary
	// function as a Rollbacker. If f is a function with the appropriate
	// signature, RollbackFunc(f) is a Rollbacker that calls f.
	RollbackFunc func(context.Context, *Tx) error

	// RollbackHook defines the "rollback middleware". A function that gets a Rollbacker
	// and returns a Rollbacker. For example:
	//
	//	hook := func(next ent.Rollbacker) ent.Rollbacker {
	//		return ent.RollbackFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Rollback(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	RollbackHook func(Rollbacker) Rollbacker
)

// Rollback calls f(ctx, m).
func (f RollbackFunc) Rollback(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Rollback rollbacks the transaction.
func (tx *Tx) Rollback() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Rollbacker = RollbackFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Rollback()
	})
	txDriver.mu.Lock()
	hooks := append([]RollbackHook(nil), txDriver.onRollback...)
	txDriver.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Rollback(tx.ctx, tx)
}

// OnRollback adds a hook to call on rollback.
func (tx *Tx) OnRollback(f RollbackHook) {
	txDriver := tx.config.driver.(*txDriver)
	txDriver.mu.Lock()
	txDriver.onRollback = append(txDriver.onRollback, f)
	txDriver.mu.Unlock()
}

// Client returns a Client that binds to current transaction.
func (tx *Tx) Client() *Client {
	tx.clientOnce.Do(func() {
		tx.client = &Client{config: tx.config}
		tx.client.init()
	})
	return tx.client
}

func (tx *Tx) init() {
	tx.Catalog = NewCatalogClient(tx.config)
	tx.Connector = NewConnectorClient(tx.config)
	tx.CostReport = NewCostReportClient(tx.config)
	tx.DistributeLock = NewDistributeLockClient(tx.config)
	tx.Environment = NewEnvironmentClient(tx.config)
	tx.EnvironmentConnectorRelationship = NewEnvironmentConnectorRelationshipClient(tx.config)
	tx.Perspective = NewPerspectiveClient(tx.config)
	tx.Project = NewProjectClient(tx.config)
	tx.Role = NewRoleClient(tx.config)
	tx.Service = NewServiceClient(tx.config)
	tx.ServiceRelationship = NewServiceRelationshipClient(tx.config)
	tx.ServiceResource = NewServiceResourceClient(tx.config)
	tx.ServiceResourceRelationship = NewServiceResourceRelationshipClient(tx.config)
	tx.ServiceRevision = NewServiceRevisionClient(tx.config)
	tx.Setting = NewSettingClient(tx.config)
	tx.Subject = NewSubjectClient(tx.config)
	tx.SubjectRoleRelationship = NewSubjectRoleRelationshipClient(tx.config)
	tx.Template = NewTemplateClient(tx.config)
	tx.TemplateVersion = NewTemplateVersionClient(tx.config)
	tx.Token = NewTokenClient(tx.config)
	tx.Variable = NewVariableClient(tx.config)
}

// txDriver wraps the given dialect.Tx with a nop dialect.Driver implementation.
// The idea is to support transactions without adding any extra code to the builders.
// When a builder calls to driver.Tx(), it gets the same dialect.Tx instance.
// Commit and Rollback are nop for the internal builders and the user must call one
// of them in order to commit or rollback the transaction.
//
// If a closed transaction is embedded in one of the generated entities, and the entity
// applies a query, for example: Catalog.QueryXXX(), the query will be executed
// through the driver which created this transaction.
//
// Note that txDriver is not goroutine safe.
type txDriver struct {
	// the driver we started the transaction from.
	drv dialect.Driver
	// tx is the underlying transaction.
	tx dialect.Tx
	// completion hooks.
	mu         sync.Mutex
	onCommit   []CommitHook
	onRollback []RollbackHook
}

// newTx creates a new transactional driver.
func newTx(ctx context.Context, drv dialect.Driver) (*txDriver, error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &txDriver{tx: tx, drv: drv}, nil
}

// Tx returns the transaction wrapper (txDriver) to avoid Commit or Rollback calls
// from the internal builders. Should be called only by the internal builders.
func (tx *txDriver) Tx(context.Context) (dialect.Tx, error) { return tx, nil }

// Dialect returns the dialect of the driver we started the transaction from.
func (tx *txDriver) Dialect() string { return tx.drv.Dialect() }

// Close is a nop close.
func (*txDriver) Close() error { return nil }

// Commit is a nop commit for the internal builders.
// User must call `Tx.Commit` in order to commit the transaction.
func (*txDriver) Commit() error { return nil }

// Rollback is a nop rollback for the internal builders.
// User must call `Tx.Rollback` in order to rollback the transaction.
func (*txDriver) Rollback() error { return nil }

// Exec calls tx.Exec.
func (tx *txDriver) Exec(ctx context.Context, query string, args, v any) error {
	return tx.tx.Exec(ctx, query, args, v)
}

// Query calls tx.Query.
func (tx *txDriver) Query(ctx context.Context, query string, args, v any) error {
	return tx.tx.Query(ctx, query, args, v)
}

var _ dialect.Driver = (*txDriver)(nil)

// Catalogs implements the ClientSet.
func (tx *Tx) Catalogs() *CatalogClient {
	return tx.Catalog
}

// Connectors implements the ClientSet.
func (tx *Tx) Connectors() *ConnectorClient {
	return tx.Connector
}

// CostReports implements the ClientSet.
func (tx *Tx) CostReports() *CostReportClient {
	return tx.CostReport
}

// DistributeLocks implements the ClientSet.
func (tx *Tx) DistributeLocks() *DistributeLockClient {
	return tx.DistributeLock
}

// Environments implements the ClientSet.
func (tx *Tx) Environments() *EnvironmentClient {
	return tx.Environment
}

// EnvironmentConnectorRelationships implements the ClientSet.
func (tx *Tx) EnvironmentConnectorRelationships() *EnvironmentConnectorRelationshipClient {
	return tx.EnvironmentConnectorRelationship
}

// Perspectives implements the ClientSet.
func (tx *Tx) Perspectives() *PerspectiveClient {
	return tx.Perspective
}

// Projects implements the ClientSet.
func (tx *Tx) Projects() *ProjectClient {
	return tx.Project
}

// Roles implements the ClientSet.
func (tx *Tx) Roles() *RoleClient {
	return tx.Role
}

// Services implements the ClientSet.
func (tx *Tx) Services() *ServiceClient {
	return tx.Service
}

// ServiceRelationships implements the ClientSet.
func (tx *Tx) ServiceRelationships() *ServiceRelationshipClient {
	return tx.ServiceRelationship
}

// ServiceResources implements the ClientSet.
func (tx *Tx) ServiceResources() *ServiceResourceClient {
	return tx.ServiceResource
}

// ServiceResourceRelationships implements the ClientSet.
func (tx *Tx) ServiceResourceRelationships() *ServiceResourceRelationshipClient {
	return tx.ServiceResourceRelationship
}

// ServiceRevisions implements the ClientSet.
func (tx *Tx) ServiceRevisions() *ServiceRevisionClient {
	return tx.ServiceRevision
}

// Settings implements the ClientSet.
func (tx *Tx) Settings() *SettingClient {
	return tx.Setting
}

// Subjects implements the ClientSet.
func (tx *Tx) Subjects() *SubjectClient {
	return tx.Subject
}

// SubjectRoleRelationships implements the ClientSet.
func (tx *Tx) SubjectRoleRelationships() *SubjectRoleRelationshipClient {
	return tx.SubjectRoleRelationship
}

// Templates implements the ClientSet.
func (tx *Tx) Templates() *TemplateClient {
	return tx.Template
}

// TemplateVersions implements the ClientSet.
func (tx *Tx) TemplateVersions() *TemplateVersionClient {
	return tx.TemplateVersion
}

// Tokens implements the ClientSet.
func (tx *Tx) Tokens() *TokenClient {
	return tx.Token
}

// Variables implements the ClientSet.
func (tx *Tx) Variables() *VariableClient {
	return tx.Variable
}

// Dialect returns the dialect name of the driver.
func (tx *Tx) Dialect() string {
	return tx.driver.Dialect()
}

// Intercept adds the query interceptors to all the entity clients.
func (tx *Tx) Intercept(interceptors ...Interceptor) {
	tx.Catalog.Intercept(interceptors...)
	tx.Connector.Intercept(interceptors...)
	tx.CostReport.Intercept(interceptors...)
	tx.DistributeLock.Intercept(interceptors...)
	tx.Environment.Intercept(interceptors...)
	tx.EnvironmentConnectorRelationship.Intercept(interceptors...)
	tx.Perspective.Intercept(interceptors...)
	tx.Project.Intercept(interceptors...)
	tx.Role.Intercept(interceptors...)
	tx.Service.Intercept(interceptors...)
	tx.ServiceRelationship.Intercept(interceptors...)
	tx.ServiceResource.Intercept(interceptors...)
	tx.ServiceResourceRelationship.Intercept(interceptors...)
	tx.ServiceRevision.Intercept(interceptors...)
	tx.Setting.Intercept(interceptors...)
	tx.Subject.Intercept(interceptors...)
	tx.SubjectRoleRelationship.Intercept(interceptors...)
	tx.Template.Intercept(interceptors...)
	tx.TemplateVersion.Intercept(interceptors...)
	tx.Token.Intercept(interceptors...)
	tx.Variable.Intercept(interceptors...)
}

// ExecContext allows calling the underlying ExecContext method of the transaction if it is supported by it.
// See, database/sql#Tx.ExecContext for more information.
func (tx *txDriver) ExecContext(ctx context.Context, query string, args ...any) (stdsql.Result, error) {
	ex, ok := tx.tx.(interface {
		ExecContext(context.Context, string, ...any) (stdsql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Tx.ExecContext is not supported")
	}
	return ex.ExecContext(ctx, query, args...)
}

// QueryContext allows calling the underlying QueryContext method of the transaction if it is supported by it.
// See, database/sql#Tx.QueryContext for more information.
func (tx *txDriver) QueryContext(ctx context.Context, query string, args ...any) (*stdsql.Rows, error) {
	q, ok := tx.tx.(interface {
		QueryContext(context.Context, string, ...any) (*stdsql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Tx.QueryContext is not supported")
	}
	return q.QueryContext(ctx, query, args...)
}

// Use adds the mutation hooks to all the entity clients.
func (tx *Tx) Use(hooks ...Hook) {
	tx.Catalog.Use(hooks...)
	tx.Connector.Use(hooks...)
	tx.CostReport.Use(hooks...)
	tx.DistributeLock.Use(hooks...)
	tx.Environment.Use(hooks...)
	tx.EnvironmentConnectorRelationship.Use(hooks...)
	tx.Perspective.Use(hooks...)
	tx.Project.Use(hooks...)
	tx.Role.Use(hooks...)
	tx.Service.Use(hooks...)
	tx.ServiceRelationship.Use(hooks...)
	tx.ServiceResource.Use(hooks...)
	tx.ServiceResourceRelationship.Use(hooks...)
	tx.ServiceRevision.Use(hooks...)
	tx.Setting.Use(hooks...)
	tx.Subject.Use(hooks...)
	tx.SubjectRoleRelationship.Use(hooks...)
	tx.Template.Use(hooks...)
	tx.TemplateVersion.Use(hooks...)
	tx.Token.Use(hooks...)
	tx.Variable.Use(hooks...)
}

// WithDebug enables the debug mode of the client.
func (tx *Tx) WithDebug() ClientSet {
	if tx.debug {
		return tx
	}
	cfg := tx.config
	cfg.driver = dialect.Debug(tx.driver, tx.log)
	dtx := &Tx{
		ctx:    tx.ctx,
		config: cfg,
	}
	dtx.init()
	return dtx
}

// WithTx gives a new transactional client in the callback function,
// if already in a transaction, this will keep in the same transaction.
func (tx *Tx) WithTx(ctx context.Context, cb func(tx *Tx) error) error {
	return cb(tx)
}
