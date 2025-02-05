package terraform

import (
	"context"

	revisionbus "github.com/seal-io/walrus/pkg/bus/servicerevision"
	"github.com/seal-io/walrus/pkg/dao/model/service"
	"github.com/seal-io/walrus/pkg/dao/types/status"
)

// SyncServiceRevisionStatus updates the status of the service according to its recent finished service revision.
func SyncServiceRevisionStatus(ctx context.Context, bm revisionbus.BusMessage) (err error) {
	var (
		mc       = bm.TransactionalModelClient
		revision = bm.Refer
	)

	// Report to service.
	entity, err := mc.Services().Query().
		Where(service.ID(revision.ServiceID)).
		Select(
			service.FieldID,
			service.FieldStatus).
		Only(ctx)
	if err != nil {
		return err
	}

	if status.ServiceRevisionStatusReady.IsTrue(revision) {
		if status.ServiceStatusDeleted.IsUnknown(entity) {
			return mc.Services().DeleteOne(entity).
				Exec(ctx)
		}

		status.ServiceStatusDeployed.True(entity, "")
		status.ServiceStatusReady.Unknown(entity, "")
	} else if status.ServiceRevisionStatusReady.IsFalse(revision) {
		if status.ServiceStatusDeleted.IsUnknown(entity) {
			status.ServiceStatusDeleted.False(entity, "")
		} else {
			status.ServiceStatusDeployed.False(entity, "")
		}

		entity.Status.SummaryStatusMessage = revision.Status.SummaryStatusMessage
	}

	entity.Status.SetSummary(status.WalkService(&entity.Status))

	return mc.Services().UpdateOne(entity).
		SetStatus(entity.Status).
		Exec(ctx)
}
