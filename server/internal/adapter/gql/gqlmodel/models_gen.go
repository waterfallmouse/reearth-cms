// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlmodel

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/reearth/reearth-cms/server/internal/usecase"
	"golang.org/x/text/language"
)

type Node interface {
	IsNode()
}

type SchemaFieldTypeProperty interface {
	IsSchemaFieldTypeProperty()
}

type AddMemberToWorkspaceInput struct {
	WorkspaceID ID   `json:"workspaceId"`
	UserID      ID   `json:"userId"`
	Role        Role `json:"role"`
}

type AddMemberToWorkspacePayload struct {
	Workspace *Workspace `json:"workspace"`
}

type Asset struct {
	ID          ID           `json:"id"`
	Project     *Project     `json:"project"`
	ProjectID   ID           `json:"projectId"`
	CreatedAt   time.Time    `json:"createdAt"`
	CreatedBy   *User        `json:"createdBy"`
	CreatedByID ID           `json:"createdById"`
	FileName    string       `json:"fileName"`
	Size        int64        `json:"size"`
	PreviewType *PreviewType `json:"previewType"`
	File        *AssetFile   `json:"file"`
	Hash        string       `json:"hash"`
}

func (Asset) IsNode() {}

type AssetConnection struct {
	Edges      []*AssetEdge `json:"edges"`
	Nodes      []*Asset     `json:"nodes"`
	PageInfo   *PageInfo    `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

type AssetEdge struct {
	Cursor usecase.Cursor `json:"cursor"`
	Node   *Asset         `json:"node"`
}

type AssetFile struct {
	Name        string       `json:"name"`
	Size        int64        `json:"size"`
	ContentType *string      `json:"contentType"`
	Path        string       `json:"path"`
	Children    []*AssetFile `json:"children"`
}

type CreateAssetInput struct {
	ProjectID ID             `json:"projectId"`
	File      graphql.Upload `json:"file"`
}

type CreateAssetPayload struct {
	Asset *Asset `json:"asset"`
}

type CreateFieldInput struct {
	ModelID      ID                            `json:"modelId"`
	Type         SchemaFiledType               `json:"type"`
	Title        string                        `json:"title"`
	Description  *string                       `json:"description"`
	Key          string                        `json:"key"`
	MultiValue   bool                          `json:"multiValue"`
	Unique       bool                          `json:"unique"`
	Required     bool                          `json:"required"`
	TypeProperty *SchemaFieldTypePropertyInput `json:"typeProperty"`
}

type CreateItemInput struct {
	ModelID ID                `json:"modelId"`
	Fields  []*ItemFieldInput `json:"fields"`
}

type CreateModelInput struct {
	ProjectID   ID      `json:"projectId"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Key         *string `json:"key"`
}

type CreateProjectInput struct {
	WorkspaceID ID      `json:"workspaceId"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Alias       *string `json:"alias"`
}

type CreateWorkspaceInput struct {
	Name string `json:"name"`
}

type CreateWorkspacePayload struct {
	Workspace *Workspace `json:"workspace"`
}

type DeleteAssetInput struct {
	AssetID ID `json:"assetId"`
}

type DeleteAssetPayload struct {
	AssetID ID `json:"assetId"`
}

type DeleteFieldInput struct {
	FieldID ID `json:"fieldId"`
}

type DeleteFieldPayload struct {
	FieldID ID `json:"fieldId"`
}

type DeleteItemInput struct {
	ItemID ID `json:"itemId"`
}

type DeleteItemPayload struct {
	ItemID ID `json:"itemId"`
}

type DeleteMeInput struct {
	UserID ID `json:"userId"`
}

type DeleteMePayload struct {
	UserID ID `json:"userId"`
}

type DeleteModelInput struct {
	ModelID ID `json:"modelId"`
}

type DeleteModelPayload struct {
	ModelID ID `json:"modelId"`
}

type DeleteProjectInput struct {
	ProjectID ID `json:"projectId"`
}

type DeleteProjectPayload struct {
	ProjectID ID `json:"projectId"`
}

type DeleteWorkspaceInput struct {
	WorkspaceID ID `json:"workspaceId"`
}

type DeleteWorkspacePayload struct {
	WorkspaceID ID `json:"workspaceId"`
}

type FieldPayload struct {
	Field *SchemaField `json:"field"`
}

type Item struct {
	ID            ID             `json:"id"`
	ModelID       ID             `json:"modelId"`
	Model         *Model         `json:"model"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	LatestVersion *ItemVersion   `json:"latestVersion"`
	PublicVersion string         `json:"publicVersion"`
	Versions      []*ItemVersion `json:"versions"`
}

func (Item) IsNode() {}

type ItemConnection struct {
	Edges      []*ItemEdge `json:"edges"`
	Nodes      []*Item     `json:"nodes"`
	PageInfo   *PageInfo   `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

type ItemEdge struct {
	Cursor usecase.Cursor `json:"cursor"`
	Node   *Item          `json:"node"`
}

type ItemField struct {
	FieldID ID          `json:"fieldId"`
	Value   interface{} `json:"value"`
}

type ItemFieldInput struct {
	FieldID ID          `json:"fieldId"`
	Value   interface{} `json:"value"`
}

type ItemPayload struct {
	Item *Item `json:"item"`
}

type ItemVersion struct {
	Version string       `json:"version"`
	Parent  []string     `json:"parent"`
	Ref     []string     `json:"ref"`
	Fields  []*ItemField `json:"fields"`
}

type KeyAvailability struct {
	Key       string `json:"key"`
	Available bool   `json:"available"`
}

type Me struct {
	ID            ID           `json:"id"`
	Name          string       `json:"name"`
	Email         string       `json:"email"`
	Lang          language.Tag `json:"lang"`
	Theme         Theme        `json:"theme"`
	MyWorkspaceID ID           `json:"myWorkspaceId"`
	Auths         []string     `json:"auths"`
	Workspaces    []*Workspace `json:"workspaces"`
	MyWorkspace   *Workspace   `json:"myWorkspace"`
}

type Model struct {
	ID          ID        `json:"id"`
	ProjectID   ID        `json:"projectId"`
	SchemaID    ID        `json:"schemaId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Key         string    `json:"key"`
	Project     *Project  `json:"project"`
	Schema      *Schema   `json:"schema"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (Model) IsNode() {}

type ModelConnection struct {
	Edges      []*ModelEdge `json:"edges"`
	Nodes      []*Model     `json:"nodes"`
	PageInfo   *PageInfo    `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

type ModelEdge struct {
	Cursor usecase.Cursor `json:"cursor"`
	Node   *Model         `json:"node"`
}

type ModelPayload struct {
	Model *Model `json:"model"`
}

type PageInfo struct {
	StartCursor     *usecase.Cursor `json:"startCursor"`
	EndCursor       *usecase.Cursor `json:"endCursor"`
	HasNextPage     bool            `json:"hasNextPage"`
	HasPreviousPage bool            `json:"hasPreviousPage"`
}

type Pagination struct {
	First  *int            `json:"first"`
	Last   *int            `json:"last"`
	After  *usecase.Cursor `json:"after"`
	Before *usecase.Cursor `json:"before"`
}

type Project struct {
	ID          ID         `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Alias       string     `json:"alias"`
	WorkspaceID ID         `json:"workspaceId"`
	Workspace   *Workspace `json:"workspace"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func (Project) IsNode() {}

type ProjectAliasAvailability struct {
	Alias     string `json:"alias"`
	Available bool   `json:"available"`
}

type ProjectConnection struct {
	Edges      []*ProjectEdge `json:"edges"`
	Nodes      []*Project     `json:"nodes"`
	PageInfo   *PageInfo      `json:"pageInfo"`
	TotalCount int            `json:"totalCount"`
}

type ProjectEdge struct {
	Cursor usecase.Cursor `json:"cursor"`
	Node   *Project       `json:"node"`
}

type ProjectPayload struct {
	Project *Project `json:"project"`
}

type PublishModelInput struct {
	ModelID ID   `json:"modelId"`
	Status  bool `json:"status"`
}

type PublishModelPayload struct {
	ModelID ID   `json:"modelId"`
	Status  bool `json:"status"`
}

type RemoveMemberFromWorkspaceInput struct {
	WorkspaceID ID `json:"workspaceId"`
	UserID      ID `json:"userId"`
}

type RemoveMemberFromWorkspacePayload struct {
	Workspace *Workspace `json:"workspace"`
}

type RemoveMyAuthInput struct {
	Auth string `json:"auth"`
}

type Schema struct {
	ID        ID             `json:"id"`
	ProjectID ID             `json:"projectId"`
	Fields    []*SchemaField `json:"fields"`
	Project   *Project       `json:"project"`
}

func (Schema) IsNode() {}

type SchemaField struct {
	ID           ID                      `json:"id"`
	ModelID      ID                      `json:"modelId"`
	Model        *Model                  `json:"model"`
	Type         SchemaFiledType         `json:"type"`
	TypeProperty SchemaFieldTypeProperty `json:"typeProperty"`
	Key          string                  `json:"key"`
	Title        string                  `json:"title"`
	Description  *string                 `json:"description"`
	MultiValue   bool                    `json:"multiValue"`
	Unique       bool                    `json:"unique"`
	Required     bool                    `json:"required"`
	CreatedAt    time.Time               `json:"createdAt"`
	UpdatedAt    time.Time               `json:"updatedAt"`
}

type SchemaFieldAsset struct {
	DefaultValue *ID `json:"defaultValue"`
}

func (SchemaFieldAsset) IsSchemaFieldTypeProperty() {}

type SchemaFieldAssetInput struct {
	DefaultValue *ID `json:"defaultValue"`
}

type SchemaFieldBool struct {
	DefaultValue *bool `json:"defaultValue"`
}

func (SchemaFieldBool) IsSchemaFieldTypeProperty() {}

type SchemaFieldBoolInput struct {
	DefaultValue *bool `json:"defaultValue"`
}

type SchemaFieldDate struct {
	DefaultValue *time.Time `json:"defaultValue"`
}

func (SchemaFieldDate) IsSchemaFieldTypeProperty() {}

type SchemaFieldDateInput struct {
	DefaultValue *time.Time `json:"defaultValue"`
}

type SchemaFieldInteger struct {
	DefaultValue *int `json:"defaultValue"`
	Min          *int `json:"min"`
	Max          *int `json:"max"`
}

func (SchemaFieldInteger) IsSchemaFieldTypeProperty() {}

type SchemaFieldIntegerInput struct {
	DefaultValue *int `json:"defaultValue"`
	Min          *int `json:"min"`
	Max          *int `json:"max"`
}

type SchemaFieldReference struct {
	ModelID *ID `json:"modelId"`
}

func (SchemaFieldReference) IsSchemaFieldTypeProperty() {}

type SchemaFieldReferenceInput struct {
	ModelID *ID `json:"modelId"`
}

type SchemaFieldRichText struct {
	DefaultValue *string `json:"defaultValue"`
	MaxLength    *int    `json:"maxLength"`
}

func (SchemaFieldRichText) IsSchemaFieldTypeProperty() {}

type SchemaFieldRichTextInput struct {
	DefaultValue *string `json:"defaultValue"`
	MaxLength    *int    `json:"maxLength"`
}

type SchemaFieldSelect struct {
	Values       []string `json:"values"`
	DefaultValue *string  `json:"defaultValue"`
}

func (SchemaFieldSelect) IsSchemaFieldTypeProperty() {}

type SchemaFieldSelectInput struct {
	Values       []string `json:"values"`
	DefaultValue *string  `json:"defaultValue"`
}

type SchemaFieldTag struct {
	Values       []string `json:"values"`
	DefaultValue *string  `json:"defaultValue"`
}

func (SchemaFieldTag) IsSchemaFieldTypeProperty() {}

type SchemaFieldTagInput struct {
	Values       []string `json:"values"`
	DefaultValue *string  `json:"defaultValue"`
}

type SchemaFieldText struct {
	DefaultValue *string `json:"defaultValue"`
	MaxLength    *int    `json:"maxLength"`
}

func (SchemaFieldText) IsSchemaFieldTypeProperty() {}

type SchemaFieldTextArea struct {
	DefaultValue *string `json:"defaultValue"`
	MaxLength    *int    `json:"maxLength"`
}

func (SchemaFieldTextArea) IsSchemaFieldTypeProperty() {}

type SchemaFieldTextAreaInput struct {
	DefaultValue *string `json:"defaultValue"`
	MaxLength    *int    `json:"maxLength"`
}

type SchemaFieldTextInput struct {
	DefaultValue *string `json:"defaultValue"`
	MaxLength    *int    `json:"maxLength"`
}

type SchemaFieldTypePropertyInput struct {
	Text         *SchemaFieldTextInput      `json:"text"`
	TextArea     *SchemaFieldTextAreaInput  `json:"textArea"`
	RichText     *SchemaFieldRichTextInput  `json:"richText"`
	MarkdownText *SchemaMarkdownTextInput   `json:"markdownText"`
	Asset        *SchemaFieldAssetInput     `json:"asset"`
	Date         *SchemaFieldDateInput      `json:"date"`
	Bool         *SchemaFieldBoolInput      `json:"bool"`
	Select       *SchemaFieldSelectInput    `json:"select"`
	Tag          *SchemaFieldTagInput       `json:"tag"`
	Integer      *SchemaFieldIntegerInput   `json:"integer"`
	Reference    *SchemaFieldReferenceInput `json:"reference"`
	URL          *SchemaFieldURLInput       `json:"url"`
}

type SchemaFieldURL struct {
	DefaultValue *string `json:"defaultValue"`
}

func (SchemaFieldURL) IsSchemaFieldTypeProperty() {}

type SchemaFieldURLInput struct {
	DefaultValue *string `json:"defaultValue"`
}

type SchemaMarkdownText struct {
	DefaultValue *string `json:"defaultValue"`
	MaxLength    *int    `json:"maxLength"`
}

func (SchemaMarkdownText) IsSchemaFieldTypeProperty() {}

type SchemaMarkdownTextInput struct {
	DefaultValue *string `json:"defaultValue"`
	MaxLength    *int    `json:"maxLength"`
}

type SignupInput struct {
	Lang        *language.Tag `json:"lang"`
	Theme       *Theme        `json:"theme"`
	UserID      *ID           `json:"userId"`
	WorkspaceID *ID           `json:"workspaceId"`
	Secret      *string       `json:"secret"`
}

type SignupPayload struct {
	User      *User      `json:"user"`
	Workspace *Workspace `json:"workspace"`
}

type UpdateFieldInput struct {
	FieldID      ID                            `json:"fieldId"`
	Title        *string                       `json:"title"`
	Description  *string                       `json:"description"`
	Key          *string                       `json:"key"`
	TypeProperty *SchemaFieldTypePropertyInput `json:"typeProperty"`
}

type UpdateItemInput struct {
	ItemID ID                `json:"itemId"`
	Fields []*ItemFieldInput `json:"fields"`
}

type UpdateMeInput struct {
	Name                 *string       `json:"name"`
	Email                *string       `json:"email"`
	Lang                 *language.Tag `json:"lang"`
	Theme                *Theme        `json:"theme"`
	Password             *string       `json:"password"`
	PasswordConfirmation *string       `json:"passwordConfirmation"`
}

type UpdateMePayload struct {
	Me *Me `json:"me"`
}

type UpdateMemberOfWorkspaceInput struct {
	WorkspaceID ID   `json:"workspaceId"`
	UserID      ID   `json:"userId"`
	Role        Role `json:"role"`
}

type UpdateMemberOfWorkspacePayload struct {
	Workspace *Workspace `json:"workspace"`
}

type UpdateModelInput struct {
	ModelID     ID      `json:"modelId"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Key         *string `json:"key"`
}

type UpdateProjectInput struct {
	ProjectID   ID      `json:"projectId"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type UpdateWorkspaceInput struct {
	WorkspaceID ID     `json:"workspaceId"`
	Name        string `json:"name"`
}

type UpdateWorkspacePayload struct {
	Workspace *Workspace `json:"workspace"`
}

type User struct {
	ID    ID     `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (User) IsNode() {}

type Workspace struct {
	ID       ID                 `json:"id"`
	Name     string             `json:"name"`
	Members  []*WorkspaceMember `json:"members"`
	Personal bool               `json:"personal"`
}

func (Workspace) IsNode() {}

type WorkspaceMember struct {
	UserID ID    `json:"userId"`
	Role   Role  `json:"role"`
	User   *User `json:"user"`
}

type AssetSortType string

const (
	AssetSortTypeDate AssetSortType = "DATE"
	AssetSortTypeSize AssetSortType = "SIZE"
	AssetSortTypeName AssetSortType = "NAME"
)

var AllAssetSortType = []AssetSortType{
	AssetSortTypeDate,
	AssetSortTypeSize,
	AssetSortTypeName,
}

func (e AssetSortType) IsValid() bool {
	switch e {
	case AssetSortTypeDate, AssetSortTypeSize, AssetSortTypeName:
		return true
	}
	return false
}

func (e AssetSortType) String() string {
	return string(e)
}

func (e *AssetSortType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AssetSortType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AssetSortType", str)
	}
	return nil
}

func (e AssetSortType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type NodeType string

const (
	NodeTypeUser      NodeType = "USER"
	NodeTypeWorkspace NodeType = "WORKSPACE"
	NodeTypeProject   NodeType = "PROJECT"
)

var AllNodeType = []NodeType{
	NodeTypeUser,
	NodeTypeWorkspace,
	NodeTypeProject,
}

func (e NodeType) IsValid() bool {
	switch e {
	case NodeTypeUser, NodeTypeWorkspace, NodeTypeProject:
		return true
	}
	return false
}

func (e NodeType) String() string {
	return string(e)
}

func (e *NodeType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NodeType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NodeType", str)
	}
	return nil
}

func (e NodeType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PreviewType string

const (
	PreviewTypeImage   PreviewType = "IMAGE"
	PreviewTypeGeo     PreviewType = "GEO"
	PreviewTypeGeo3d   PreviewType = "GEO3D"
	PreviewTypeModel3d PreviewType = "MODEL3D"
)

var AllPreviewType = []PreviewType{
	PreviewTypeImage,
	PreviewTypeGeo,
	PreviewTypeGeo3d,
	PreviewTypeModel3d,
}

func (e PreviewType) IsValid() bool {
	switch e {
	case PreviewTypeImage, PreviewTypeGeo, PreviewTypeGeo3d, PreviewTypeModel3d:
		return true
	}
	return false
}

func (e PreviewType) String() string {
	return string(e)
}

func (e *PreviewType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PreviewType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PreviewType", str)
	}
	return nil
}

func (e PreviewType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Role string

const (
	RoleReader Role = "READER"
	RoleWriter Role = "WRITER"
	RoleOwner  Role = "OWNER"
)

var AllRole = []Role{
	RoleReader,
	RoleWriter,
	RoleOwner,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleReader, RoleWriter, RoleOwner:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SchemaFiledType string

const (
	SchemaFiledTypeText         SchemaFiledType = "Text"
	SchemaFiledTypeTextArea     SchemaFiledType = "TextArea"
	SchemaFiledTypeRichText     SchemaFiledType = "RichText"
	SchemaFiledTypeMarkdownText SchemaFiledType = "MarkdownText"
	SchemaFiledTypeAsset        SchemaFiledType = "Asset"
	SchemaFiledTypeDate         SchemaFiledType = "Date"
	SchemaFiledTypeBool         SchemaFiledType = "Bool"
	SchemaFiledTypeSelect       SchemaFiledType = "Select"
	SchemaFiledTypeTag          SchemaFiledType = "Tag"
	SchemaFiledTypeInteger      SchemaFiledType = "Integer"
	SchemaFiledTypeReference    SchemaFiledType = "Reference"
	SchemaFiledTypeURL          SchemaFiledType = "URL"
)

var AllSchemaFiledType = []SchemaFiledType{
	SchemaFiledTypeText,
	SchemaFiledTypeTextArea,
	SchemaFiledTypeRichText,
	SchemaFiledTypeMarkdownText,
	SchemaFiledTypeAsset,
	SchemaFiledTypeDate,
	SchemaFiledTypeBool,
	SchemaFiledTypeSelect,
	SchemaFiledTypeTag,
	SchemaFiledTypeInteger,
	SchemaFiledTypeReference,
	SchemaFiledTypeURL,
}

func (e SchemaFiledType) IsValid() bool {
	switch e {
	case SchemaFiledTypeText, SchemaFiledTypeTextArea, SchemaFiledTypeRichText, SchemaFiledTypeMarkdownText, SchemaFiledTypeAsset, SchemaFiledTypeDate, SchemaFiledTypeBool, SchemaFiledTypeSelect, SchemaFiledTypeTag, SchemaFiledTypeInteger, SchemaFiledTypeReference, SchemaFiledTypeURL:
		return true
	}
	return false
}

func (e SchemaFiledType) String() string {
	return string(e)
}

func (e *SchemaFiledType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SchemaFiledType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SchemaFiledType", str)
	}
	return nil
}

func (e SchemaFiledType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Theme string

const (
	ThemeDefault Theme = "DEFAULT"
	ThemeLight   Theme = "LIGHT"
	ThemeDark    Theme = "DARK"
)

var AllTheme = []Theme{
	ThemeDefault,
	ThemeLight,
	ThemeDark,
}

func (e Theme) IsValid() bool {
	switch e {
	case ThemeDefault, ThemeLight, ThemeDark:
		return true
	}
	return false
}

func (e Theme) String() string {
	return string(e)
}

func (e *Theme) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Theme(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Theme", str)
	}
	return nil
}

func (e Theme) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
