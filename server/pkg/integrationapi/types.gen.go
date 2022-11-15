// Package integrationapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.2 DO NOT EDIT.
package integrationapi

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/reearth/reearth-cms/server/pkg/id"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for AssetPreviewType.
const (
	Geo     AssetPreviewType = "geo"
	Geo3d   AssetPreviewType = "geo3d"
	Image   AssetPreviewType = "image"
	Model3d AssetPreviewType = "model3d"
	Unknown AssetPreviewType = "unknown"
)

// Defines values for CommentAuthorType.
const (
	Integrtaion CommentAuthorType = "integrtaion"
	User        CommentAuthorType = "user"
)

// Defines values for RefOrVersionRef.
const (
	RefOrVersionRefLatest RefOrVersionRef = "latest"
	RefOrVersionRefPublic RefOrVersionRef = "public"
)

// Defines values for ValueType.
const (
	ValueTypeAsset     ValueType = "asset"
	ValueTypeBool      ValueType = "bool"
	ValueTypeDate      ValueType = "date"
	ValueTypeInteger   ValueType = "integer"
	ValueTypeMarkdown  ValueType = "markdown"
	ValueTypeReference ValueType = "reference"
	ValueTypeRichText  ValueType = "richText"
	ValueTypeSelect    ValueType = "select"
	ValueTypeTag       ValueType = "tag"
	ValueTypeText      ValueType = "text"
	ValueTypeTextArea  ValueType = "textArea"
	ValueTypeUrl       ValueType = "url"
)

// Defines values for RefParam.
const (
	RefParamLatest RefParam = "latest"
	RefParamPublic RefParam = "public"
)

// Defines values for SortDirParam.
const (
	SortDirParamAsc  SortDirParam = "asc"
	SortDirParamDesc SortDirParam = "desc"
)

// Defines values for SortParam.
const (
	SortParamCreatedAt SortParam = "createdAt"
	SortParamUpdatedAt SortParam = "updatedAt"
)

// Defines values for ItemGetParamsRef.
const (
	ItemGetParamsRefLatest ItemGetParamsRef = "latest"
	ItemGetParamsRefPublic ItemGetParamsRef = "public"
)

// Defines values for ItemFilterParamsSort.
const (
	ItemFilterParamsSortCreatedAt ItemFilterParamsSort = "createdAt"
	ItemFilterParamsSortUpdatedAt ItemFilterParamsSort = "updatedAt"
)

// Defines values for ItemFilterParamsDir.
const (
	ItemFilterParamsDirAsc  ItemFilterParamsDir = "asc"
	ItemFilterParamsDirDesc ItemFilterParamsDir = "desc"
)

// Defines values for ItemFilterParamsRef.
const (
	ItemFilterParamsRefLatest ItemFilterParamsRef = "latest"
	ItemFilterParamsRefPublic ItemFilterParamsRef = "public"
)

// Defines values for AssetFilterParamsSort.
const (
	CreatedAt AssetFilterParamsSort = "createdAt"
	UpdatedAt AssetFilterParamsSort = "updatedAt"
)

// Defines values for AssetFilterParamsDir.
const (
	Asc  AssetFilterParamsDir = "asc"
	Desc AssetFilterParamsDir = "desc"
)

// Asset defines model for asset.
type Asset struct {
	ContentType *string             `json:"contentType,omitempty"`
	CreatedAt   *openapi_types.Date `json:"createdAt,omitempty"`
	File        *File               `json:"file,omitempty"`
	Id          *id.AssetID         `json:"id,omitempty"`
	Name        *string             `json:"name,omitempty"`
	PreviewType *AssetPreviewType   `json:"previewType,omitempty"`
	ProjectId   *id.ProjectID       `json:"projectId,omitempty"`
	TotalSize   *float32            `json:"totalSize,omitempty"`
	UpdatedAt   *openapi_types.Date `json:"updatedAt,omitempty"`
	Url         *string             `json:"url,omitempty"`
}

// AssetPreviewType defines model for Asset.PreviewType.
type AssetPreviewType string

// Comment defines model for comment.
type Comment struct {
	AuthorId   *any                `json:"authorId,omitempty"`
	AuthorType *CommentAuthorType  `json:"authorType,omitempty"`
	Content    *string             `json:"content,omitempty"`
	CreatedAt  *openapi_types.Date `json:"createdAt,omitempty"`
	Id         *id.CommentID       `json:"id,omitempty"`
}

// CommentAuthorType defines model for Comment.AuthorType.
type CommentAuthorType string

// Field defines model for field.
type Field struct {
	Id    *id.FieldID  `json:"id,omitempty"`
	Type  *ValueType   `json:"type,omitempty"`
	Value *interface{} `json:"value,omitempty"`
}

// File defines model for file.
type File struct {
	Children    *[]File  `json:"children,omitempty"`
	ContentType *string  `json:"contentType,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Path        *string  `json:"path,omitempty"`
	Size        *float32 `json:"size,omitempty"`
}

// Item defines model for item.
type Item struct {
	CreatedAt *openapi_types.Date `json:"createdAt,omitempty"`
	Fields    *[]Field            `json:"fields,omitempty"`
	Id        *id.ItemID          `json:"id,omitempty"`
	ModelId   *string             `json:"modelId,omitempty"`
	UpdatedAt *openapi_types.Date `json:"updatedAt,omitempty"`
}

// RefOrVersion defines model for refOrVersion.
type RefOrVersion struct {
	Ref     *RefOrVersionRef    `json:"ref,omitempty"`
	Version *openapi_types.UUID `json:"version,omitempty"`
}

// RefOrVersionRef defines model for RefOrVersion.Ref.
type RefOrVersionRef string

// Schema defines model for schema.
type Schema struct {
	CreatedAt *openapi_types.Date `json:"createdAt,omitempty"`
	Fields    *[]SchemaField      `json:"fields,omitempty"`
	Id        *id.SchemaID        `json:"id,omitempty"`
	ProjectId *id.ProjectID       `json:"projectId,omitempty"`
	UpdatedAt *openapi_types.Date `json:"updatedAt,omitempty"`
}

// SchemaField defines model for schemaField.
type SchemaField struct {
	Id       *id.FieldID `json:"id,omitempty"`
	Key      *string     `json:"key,omitempty"`
	Required *bool       `json:"required,omitempty"`
	Type     *ValueType  `json:"type,omitempty"`
}

// ValueType defines model for valueType.
type ValueType string

// Version defines model for version.
type Version struct {
	Parents *[]openapi_types.UUID `json:"parents,omitempty"`
	Refs    *[]openapi_types.UUID `json:"refs,omitempty"`
	Version *openapi_types.UUID   `json:"version,omitempty"`
}

// VersionedItem defines model for versionedItem.
type VersionedItem struct {
	CreatedAt *openapi_types.Date   `json:"createdAt,omitempty"`
	Fields    *[]Field              `json:"fields,omitempty"`
	Id        *id.ItemID            `json:"id,omitempty"`
	ModelId   *string               `json:"modelId,omitempty"`
	Parents   *[]openapi_types.UUID `json:"parents,omitempty"`
	Refs      *[]string             `json:"refs,omitempty"`
	UpdatedAt *openapi_types.Date   `json:"updatedAt,omitempty"`
	Version   *openapi_types.UUID   `json:"version,omitempty"`
}

// AssetIdParam defines model for assetIdParam.
type AssetIdParam = id.AssetID

// CommentIdParam defines model for commentIdParam.
type CommentIdParam = id.CommentID

// ItemIdParam defines model for itemIdParam.
type ItemIdParam = id.ItemID

// ModelIdParam defines model for modelIdParam.
type ModelIdParam = id.ModelID

// PageParam defines model for pageParam.
type PageParam = int

// PerPageParam defines model for perPageParam.
type PerPageParam = int

// ProjectIdParam defines model for projectIdParam.
type ProjectIdParam = id.ProjectID

// RefParam defines model for refParam.
type RefParam string

// SortDirParam defines model for sortDirParam.
type SortDirParam string

// SortParam defines model for sortParam.
type SortParam string

// ItemGetParams defines parameters for ItemGet.
type ItemGetParams struct {
	// Ref Used to select a ref or ver
	Ref *ItemGetParamsRef `form:"ref,omitempty" json:"ref,omitempty"`
}

// ItemGetParamsRef defines parameters for ItemGet.
type ItemGetParamsRef string

// ItemUpdateJSONBody defines parameters for ItemUpdate.
type ItemUpdateJSONBody struct {
	Fields *[]Field `json:"fields,omitempty"`
}

// ItemFilterParams defines parameters for ItemFilter.
type ItemFilterParams struct {
	// Sort Used to define the order of the response list
	Sort *ItemFilterParamsSort `form:"sort,omitempty" json:"sort,omitempty"`

	// Dir Used to define the order direction of the response list, will be ignored if the order is not presented
	Dir *ItemFilterParamsDir `form:"dir,omitempty" json:"dir,omitempty"`

	// Page Used to select the page
	Page *PageParam `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Used to select the page
	PerPage *PerPageParam `form:"perPage,omitempty" json:"perPage,omitempty"`

	// Ref Used to select a ref or ver
	Ref *ItemFilterParamsRef `form:"ref,omitempty" json:"ref,omitempty"`
}

// ItemFilterParamsSort defines parameters for ItemFilter.
type ItemFilterParamsSort string

// ItemFilterParamsDir defines parameters for ItemFilter.
type ItemFilterParamsDir string

// ItemFilterParamsRef defines parameters for ItemFilter.
type ItemFilterParamsRef string

// ItemCreateJSONBody defines parameters for ItemCreate.
type ItemCreateJSONBody struct {
	Fields *[]struct {
		Id    *id.FieldID  `json:"id,omitempty"`
		Type  *ValueType   `json:"type,omitempty"`
		Value *interface{} `json:"value,omitempty"`
	} `json:"fields,omitempty"`
}

// AssetFilterParams defines parameters for AssetFilter.
type AssetFilterParams struct {
	// Sort Used to define the order of the response list
	Sort *AssetFilterParamsSort `form:"sort,omitempty" json:"sort,omitempty"`

	// Dir Used to define the order direction of the response list, will be ignored if the order is not presented
	Dir *AssetFilterParamsDir `form:"dir,omitempty" json:"dir,omitempty"`

	// Page Used to select the page
	Page *PageParam `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Used to select the page
	PerPage *PerPageParam `form:"perPage,omitempty" json:"perPage,omitempty"`
}

// AssetFilterParamsSort defines parameters for AssetFilter.
type AssetFilterParamsSort string

// AssetFilterParamsDir defines parameters for AssetFilter.
type AssetFilterParamsDir string

// AssetCreateJSONBody defines parameters for AssetCreate.
type AssetCreateJSONBody struct {
	Url *string `json:"url,omitempty"`
}

// AssetCreateMultipartBody defines parameters for AssetCreate.
type AssetCreateMultipartBody struct {
	File *openapi_types.File `json:"file,omitempty"`
}

// AssetCommentCreateJSONBody defines parameters for AssetCommentCreate.
type AssetCommentCreateJSONBody struct {
	Content *string `json:"content,omitempty"`
}

// AssetCommentUpdateJSONBody defines parameters for AssetCommentUpdate.
type AssetCommentUpdateJSONBody struct {
	Content *string `json:"content,omitempty"`
}

// ItemUpdateJSONRequestBody defines body for ItemUpdate for application/json ContentType.
type ItemUpdateJSONRequestBody ItemUpdateJSONBody

// ItemCreateJSONRequestBody defines body for ItemCreate for application/json ContentType.
type ItemCreateJSONRequestBody ItemCreateJSONBody

// AssetCreateJSONRequestBody defines body for AssetCreate for application/json ContentType.
type AssetCreateJSONRequestBody AssetCreateJSONBody

// AssetCreateMultipartRequestBody defines body for AssetCreate for multipart/form-data ContentType.
type AssetCreateMultipartRequestBody AssetCreateMultipartBody

// AssetCommentCreateJSONRequestBody defines body for AssetCommentCreate for application/json ContentType.
type AssetCommentCreateJSONRequestBody AssetCommentCreateJSONBody

// AssetCommentUpdateJSONRequestBody defines body for AssetCommentUpdate for application/json ContentType.
type AssetCommentUpdateJSONRequestBody AssetCommentUpdateJSONBody
