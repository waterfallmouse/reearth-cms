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
	gst := schema.GeometryObjectSupportedTypeList{schema.GeometryObjectSupportedTypePoint, schema.GeometryObjectSupportedTypeLineString, schema.GeometryObjectSupportedTypePolygon}
	gest := schema.GeometryEditorSupportedTypeList{schema.GeometryEditorSupportedTypePoint, schema.GeometryEditorSupportedTypeLineString, schema.GeometryEditorSupportedTypePolygon}
	sf1 := schema.NewField(schema.NewGeometryObject(gst).TypeProperty()).NewID().Name("LineString").Key(key.Random()).MustBuild()
	sf2 := schema.NewField(schema.NewText(lo.ToPtr(10)).TypeProperty()).NewID().Name("Name").Key(key.Random()).Multiple(true).MustBuild()
	sf3 := schema.NewField(schema.NewGeometryEditor(gest).TypeProperty()).NewID().Name("Polygon").Key(key.Random()).MustBuild()
	in4, _ := schema.NewInteger(lo.ToPtr(int64(1)), lo.ToPtr(int64(100)))
	tp4 := in4.TypeProperty()
	sf4 := schema.NewField(tp4).NewID().Name("Age").Key(key.Random()).MustBuild()
	sf5 := schema.NewField(schema.NewBool().TypeProperty()).NewID().Name("IsMarried").Key(key.Random()).MustBuild()
	fi1 := item.NewField(sf1.ID(), value.TypeGeometryObject.Value("{\"coordinates\":[[139.65439725962517,36.34793305387103],[139.61688622815393,35.910803456352724]],\"type\":\"LineString\"}").AsMultiple(), nil)
	fi2 := item.NewField(sf2.ID(), value.MultipleFrom(value.TypeText, []*value.Value{value.TypeText.Value("a"), value.TypeText.Value("b"), value.TypeText.Value("c")}), nil)
	fi3 := item.NewField(sf3.ID(), value.TypeGeometryEditor.Value("{\"coordinates\": [[[138.90306434425662,36.11737907906834],[138.90306434425662,36.33622175736386],[138.67187898370287,36.33622175736386],[138.67187898370287,36.11737907906834],[138.90306434425662,36.11737907906834]]],\"type\": \"Polygon\"}").AsMultiple(), nil)
	fi4 := item.NewField(sf4.ID(), value.TypeInteger.Value(30).AsMultiple(), nil)
	fi5 := item.NewField(sf5.ID(), value.TypeBool.Value(true).AsMultiple(), nil)
	s1 := schema.New().ID(sid).Fields([]*schema.Field{sf1, sf2, sf3, sf4, sf5}).Workspace(accountdomain.NewWorkspaceID()).TitleField(sf1.ID().Ref()).Project(pid).MustBuild()
	s2 := schema.New().ID(sid).Fields([]*schema.Field{sf2}).Workspace(accountdomain.NewWorkspaceID()).TitleField(sf2.ID().Ref()).Project(pid).MustBuild()
	i1 := item.New().
		ID(iid).
		Schema(sid).
		Project(pid).
		Fields([]*item.Field{fi1, fi2, fi3, fi4, fi5}).
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
	lineString := [][]float64{
		{139.65439725962517, 36.34793305387103},
		{139.61688622815393, 35.910803456352724},
	}
	jsonBytes, err := json.Marshal(lineString)
	assert.Nil(t, err)
	c := Geometry_Coordinates{
		union: jsonBytes,
	}
	g := Geometry{
		Type:        lo.ToPtr(GeometryTypeLineString),
		Coordinates: &c,
	}
	p := make(map[string]interface{})
	p["Name"] = []string{"a", "b", "c"}
	p["Age"] = "30"
	p["IsMarried"] = "true"

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

func TestToGeoJSONProp(t *testing.T) {
	sf1 := schema.NewField(schema.NewText(lo.ToPtr(10)).TypeProperty()).NewID().Key(key.Random()).MustBuild()
	if1 := item.NewField(sf1.ID(), value.TypeText.Value("Nour").AsMultiple(), nil)
	s1, ok1 := toGeoJSONProp(if1)
	assert.Equal(t, "Nour", s1)
	assert.True(t, ok1)
	
	var if2 *item.Field
	s2, ok2 := toGeoJSONProp(if2)
	assert.Empty(t, s2)
	assert.False(t, ok2)

	sf3 := schema.NewField(schema.NewText(lo.ToPtr(10)).TypeProperty()).NewID().Key(key.Random()).MustBuild()
	if3 := item.NewField(sf3.ID(), value.MultipleFrom(value.TypeText, []*value.Value{value.TypeText.Value("a"), value.TypeText.Value("b"), value.TypeText.Value("c")}), nil)
	s3, ok3 := toGeoJSONProp(if3)
	assert.Equal(t, []string{"a","b", "c"}, s3)
	assert.True(t, ok3)
}