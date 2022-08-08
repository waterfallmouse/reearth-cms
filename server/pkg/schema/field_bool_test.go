package schema

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestFieldBoolFrom(t *testing.T) {
	tests := []struct {
		name string
		arg  *bool
		want *FieldBool
	}{
		{
			name: "success default nil",
			arg:  nil,
			want: &FieldBool{defaultValue: nil},
		},
		{
			name: "success default true",
			arg:  lo.ToPtr(true),
			want: &FieldBool{defaultValue: lo.ToPtr(true)},
		},
		{
			name: "success default false",
			arg:  lo.ToPtr(false),
			want: &FieldBool{defaultValue: lo.ToPtr(false)},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, FieldBoolFrom(tc.arg))
		})
	}
}

func TestFieldBool_TypeProperty(t *testing.T) {
	tests := []struct {
		name string
		f    *FieldBool
		want *TypeProperty
	}{
		{
			name: "nil",
			f:    nil,
			want: &TypeProperty{},
		},
		{
			name: "success default nil",
			f:    &FieldBool{defaultValue: nil},
			want: &TypeProperty{bool: &FieldBool{defaultValue: nil}},
		},
		{
			name: "success default true",
			f:    &FieldBool{defaultValue: lo.ToPtr(true)},
			want: &TypeProperty{bool: &FieldBool{defaultValue: lo.ToPtr(true)}},
		},
		{
			name: "success default false",
			f:    &FieldBool{defaultValue: lo.ToPtr(false)},
			want: &TypeProperty{bool: &FieldBool{defaultValue: lo.ToPtr(false)}},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.want, tc.f.TypeProperty())
		})
	}
}

/*func TestnewFieldBool(t *testing.T) {
	tests := []struct {
		name string
		want *FieldBool
	}{
		{
			name: "new",
			want: &FieldBool{
				defaultValue: nil,
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.want, newFieldBool())
		})
	}
}*/

func TestFieldBool_DefaultValue(t *testing.T) {
	type fields struct {
		defaultValue *bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *bool
	}{
		{
			name: "test",
			fields: fields{
				defaultValue: lo.ToPtr(true),
			},
			want: lo.ToPtr(true),
		},
		{
			name: "test",
			fields: fields{
				defaultValue: lo.ToPtr(false),
			},
			want: lo.ToPtr(false),
		},
		{
			name: "test",
			fields: fields{
				defaultValue: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			f := &FieldBool{
				defaultValue: tt.fields.defaultValue,
			}
			assert.Equal(t, tt.want, f.DefaultValue())
		})
	}
}
