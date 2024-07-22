package integrationapi

import (
	"encoding/json"
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/item"
	"github.com/reearth/reearth-cms/server/pkg/key"
	"github.com/reearth/reearth-cms/server/pkg/schema"
	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/reearth/reearth-cms/server/pkg/version"
	"github.com/reearth/reearthx/account/accountdomain"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestNewFeatureCollection(t *testing.T) {
	iid := id.NewItemID()
	sid := id.NewSchemaID()
	mid := id.NewModelID()
	uid := accountdomain.NewUserID()
	nid := id.NewIntegrationID()
	tid := id.NewThreadID()
	pid := id.NewProjectID()
	gst := schema.GeometryObjectSupportedTypeList{schema.GeometryObjectSupportedTypePoint, schema.GeometryObjectSupportedTypeLineString}
	sf1 := schema.NewField(schema.NewGeometryObject(gst).TypeProperty()).NewID().Key(key.Random()).MustBuild()
	sf2 := schema.NewField(schema.NewText(lo.ToPtr(10)).TypeProperty()).NewID().Key(key.Random()).MustBuild()
	s1 := schema.New().ID(sid).Fields([]*schema.Field{sf1}).Workspace(accountdomain.NewWorkspaceID()).TitleField(sf1.ID().Ref()).Project(pid).MustBuild()
	s2 := schema.New().ID(sid).Fields([]*schema.Field{sf2}).Workspace(accountdomain.NewWorkspaceID()).TitleField(sf2.ID().Ref()).Project(pid).MustBuild()
	i1 := item.New().
		ID(iid).
		Schema(sid).
		Project(pid).
		Fields([]*item.Field{item.NewField(sf1.ID(), value.TypeGeometryObject.Value("{\n\"type\": \"Point\",\n\t\"coordinates\": [102.1,0.5]\n}").AsMultiple(), nil)}).
		Model(mid).
		Thread(tid).
		User(uid).
		Integration(nid).
		MustBuild()
	v1 := version.New()
	vi1 := version.MustBeValue(v1, nil, version.NewRefs(version.Latest), util.Now(), i1)
	i2 := item.New().
		ID(iid).
		Schema(sid).
		Project(pid).
		Fields([]*item.Field{item.NewField(sf1.ID(), value.TypeText.Value("test").AsMultiple(), nil)}).
		Model(mid).
		Thread(tid).
		User(uid).
		Integration(nid).
		MustBuild()
	v2 := version.New()
	vi2 := version.MustBeValue(v2, nil, version.NewRefs(version.Latest), util.Now(), i2)

	// with geometry fields
	ver1 := item.VersionedList{vi1}
	point := []float32{102.1, 0.5}
	jsonBytes, err := json.Marshal(point)
	assert.Nil(t, err)
	c := Geometry_Coordinates{
		union: jsonBytes,
	}
	g := Geometry{
		Type:        lo.ToPtr(GeometryTypePoint),
		Coordinates: &c,
	}
	p := make(map[string]interface{})
	f := Feature{
		Type:       lo.ToPtr(FeatureTypeFeature),
		Geometry:   &g,
		Properties: &p,
		Id:         vi1.Value().ID().Ref(),
	}

	expected1 := &FeatureCollection{
		Type:     lo.ToPtr(FeatureCollectionTypeFeatureCollection),
		Features: &[]Feature{f},
	}

	fc1, err1 := FeatureCollectionFromItems(ver1, s1)
	assert.Nil(t, err1)
	assert.Equal(t, expected1, fc1)

	// no geometry fields
	ver2 := item.VersionedList{vi2}
	expectErr2 := noGeometryFieldError

	fc, err := FeatureCollectionFromItems(ver2, s2)
	assert.Equal(t, expectErr2, err)
	assert.Nil(t, fc)
}

func TestStringToGeometry(t *testing.T) {
	validGeoStringPoint := `
	{
		"type": "Point",
		"coordinates": [139.7112596, 35.6424892]
	}`
	geo, err := StringToGeometry(validGeoStringPoint)
	assert.NoError(t, err)
	assert.NotNil(t, geo)
	assert.Equal(t, GeometryTypePoint, *geo.Type)
	expected := []float32{139.7112596, 35.6424892}
	actual, err := geo.Coordinates.AsPoint()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)

	// Invalid Geometry string
	invalidGeometryString := "InvalidGeometry"
	geo, err = StringToGeometry(invalidGeometryString)
	assert.Error(t, err)
	assert.Nil(t, geo)
}
