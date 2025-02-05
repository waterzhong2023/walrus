package servicerevision

import (
	"github.com/seal-io/walrus/pkg/apis/runtime"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/servicerevision"
	"github.com/seal-io/walrus/pkg/datalisten/modelchange"
	"github.com/seal-io/walrus/utils/topic"
)

func (h Handler) Get(req GetRequest) (GetResponse, error) {
	entity, err := h.modelClient.ServiceRevisions().Get(req.Context, req.ID)
	if err != nil {
		return nil, err
	}

	return model.ExposeServiceRevision(entity), nil
}

var (
	getFields = servicerevision.WithoutFields(
		servicerevision.FieldRecord,
		servicerevision.FieldInputPlan,
		servicerevision.FieldOutput,
		servicerevision.FieldTemplateName,
		servicerevision.FieldTemplateVersion,
		servicerevision.FieldAttributes,
		servicerevision.FieldVariables,
	)
	sortFields = []string{
		servicerevision.FieldCreateTime,
	}
)

func (h Handler) CollectionGet(req CollectionGetRequest) (CollectionGetResponse, int, error) {
	query := h.modelClient.ServiceRevisions().Query()

	if req.Service != nil && req.Service.ID != "" {
		query.Where(servicerevision.ServiceID(req.Service.ID))
	}

	if stream := req.Stream; stream != nil {
		// Handle stream request.
		if fields, ok := req.Extracting(getFields, getFields...); ok {
			query.Select(fields...)
		}

		if orders, ok := req.Sorting(sortFields, model.Desc(servicerevision.FieldCreateTime)); ok {
			query.Order(orders...)
		}

		t, err := topic.Subscribe(modelchange.ServiceRevision)
		if err != nil {
			return nil, 0, err
		}

		defer func() { t.Unsubscribe() }()

		for {
			var event topic.Event

			event, err = t.Receive(stream)
			if err != nil {
				return nil, 0, err
			}

			dm, ok := event.Data.(modelchange.Event)
			if !ok {
				continue
			}

			var items []*model.ServiceRevisionOutput

			switch dm.Type {
			case modelchange.EventTypeCreate, modelchange.EventTypeUpdate:
				revisions, err := query.Clone().
					Where(servicerevision.IDIn(dm.IDs...)).
					// Must append service ID.
					Select(servicerevision.FieldServiceID).
					Unique(false).
					All(stream)
				if err != nil {
					return nil, 0, err
				}

				items = model.ExposeServiceRevisions(revisions)
			case modelchange.EventTypeDelete:
				items = make([]*model.ServiceRevisionOutput, len(dm.IDs))
				for i := range dm.IDs {
					items[i] = &model.ServiceRevisionOutput{
						ID: dm.IDs[i],
					}
				}
			}

			if len(items) == 0 {
				continue
			}

			resp := runtime.TypedResponse(dm.Type.String(), items)
			if err = stream.SendJSON(resp); err != nil {
				return nil, 0, err
			}
		}
	}

	// Handler normal request.

	// Get count.
	cnt, err := query.Clone().Count(req.Context)
	if err != nil {
		return nil, 0, err
	}

	// Get entities.
	if limit, offset, ok := req.Paging(); ok {
		query.Limit(limit).Offset(offset)
	}

	if fields, ok := req.Extracting(getFields, getFields...); ok {
		query.Select(fields...)
	}

	if orders, ok := req.Sorting(sortFields, model.Desc(servicerevision.FieldCreateTime)); ok {
		query.Order(orders...)
	}

	entities, err := query.
		// Must append service ID.
		Select(servicerevision.FieldServiceID).
		Unique(false).
		All(req.Context)
	if err != nil {
		return nil, 0, err
	}

	return model.ExposeServiceRevisions(entities), cnt, nil
}

func (h Handler) CollectionDelete(req CollectionDeleteRequest) error {
	ids := req.IDs()

	return h.modelClient.WithTx(req.Context, func(tx *model.Tx) error {
		_, err := tx.ServiceRevisions().Delete().
			Where(servicerevision.IDIn(ids...)).
			Exec(req.Context)

		return err
	})
}
