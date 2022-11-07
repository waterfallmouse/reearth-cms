package integration

import (
	"context"
	"errors"

	"github.com/reearth/reearth-cms/server/internal/adapter"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/integrationapi"
	"github.com/reearth/reearth-cms/server/pkg/item"
	"github.com/reearth/reearthx/usecasex"
	"github.com/samber/lo"
)

func (s Server) ItemFilter(ctx context.Context, request ItemFilterRequestObject) (ItemFilterResponseObject, error) {
	op := adapter.Operator(ctx)
	uc := adapter.Usecases(ctx)
	m, err := uc.Model.FindByIDs(ctx, []id.ModelID{id.ModelID(request.ModelId)}, op)
	if err != nil {
		return nil, err
	}

	p := &usecasex.Pagination{
		Before: nil,
		After:  nil,
		First:  lo.ToPtr(1000),
		Last:   nil,
	}

	items, pi, err := adapter.Usecases(ctx).Item.FindBySchema(ctx, m[0].Schema(), p, op)
	if err != nil {
		return ItemFilter400Response{}, err
	}

	itemList := lo.Map(items, func(i *item.Item, _ int) integrationapi.Item {
		ver, err := uc.Item.FindAllVersionsByID(ctx, i.ID(), op)
		if err != nil {
			return integrationapi.Item{}
		}

		return toItem(i, ver[len(ver)-1], m[0].ID())
	})

	return ItemFilter200JSONResponse{
		Items:      &itemList,
		Page:       lo.ToPtr(1),
		PerPage:    lo.ToPtr(1000),
		TotalCount: lo.ToPtr(pi.TotalCount),
	}, nil
}

func (s Server) ItemCreate(ctx context.Context, request ItemCreateRequestObject) (ItemCreateResponseObject, error) {
	op := adapter.Operator(ctx)
	uc := adapter.Usecases(ctx)

	if request.Body.Fields == nil {
		return ItemCreate400Response{}, errors.New("missing fields")
	}

	m, err := uc.Model.FindByIDs(ctx, []id.ModelID{id.ModelID(request.ModelId)}, op)
	if err != nil {
		return nil, err
	}

	cp := interfaces.CreateItemParam{
		SchemaID: m[0].Schema(),
		Fields: lo.Map(*request.Body.Fields, func(f integrationapi.Field, _ int) interfaces.ItemFieldParam {
			return toItemFieldParam(f)
		}),
	}

	i, err := uc.Item.Create(ctx, cp, op)
	if err != nil {
		return ItemCreate400Response{}, err
	}

	ver, err := uc.Item.FindAllVersionsByID(ctx, i.ID(), op)
	if err != nil {
		return nil, err
	}

	return ItemCreate200JSONResponse(toItem(i, ver[len(ver)-1], id.NewModelID())), nil
}

func (s Server) ItemDelete(ctx context.Context, request ItemDeleteRequestObject) (ItemDeleteResponseObject, error) {
	op := adapter.Operator(ctx)
	uc := adapter.Usecases(ctx)
	iId := id.ItemID(request.ItemId)

	err := uc.Item.Delete(ctx, iId, op)
	if err != nil {
		return ItemDelete400Response{}, err
	}
	return ItemDelete200JSONResponse{
		Id: &iId,
	}, nil
}

func (s Server) ItemGet(ctx context.Context, request ItemGetRequestObject) (ItemGetResponseObject, error) {
	op := adapter.Operator(ctx)
	uc := adapter.Usecases(ctx)
	iId := id.ItemID(request.ItemId)

	itm, err := uc.Item.FindByID(ctx, iId, op)
	if err != nil {
		return ItemGet400Response{}, err
	}

	ver, err := uc.Item.FindAllVersionsByID(ctx, iId, op)
	if err != nil {
		return nil, err
	}

	return ItemGet200JSONResponse(toItem(itm, ver[len(ver)-1], id.NewModelID())), nil
}

func (s Server) ItemPublish(ctx context.Context, request ItemPublishRequestObject) (ItemPublishResponseObject, error) {
	// TODO implement me
	panic("implement me")
}