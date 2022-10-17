// Package integration provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.1-0.20221010081618-45eb6bdf117a DO NOT EDIT.
package integration

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (DELETE /items/{itemId})
	ItemDelete(ctx echo.Context, itemId ItemIdParam) error
	// Returns an items.
	// (GET /items/{itemId})
	ItemGet(ctx echo.Context, itemId ItemIdParam, params ItemGetParams) error
	// Returns a list of items.
	// (GET /models/{modelId}/items)
	ItemFilter(ctx echo.Context, modelId ModelIdParam, params ItemFilterParams) error

	// (POST /models/{modelId}/items)
	ItemCreate(ctx echo.Context, modelId ModelIdParam) error
	// Set ref and version.
	// (POST /models/{modelId}/items/{itemId}/refs)
	ItemPublish(ctx echo.Context, modelId ModelIdParam, itemId ItemIdParam) error
	// Returns a list of assets.
	// (GET /projects/{projectId}/assets)
	AssetFilter(ctx echo.Context, projectId ProjectIdParam, params AssetFilterParams) error
	// Returns a list of assets.
	// (POST /projects/{projectId}/assets)
	AssetCreate(ctx echo.Context, projectId ProjectIdParam) error

	// (DELETE /projects/{projectId}/assets/{assetId})
	AssetDelete(ctx echo.Context, projectId ProjectIdParam, assetId AssetIdParam) error

	// (GET /projects/{projectId}/assets/{assetId})
	AssetGet(ctx echo.Context, projectId ProjectIdParam, assetId AssetIdParam) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ItemDelete converts echo context to params.
func (w *ServerInterfaceWrapper) ItemDelete(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "itemId" -------------
	var itemId ItemIdParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "itemId", runtime.ParamLocationPath, ctx.Param("itemId"), &itemId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter itemId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ItemDelete(ctx, itemId)
	return err
}

// ItemGet converts echo context to params.
func (w *ServerInterfaceWrapper) ItemGet(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "itemId" -------------
	var itemId ItemIdParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "itemId", runtime.ParamLocationPath, ctx.Param("itemId"), &itemId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter itemId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ItemGetParams
	// ------------- Optional query parameter "ref" -------------

	err = runtime.BindQueryParameter("form", true, false, "ref", ctx.QueryParams(), &params.Ref)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter ref: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ItemGet(ctx, itemId, params)
	return err
}

// ItemFilter converts echo context to params.
func (w *ServerInterfaceWrapper) ItemFilter(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "modelId" -------------
	var modelId ModelIdParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "modelId", runtime.ParamLocationPath, ctx.Param("modelId"), &modelId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter modelId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ItemFilterParams
	// ------------- Optional query parameter "sort" -------------

	err = runtime.BindQueryParameter("form", true, false, "sort", ctx.QueryParams(), &params.Sort)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sort: %s", err))
	}

	// ------------- Optional query parameter "dir" -------------

	err = runtime.BindQueryParameter("form", true, false, "dir", ctx.QueryParams(), &params.Dir)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter dir: %s", err))
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "perPage" -------------

	err = runtime.BindQueryParameter("form", true, false, "perPage", ctx.QueryParams(), &params.PerPage)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter perPage: %s", err))
	}

	// ------------- Optional query parameter "ref" -------------

	err = runtime.BindQueryParameter("form", true, false, "ref", ctx.QueryParams(), &params.Ref)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter ref: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ItemFilter(ctx, modelId, params)
	return err
}

// ItemCreate converts echo context to params.
func (w *ServerInterfaceWrapper) ItemCreate(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "modelId" -------------
	var modelId ModelIdParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "modelId", runtime.ParamLocationPath, ctx.Param("modelId"), &modelId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter modelId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ItemCreate(ctx, modelId)
	return err
}

// ItemPublish converts echo context to params.
func (w *ServerInterfaceWrapper) ItemPublish(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "modelId" -------------
	var modelId ModelIdParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "modelId", runtime.ParamLocationPath, ctx.Param("modelId"), &modelId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter modelId: %s", err))
	}

	// ------------- Path parameter "itemId" -------------
	var itemId ItemIdParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "itemId", runtime.ParamLocationPath, ctx.Param("itemId"), &itemId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter itemId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ItemPublish(ctx, modelId, itemId)
	return err
}

// AssetFilter converts echo context to params.
func (w *ServerInterfaceWrapper) AssetFilter(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "projectId" -------------
	var projectId ProjectIdParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "projectId", runtime.ParamLocationPath, ctx.Param("projectId"), &projectId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter projectId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params AssetFilterParams
	// ------------- Optional query parameter "sort" -------------

	err = runtime.BindQueryParameter("form", true, false, "sort", ctx.QueryParams(), &params.Sort)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sort: %s", err))
	}

	// ------------- Optional query parameter "dir" -------------

	err = runtime.BindQueryParameter("form", true, false, "dir", ctx.QueryParams(), &params.Dir)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter dir: %s", err))
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "perPage" -------------

	err = runtime.BindQueryParameter("form", true, false, "perPage", ctx.QueryParams(), &params.PerPage)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter perPage: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AssetFilter(ctx, projectId, params)
	return err
}

// AssetCreate converts echo context to params.
func (w *ServerInterfaceWrapper) AssetCreate(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "projectId" -------------
	var projectId ProjectIdParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "projectId", runtime.ParamLocationPath, ctx.Param("projectId"), &projectId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter projectId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AssetCreate(ctx, projectId)
	return err
}

// AssetDelete converts echo context to params.
func (w *ServerInterfaceWrapper) AssetDelete(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "projectId" -------------
	var projectId ProjectIdParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "projectId", runtime.ParamLocationPath, ctx.Param("projectId"), &projectId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter projectId: %s", err))
	}

	// ------------- Path parameter "assetId" -------------
	var assetId AssetIdParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "assetId", runtime.ParamLocationPath, ctx.Param("assetId"), &assetId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter assetId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AssetDelete(ctx, projectId, assetId)
	return err
}

// AssetGet converts echo context to params.
func (w *ServerInterfaceWrapper) AssetGet(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "projectId" -------------
	var projectId ProjectIdParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "projectId", runtime.ParamLocationPath, ctx.Param("projectId"), &projectId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter projectId: %s", err))
	}

	// ------------- Path parameter "assetId" -------------
	var assetId AssetIdParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "assetId", runtime.ParamLocationPath, ctx.Param("assetId"), &assetId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter assetId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AssetGet(ctx, projectId, assetId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE(baseURL+"/items/:itemId", wrapper.ItemDelete)
	router.GET(baseURL+"/items/:itemId", wrapper.ItemGet)
	router.GET(baseURL+"/models/:modelId/items", wrapper.ItemFilter)
	router.POST(baseURL+"/models/:modelId/items", wrapper.ItemCreate)
	router.POST(baseURL+"/models/:modelId/items/:itemId/refs", wrapper.ItemPublish)
	router.GET(baseURL+"/projects/:projectId/assets", wrapper.AssetFilter)
	router.POST(baseURL+"/projects/:projectId/assets", wrapper.AssetCreate)
	router.DELETE(baseURL+"/projects/:projectId/assets/:assetId", wrapper.AssetDelete)
	router.GET(baseURL+"/projects/:projectId/assets/:assetId", wrapper.AssetGet)

}

type UnauthorizedErrorResponse struct {
}

type ItemDeleteRequestObject struct {
	ItemId ItemIdParam `json:"itemId"`
}

type ItemDeleteResponseObject interface {
	VisitItemDeleteResponse(w http.ResponseWriter) error
}

type ItemDelete200JSONResponse struct {
	Id *openapi_types.UUID `json:"id,omitempty"`
}

func (response ItemDelete200JSONResponse) VisitItemDeleteResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ItemDelete400Response struct {
}

func (response ItemDelete400Response) VisitItemDeleteResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type ItemDelete401Response = UnauthorizedErrorResponse

func (response ItemDelete401Response) VisitItemDeleteResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type ItemDelete404Response struct {
}

func (response ItemDelete404Response) VisitItemDeleteResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type ItemGetRequestObject struct {
	ItemId ItemIdParam `json:"itemId"`
	Params ItemGetParams
}

type ItemGetResponseObject interface {
	VisitItemGetResponse(w http.ResponseWriter) error
}

type ItemGet200JSONResponse Item

func (response ItemGet200JSONResponse) VisitItemGetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ItemGet400Response struct {
}

func (response ItemGet400Response) VisitItemGetResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type ItemGet401Response = UnauthorizedErrorResponse

func (response ItemGet401Response) VisitItemGetResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type ItemFilterRequestObject struct {
	ModelId ModelIdParam `json:"modelId"`
	Params  ItemFilterParams
}

type ItemFilterResponseObject interface {
	VisitItemFilterResponse(w http.ResponseWriter) error
}

type ItemFilter200JSONResponse struct {
	Items      *[]Item `json:"items,omitempty"`
	Page       *int    `json:"page,omitempty"`
	PerPage    *int    `json:"perPage,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

func (response ItemFilter200JSONResponse) VisitItemFilterResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ItemFilter400Response struct {
}

func (response ItemFilter400Response) VisitItemFilterResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type ItemFilter401Response = UnauthorizedErrorResponse

func (response ItemFilter401Response) VisitItemFilterResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type ItemCreateRequestObject struct {
	ModelId ModelIdParam `json:"modelId"`
	Body    *ItemCreateJSONRequestBody
}

type ItemCreateResponseObject interface {
	VisitItemCreateResponse(w http.ResponseWriter) error
}

type ItemCreate200JSONResponse Item

func (response ItemCreate200JSONResponse) VisitItemCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ItemCreate400Response struct {
}

func (response ItemCreate400Response) VisitItemCreateResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type ItemCreate401Response = UnauthorizedErrorResponse

func (response ItemCreate401Response) VisitItemCreateResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type ItemPublishRequestObject struct {
	ModelId ModelIdParam `json:"modelId"`
	ItemId  ItemIdParam  `json:"itemId"`
	Body    *ItemPublishJSONRequestBody
}

type ItemPublishResponseObject interface {
	VisitItemPublishResponse(w http.ResponseWriter) error
}

type ItemPublish200JSONResponse Item

func (response ItemPublish200JSONResponse) VisitItemPublishResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ItemPublish400Response struct {
}

func (response ItemPublish400Response) VisitItemPublishResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type ItemPublish401Response = UnauthorizedErrorResponse

func (response ItemPublish401Response) VisitItemPublishResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type AssetFilterRequestObject struct {
	ProjectId ProjectIdParam `json:"projectId"`
	Params    AssetFilterParams
}

type AssetFilterResponseObject interface {
	VisitAssetFilterResponse(w http.ResponseWriter) error
}

type AssetFilter200JSONResponse struct {
	Items      *[]Asset `json:"items,omitempty"`
	Page       *int     `json:"page,omitempty"`
	PerPage    *int     `json:"perPage,omitempty"`
	TotalCount *int     `json:"totalCount,omitempty"`
}

func (response AssetFilter200JSONResponse) VisitAssetFilterResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AssetFilter400Response struct {
}

func (response AssetFilter400Response) VisitAssetFilterResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type AssetFilter401Response = UnauthorizedErrorResponse

func (response AssetFilter401Response) VisitAssetFilterResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type AssetCreateRequestObject struct {
	ProjectId ProjectIdParam `json:"projectId"`
	Body      *multipart.Reader
}

type AssetCreateResponseObject interface {
	VisitAssetCreateResponse(w http.ResponseWriter) error
}

type AssetCreate200JSONResponse Asset

func (response AssetCreate200JSONResponse) VisitAssetCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AssetCreate400Response struct {
}

func (response AssetCreate400Response) VisitAssetCreateResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type AssetCreate401Response = UnauthorizedErrorResponse

func (response AssetCreate401Response) VisitAssetCreateResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type AssetDeleteRequestObject struct {
	ProjectId ProjectIdParam `json:"projectId"`
	AssetId   AssetIdParam   `json:"assetId"`
}

type AssetDeleteResponseObject interface {
	VisitAssetDeleteResponse(w http.ResponseWriter) error
}

type AssetDelete200JSONResponse struct {
	Id *openapi_types.UUID `json:"id,omitempty"`
}

func (response AssetDelete200JSONResponse) VisitAssetDeleteResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AssetDelete400Response struct {
}

func (response AssetDelete400Response) VisitAssetDeleteResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type AssetDelete401Response = UnauthorizedErrorResponse

func (response AssetDelete401Response) VisitAssetDeleteResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type AssetGetRequestObject struct {
	ProjectId ProjectIdParam `json:"projectId"`
	AssetId   AssetIdParam   `json:"assetId"`
}

type AssetGetResponseObject interface {
	VisitAssetGetResponse(w http.ResponseWriter) error
}

type AssetGet200JSONResponse Asset

func (response AssetGet200JSONResponse) VisitAssetGetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AssetGet400Response struct {
}

func (response AssetGet400Response) VisitAssetGetResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type AssetGet401Response = UnauthorizedErrorResponse

func (response AssetGet401Response) VisitAssetGetResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (DELETE /items/{itemId})
	ItemDelete(ctx context.Context, request ItemDeleteRequestObject) (ItemDeleteResponseObject, error)
	// Returns an items.
	// (GET /items/{itemId})
	ItemGet(ctx context.Context, request ItemGetRequestObject) (ItemGetResponseObject, error)
	// Returns a list of items.
	// (GET /models/{modelId}/items)
	ItemFilter(ctx context.Context, request ItemFilterRequestObject) (ItemFilterResponseObject, error)

	// (POST /models/{modelId}/items)
	ItemCreate(ctx context.Context, request ItemCreateRequestObject) (ItemCreateResponseObject, error)
	// Set ref and version.
	// (POST /models/{modelId}/items/{itemId}/refs)
	ItemPublish(ctx context.Context, request ItemPublishRequestObject) (ItemPublishResponseObject, error)
	// Returns a list of assets.
	// (GET /projects/{projectId}/assets)
	AssetFilter(ctx context.Context, request AssetFilterRequestObject) (AssetFilterResponseObject, error)
	// Returns a list of assets.
	// (POST /projects/{projectId}/assets)
	AssetCreate(ctx context.Context, request AssetCreateRequestObject) (AssetCreateResponseObject, error)

	// (DELETE /projects/{projectId}/assets/{assetId})
	AssetDelete(ctx context.Context, request AssetDeleteRequestObject) (AssetDeleteResponseObject, error)

	// (GET /projects/{projectId}/assets/{assetId})
	AssetGet(ctx context.Context, request AssetGetRequestObject) (AssetGetResponseObject, error)
}

type StrictHandlerFunc func(ctx echo.Context, args interface{}) (interface{}, error)

type StrictMiddlewareFunc func(f StrictHandlerFunc, operationID string) StrictHandlerFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// ItemDelete operation middleware
func (sh *strictHandler) ItemDelete(ctx echo.Context, itemId ItemIdParam) error {
	var request ItemDeleteRequestObject

	request.ItemId = itemId

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.ItemDelete(ctx.Request().Context(), request.(ItemDeleteRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ItemDelete")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(ItemDeleteResponseObject); ok {
		return validResponse.VisitItemDeleteResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// ItemGet operation middleware
func (sh *strictHandler) ItemGet(ctx echo.Context, itemId ItemIdParam, params ItemGetParams) error {
	var request ItemGetRequestObject

	request.ItemId = itemId
	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.ItemGet(ctx.Request().Context(), request.(ItemGetRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ItemGet")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(ItemGetResponseObject); ok {
		return validResponse.VisitItemGetResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// ItemFilter operation middleware
func (sh *strictHandler) ItemFilter(ctx echo.Context, modelId ModelIdParam, params ItemFilterParams) error {
	var request ItemFilterRequestObject

	request.ModelId = modelId
	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.ItemFilter(ctx.Request().Context(), request.(ItemFilterRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ItemFilter")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(ItemFilterResponseObject); ok {
		return validResponse.VisitItemFilterResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// ItemCreate operation middleware
func (sh *strictHandler) ItemCreate(ctx echo.Context, modelId ModelIdParam) error {
	var request ItemCreateRequestObject

	request.ModelId = modelId

	var body ItemCreateJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.ItemCreate(ctx.Request().Context(), request.(ItemCreateRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ItemCreate")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(ItemCreateResponseObject); ok {
		return validResponse.VisitItemCreateResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// ItemPublish operation middleware
func (sh *strictHandler) ItemPublish(ctx echo.Context, modelId ModelIdParam, itemId ItemIdParam) error {
	var request ItemPublishRequestObject

	request.ModelId = modelId
	request.ItemId = itemId

	var body ItemPublishJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.ItemPublish(ctx.Request().Context(), request.(ItemPublishRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ItemPublish")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(ItemPublishResponseObject); ok {
		return validResponse.VisitItemPublishResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// AssetFilter operation middleware
func (sh *strictHandler) AssetFilter(ctx echo.Context, projectId ProjectIdParam, params AssetFilterParams) error {
	var request AssetFilterRequestObject

	request.ProjectId = projectId
	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AssetFilter(ctx.Request().Context(), request.(AssetFilterRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AssetFilter")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(AssetFilterResponseObject); ok {
		return validResponse.VisitAssetFilterResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// AssetCreate operation middleware
func (sh *strictHandler) AssetCreate(ctx echo.Context, projectId ProjectIdParam) error {
	var request AssetCreateRequestObject

	request.ProjectId = projectId

	if reader, err := ctx.Request().MultipartReader(); err != nil {
		return err
	} else {
		request.Body = reader
	}

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AssetCreate(ctx.Request().Context(), request.(AssetCreateRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AssetCreate")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(AssetCreateResponseObject); ok {
		return validResponse.VisitAssetCreateResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// AssetDelete operation middleware
func (sh *strictHandler) AssetDelete(ctx echo.Context, projectId ProjectIdParam, assetId AssetIdParam) error {
	var request AssetDeleteRequestObject

	request.ProjectId = projectId
	request.AssetId = assetId

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AssetDelete(ctx.Request().Context(), request.(AssetDeleteRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AssetDelete")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(AssetDeleteResponseObject); ok {
		return validResponse.VisitAssetDeleteResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// AssetGet operation middleware
func (sh *strictHandler) AssetGet(ctx echo.Context, projectId ProjectIdParam, assetId AssetIdParam) error {
	var request AssetGetRequestObject

	request.ProjectId = projectId
	request.AssetId = assetId

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AssetGet(ctx.Request().Context(), request.(AssetGetRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AssetGet")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(AssetGetResponseObject); ok {
		return validResponse.VisitAssetGetResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+RaTXPbNhD9Kxi0R1qU6/SimxqnHXcmiaduevH4AJFLCQkJMAtQtqPhf+9gQVGUREqk",
	"rY7j9NJaJD7e7j7sPiyz4pHOcq1AWcMnK54LFBlYQPoljAF7FV+7h+53DCZCmVupFZ/wq0umE2YXwAyk",
	"EFmIGU3gAZfufS7sggdciQz4ZL0WDzjC10IixHxisYCAm2gBmXDrJxozYfmEF4V0I+1j7qYai1LNecAf",
	"zub6rHo4rdYry4BLC9kQmG58O0q/0slAXvnlHMZMx5AOAUkT2L20C6m23+SoP0PU4edqn5OZ8L5az9mQ",
	"izl0GPDJQMysrkASXjd6jfFrAfi4AVm92iCKIRFFavnkPOCZVDIrMvq7AiGVhTmgBwF4fTIcfq12KL+O",
	"A56JhwrLeHwcmY/LkCBXU1rDfK/xi8lFBO2Brnc7Waiv6xWdNQhJPx8LhpAwjWwJ2OFnhKTdxzwVFoyj",
	"Mijn2NvNg7yYpTLid7uYHTaj0V5KPIIvhkQqIKdqjAFZLBEiN2gdBwSTa2WApdLYgN3LNGUzYHKuNLo8",
	"kTQmS8OUtixHMKAsxB2mxhI7THUgG4YK+kUPO20camCbWR043fIdQCMEYSGeNsPSfFbkcfV3C3Bijt+e",
	"asgnJQq70Ci/QfwOUeO+OdMoAmOY1V9AOTdn0hip5o5SUi1FKj0hPdRNYaJ6hToHtNLvFWllQdm/CdJq",
	"F1rQMKJ5MpwxeyejDHgiU1rmZ8feCf8p3JTKsMIS0hhXgeLjh61c+74FWY6wlHC/Rr72usx8dpqD9v+9",
	"cMtSir+IW7zfyEG98FhtRXojvzVBqSKbuYTWjPNxd5X1Ez2j4kQOhDTej5L3FTyILHcO5jzoAXTHMRYe",
	"HBPd/6YIggeqSNM2fyxFWsDWfuc9sfvg7xBsIdMYQZEZFjLTlx/V+gJRPBITjzC1kyimM1iYtkxoM43U",
	"z55pAqOFXEK8lQoSkRqol5hpnYJQTzlJkMZmgNMcb1q81vOUrSVQn7G5wLX2rbH1ZOMGF0KyvcLRGUOO",
	"VsCXgIZS5VFsbfFGSD7iP4D7MacYrAaUXoLyJBiOuhAVKO3jjQuzBzADgYDTwi7cL4o/EY0eb5ZdWJv7",
	"0iJVovcLyF/wTqBdnL19f8OunCBDQVV+en3lFpGW8szhUbWP+floPBo7W3UOSuSST/jFaDy6cI4RdkHA",
	"Qwp1uPL3hdJDSsHS2XQupqUdB+kOcOnf7ZTGX8bjRtWiQ5jnqYxobvjZ+IhvKnRbGh0eiDLY8Z4HzoTy",
	"96Iy4G88sB316ksxc1oTjGX1VZH5HEvzzrvOdm14uC8IaOab/R0/aMsSXah4iz58crtNnNu78q501dG2",
	"McMWqMzauBEPWqLzB91Zm1ff23YjNkPCWh2Xd8+M6qFM6AOyH7Lpy8XqWCQCboosE/i4730zqjLuAD83",
	"r/YU5pCyuwlXVZYvwzrrHmYAiWGnkGk83bhYIlML7qQwoWKvoaWat7Pkdxo7mCgbGV8GvQbX95oe4zcX",
	"8j6DmxfnHuNPR/GdxLWOVy8xsOb5bgWlW/1kdfBOXjcLjg8kJfxWF96ceuy49aJ/NKVO2Z83Hz8wAuso",
	"VxhA5jSdeTVHdvu4POHkbjW83Ca5Nra9PL4lNVn1McDY33T8+AyCnUZutsd5u9NSvkTuf1XcKrtTdq2e",
	"wrV+fg7BgoGlZMPHbV/dgKV+lisIlSZkicb9BnJ7lbh20tksnsHlQ5yopfz3wsXXoUNaYjqiFcKqXWLC",
	"Vd04KUPqLg1RFH7CMElB3y9+ME3xkjrBtwRfm1ComEN92lenDDz4J0iDnQ8l3cn4hEeth8bIitTKXKAN",
	"3a36LBZWHJYZvkVYX8FnUglqsfe4hP+X2bo6Cj8q247k7XBVfWo+2JUhTnzXbZmXDtampdLiOd8y+f8x",
	"+Nmp7nhJ3fpHF15AG8Dler9tWzMhFfOvSaUK+pTmUqLTq1FmmGx0OkUuefWlgHqqJgyjzIwQQKBdjKQO",
	"RS7D5bnb9t8AAAD//4CNOJQaIgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}