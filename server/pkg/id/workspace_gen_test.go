// Code generated by gen, DO NOT EDIT.

package id

import (
	"encoding/json"
	"testing"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestNewWorkspaceID(t *testing.T) {
	id := NewWorkspaceID()
	assert.NotNil(t, id)
	u, err := ulid.Parse(id.String())
	assert.NotNil(t, u)
	assert.Nil(t, err)
}

func TestWorkspaceIDFrom(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected struct {
			result WorkspaceID
			err    error
		}
	}{
		{
			name:  "Fail:Not valid string",
			input: "testMustFail",
			expected: struct {
				result WorkspaceID
				err    error
			}{
				result: WorkspaceID{},
				err:    ErrInvalidID,
			},
		},
		{
			name:  "Fail:Not valid string",
			input: "",
			expected: struct {
				result WorkspaceID
				err    error
			}{
				result: WorkspaceID{},
				err:    ErrInvalidID,
			},
		},
		{
			name:  "success:valid string",
			input: "01f2r7kg1fvvffp0gmexgy5hxy",
			expected: struct {
				result WorkspaceID
				err    error
			}{
				result: WorkspaceID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
				err:    nil,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result, err := WorkspaceIDFrom(tt.input)
			assert.Equal(t, tt.expected.result, result)
			if tt.expected.err != nil {
				assert.Equal(t, tt.expected.err, err)
			}
		})
	}
}

func TestMustWorkspaceID(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		shouldPanic bool
		expected    WorkspaceID
	}{
		{
			name:        "Fail:Not valid string",
			input:       "testMustFail",
			shouldPanic: true,
		},
		{
			name:        "Fail:Not valid string",
			input:       "",
			shouldPanic: true,
		},
		{
			name:        "success:valid string",
			input:       "01f2r7kg1fvvffp0gmexgy5hxy",
			shouldPanic: false,
			expected:    WorkspaceID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.shouldPanic {
				assert.Panics(t, func() { MustBeID(tt.input) })
				return
			}
			result := MustWorkspaceID(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestWorkspaceIDFromRef(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *WorkspaceID
	}{
		{
			name:     "Fail:Not valid string",
			input:    "testMustFail",
			expected: nil,
		},
		{
			name:     "Fail:Not valid string",
			input:    "",
			expected: nil,
		},
		{
			name:     "success:valid string",
			input:    "01f2r7kg1fvvffp0gmexgy5hxy",
			expected: &WorkspaceID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := WorkspaceIDFromRef(&tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestWorkspaceIDFromRefID(t *testing.T) {
	id := New()
	id2 := WorkspaceIDFromRefID(&id)
	assert.Equal(t, id.id, id2.id)
	assert.Nil(t, WorkspaceIDFromRefID(nil))
	assert.Nil(t, WorkspaceIDFromRefID(&ID{}))
}

func TestWorkspaceID_ID(t *testing.T) {
	id := New()
	id2 := WorkspaceIDFromRefID(&id)
	assert.Equal(t, id, id2.ID())
}

func TestWorkspaceID_String(t *testing.T) {
	id := New()
	id2 := WorkspaceIDFromRefID(&id)
	assert.Equal(t, id.String(), id2.String())
	assert.Equal(t, "", WorkspaceID{}.String())
}

func TestWorkspaceID_RefString(t *testing.T) {
	id := NewWorkspaceID()
	assert.Equal(t, id.String(), *id.RefString())
	assert.Nil(t, WorkspaceID{}.RefString())
}

func TestWorkspaceID_GoString(t *testing.T) {
	id := New()
	id2 := WorkspaceIDFromRefID(&id)
	assert.Equal(t, "WorkspaceID("+id.String()+")", id2.GoString())
	assert.Equal(t, "WorkspaceID()", WorkspaceID{}.GoString())
}

func TestWorkspaceID_Ref(t *testing.T) {
	id := NewWorkspaceID()
	assert.Equal(t, WorkspaceID(id), *id.Ref())
	assert.Nil(t, (&WorkspaceID{}).Ref())
}

func TestWorkspaceID_Contains(t *testing.T) {
	id := NewWorkspaceID()
	id2 := NewWorkspaceID()
	assert.True(t, id.Contains([]WorkspaceID{id, id2}))
	assert.False(t, WorkspaceID{}.Contains([]WorkspaceID{id, id2, {}}))
	assert.False(t, id.Contains([]WorkspaceID{id2}))
}

func TestWorkspaceID_CopyRef(t *testing.T) {
	id := NewWorkspaceID().Ref()
	id2 := id.CopyRef()
	assert.Equal(t, id, id2)
	assert.NotSame(t, id, id2)
	assert.Nil(t, (*WorkspaceID)(nil).CopyRef())
}

func TestWorkspaceID_IDRef(t *testing.T) {
	id := New()
	id2 := WorkspaceIDFromRefID(&id)
	assert.Equal(t, &id, id2.IDRef())
	assert.Nil(t, (&WorkspaceID{}).IDRef())
	assert.Nil(t, (*WorkspaceID)(nil).IDRef())
}

func TestWorkspaceID_StringRef(t *testing.T) {
	id := NewWorkspaceID()
	assert.Equal(t, id.String(), *id.StringRef())
	assert.Nil(t, (&WorkspaceID{}).StringRef())
	assert.Nil(t, (*WorkspaceID)(nil).StringRef())
}

func TestWorkspaceID_MarhsalJSON(t *testing.T) {
	id := NewWorkspaceID()
	res, err := id.MarhsalJSON()
	assert.Nil(t, err)
	exp, _ := json.Marshal(id.String())
	assert.Equal(t, exp, res)

	res, err = (&WorkspaceID{}).MarhsalJSON()
	assert.Nil(t, err)
	assert.Nil(t, res)

	res, err = (*WorkspaceID)(nil).MarhsalJSON()
	assert.Nil(t, err)
	assert.Nil(t, res)
}

func TestWorkspaceID_UnmarhsalJSON(t *testing.T) {
	jsonString := "\"01f3zhkysvcxsnzepyyqtq21fb\""
	id := MustWorkspaceID("01f3zhkysvcxsnzepyyqtq21fb")
	id2 := &WorkspaceID{}
	err := id2.UnmarhsalJSON([]byte(jsonString))
	assert.Nil(t, err)
	assert.Equal(t, id, *id2)
}

func TestWorkspaceID_MarshalText(t *testing.T) {
	id := New()
	res, err := WorkspaceIDFromRefID(&id).MarshalText()
	assert.Nil(t, err)
	assert.Equal(t, []byte(id.String()), res)

	res, err = (&WorkspaceID{}).MarshalText()
	assert.Nil(t, err)
	assert.Nil(t, res)

	res, err = (*WorkspaceID)(nil).MarshalText()
	assert.Nil(t, err)
	assert.Nil(t, res)
}

func TestWorkspaceID_UnmarshalText(t *testing.T) {
	text := []byte("01f3zhcaq35403zdjnd6dcm0t2")
	id2 := &WorkspaceID{}
	err := id2.UnmarshalText(text)
	assert.Nil(t, err)
	assert.Equal(t, "01f3zhcaq35403zdjnd6dcm0t2", id2.String())
}

func TestWorkspaceID_IsNil(t *testing.T) {
	assert.True(t, WorkspaceID{}.IsNil())
	assert.False(t, NewWorkspaceID().IsNil())
}

func TestWorkspaceID_IsNilRef(t *testing.T) {
	assert.True(t, WorkspaceID{}.Ref().IsNilRef())
	assert.True(t, (*WorkspaceID)(nil).IsNilRef())
	assert.False(t, NewWorkspaceID().Ref().IsNilRef())
}

func TestWorkspaceIDsToStrings(t *testing.T) {
	tests := []struct {
		name     string
		input    []WorkspaceID
		expected []string
	}{
		{
			name:     "Empty slice",
			input:    make([]WorkspaceID, 0),
			expected: make([]string, 0),
		},
		{
			name:     "1 element",
			input:    []WorkspaceID{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
		},
		{
			name: "multiple elements",
			input: []WorkspaceID{
				MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: []string{
				"01f3zhcaq35403zdjnd6dcm0t1",
				"01f3zhcaq35403zdjnd6dcm0t2",
				"01f3zhcaq35403zdjnd6dcm0t3",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, WorkspaceIDsToStrings(tt.input))
		})
	}
}

func TestWorkspaceIDsFrom(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected struct {
			res []WorkspaceID
			err error
		}
	}{
		{
			name:  "Empty slice",
			input: make([]string, 0),
			expected: struct {
				res []WorkspaceID
				err error
			}{
				res: make([]WorkspaceID, 0),
				err: nil,
			},
		},
		{
			name:  "1 element",
			input: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
			expected: struct {
				res []WorkspaceID
				err error
			}{
				res: []WorkspaceID{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2")},
				err: nil,
			},
		},
		{
			name: "multiple elements",
			input: []string{
				"01f3zhcaq35403zdjnd6dcm0t1",
				"01f3zhcaq35403zdjnd6dcm0t2",
				"01f3zhcaq35403zdjnd6dcm0t3",
			},
			expected: struct {
				res []WorkspaceID
				err error
			}{
				res: []WorkspaceID{
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
				err: nil,
			},
		},
		{
			name: "error",
			input: []string{
				"01f3zhcaq35403zdjnd6dcm0t1",
				"x",
				"01f3zhcaq35403zdjnd6dcm0t3",
			},
			expected: struct {
				res []WorkspaceID
				err error
			}{
				res: nil,
				err: ErrInvalidID,
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := WorkspaceIDsFrom(tc.input)
			if tc.expected.err != nil {
				assert.Equal(t, tc.expected.err, err)
				assert.Nil(t, res)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.expected.res, res)
			}
		})
	}
}

func TestWorkspaceIDsFromID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    []ID
		expected []WorkspaceID
	}{
		{
			name:     "Empty slice",
			input:    make([]ID, 0),
			expected: make([]WorkspaceID, 0),
		},
		{
			name:     "1 element",
			input:    []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []WorkspaceID{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []ID{
				MustBeID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: []WorkspaceID{
				MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := WorkspaceIDsFromID(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestWorkspaceIDsFromIDRef(t *testing.T) {
	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")

	tests := []struct {
		name     string
		input    []*ID
		expected []WorkspaceID
	}{
		{
			name:     "Empty slice",
			input:    make([]*ID, 0),
			expected: make([]WorkspaceID, 0),
		},
		{
			name:     "1 element",
			input:    []*ID{&id1},
			expected: []WorkspaceID{MustWorkspaceID(id1.String())},
		},
		{
			name:  "multiple elements",
			input: []*ID{&id1, &id2, &id3},
			expected: []WorkspaceID{
				MustWorkspaceID(id1.String()),
				MustWorkspaceID(id2.String()),
				MustWorkspaceID(id3.String()),
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := WorkspaceIDsFromIDRef(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestWorkspaceIDsToID(t *testing.T) {
	tests := []struct {
		name     string
		input    []WorkspaceID
		expected []ID
	}{
		{
			name:     "Empty slice",
			input:    make([]WorkspaceID, 0),
			expected: make([]ID, 0),
		},
		{
			name:     "1 element",
			input:    []WorkspaceID{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []WorkspaceID{
				MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: []ID{
				MustBeID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := WorkspaceIDsToID(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestWorkspaceIDsToIDRef(t *testing.T) {
	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	id21 := MustWorkspaceID(id1.String())
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	id22 := MustWorkspaceID(id2.String())
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")
	id23 := MustWorkspaceID(id3.String())

	tests := []struct {
		name     string
		input    []*WorkspaceID
		expected []*ID
	}{
		{
			name:     "Empty slice",
			input:    make([]*WorkspaceID, 0),
			expected: make([]*ID, 0),
		},
		{
			name:     "1 element",
			input:    []*WorkspaceID{&id21},
			expected: []*ID{&id1},
		},
		{
			name:     "multiple elements",
			input:    []*WorkspaceID{&id21, &id22, &id23},
			expected: []*ID{&id1, &id2, &id3},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := WorkspaceIDsToIDRef(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestNewWorkspaceIDSet(t *testing.T) {
	WorkspaceIdSet := NewWorkspaceIDSet()
	assert.NotNil(t, WorkspaceIdSet)
	assert.Empty(t, WorkspaceIdSet.m)
	assert.Empty(t, WorkspaceIdSet.s)
}

func TestWorkspaceIDSet_Add(t *testing.T) {
	tests := []struct {
		name     string
		input    []WorkspaceID
		expected *WorkspaceIDSet
	}{
		{
			name:  "Empty slice",
			input: make([]WorkspaceID, 0),
			expected: &WorkspaceIDSet{
				m: map[WorkspaceID]struct{}{},
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: []WorkspaceID{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: &WorkspaceIDSet{
				m: map[WorkspaceID]struct{}{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []WorkspaceID{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: []WorkspaceID{
				MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &WorkspaceIDSet{
				m: map[WorkspaceID]struct{}{
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []WorkspaceID{
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
		{
			name: "multiple elements with duplication",
			input: []WorkspaceID{
				MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &WorkspaceIDSet{
				m: map[WorkspaceID]struct{}{
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []WorkspaceID{
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			set := NewWorkspaceIDSet()
			set.Add(tc.input...)
			assert.Equal(t, tc.expected, set)
		})
	}
}

func TestWorkspaceIDSet_AddRef(t *testing.T) {
	tests := []struct {
		name     string
		input    *WorkspaceID
		expected *WorkspaceIDSet
	}{
		{
			name:  "Empty slice",
			input: nil,
			expected: &WorkspaceIDSet{
				m: nil,
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1").Ref(),
			expected: &WorkspaceIDSet{
				m: map[WorkspaceID]struct{}{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []WorkspaceID{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			set := NewWorkspaceIDSet()
			set.AddRef(tc.input)
			assert.Equal(t, tc.expected, set)
		})
	}
}

func TestWorkspaceIDSet_Has(t *testing.T) {
	tests := []struct {
		name     string
		target   *WorkspaceIDSet
		input    WorkspaceID
		expected bool
	}{
		{
			name:     "Empty Set",
			target:   &WorkspaceIDSet{},
			input:    MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"),
			expected: false,
		},
		{
			name: "Set Contains the element",
			target: &WorkspaceIDSet{
				m: map[WorkspaceID]struct{}{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []WorkspaceID{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			input:    MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"),
			expected: true,
		},
		{
			name: "Set does not Contains the element",
			target: &WorkspaceIDSet{
				m: map[WorkspaceID]struct{}{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []WorkspaceID{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			input:    MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2"),
			expected: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, tc.target.Has(tc.input))
		})
	}
}

func TestWorkspaceIDSet_Clear(t *testing.T) {
	tests := []struct {
		name     string
		input    *WorkspaceIDSet
		expected *WorkspaceIDSet
	}{
		{
			name:     "Empty set",
			input:    &WorkspaceIDSet{},
			expected: &WorkspaceIDSet{},
		},
		{
			name:     "Nil set",
			input:    nil,
			expected: nil,
		},
		{
			name: "Contains the element",
			input: &WorkspaceIDSet{
				m: map[WorkspaceID]struct{}{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []WorkspaceID{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: &WorkspaceIDSet{
				m: nil,
				s: nil,
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.input.Clear()
			assert.Equal(t, tc.expected, tc.input)
		})
	}
}

func TestWorkspaceIDSet_All(t *testing.T) {
	tests := []struct {
		name     string
		input    *WorkspaceIDSet
		expected []WorkspaceID
	}{
		{
			name: "Empty",
			input: &WorkspaceIDSet{
				m: map[WorkspaceID]struct{}{},
				s: nil,
			},
			expected: make([]WorkspaceID, 0),
		},
		{
			name:     "Nil",
			input:    nil,
			expected: nil,
		},
		{
			name: "1 element",
			input: &WorkspaceIDSet{
				m: map[WorkspaceID]struct{}{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []WorkspaceID{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: []WorkspaceID{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1")},
		},
		{
			name: "multiple elements",
			input: &WorkspaceIDSet{
				m: map[WorkspaceID]struct{}{
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []WorkspaceID{
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: []WorkspaceID{
				MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, tc.input.All())
		})
	}
}

func TestWorkspaceIDSet_Clone(t *testing.T) {
	tests := []struct {
		name     string
		input    *WorkspaceIDSet
		expected *WorkspaceIDSet
	}{
		{
			name:     "nil set",
			input:    nil,
			expected: NewWorkspaceIDSet(),
		},
		{
			name:     "Empty set",
			input:    NewWorkspaceIDSet(),
			expected: NewWorkspaceIDSet(),
		},
		{
			name: "1 element",
			input: &WorkspaceIDSet{
				m: map[WorkspaceID]struct{}{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []WorkspaceID{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: &WorkspaceIDSet{
				m: map[WorkspaceID]struct{}{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []WorkspaceID{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: &WorkspaceIDSet{
				m: map[WorkspaceID]struct{}{
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []WorkspaceID{
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: &WorkspaceIDSet{
				m: map[WorkspaceID]struct{}{
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []WorkspaceID{
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			clone := tc.input.Clone()
			assert.Equal(t, tc.expected, clone)
			assert.NotSame(t, tc.input, clone)
		})
	}
}

func TestWorkspaceIDSet_Merge(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			a *WorkspaceIDSet
			b *WorkspaceIDSet
		}
		expected *WorkspaceIDSet
	}{
		{
			name: "Nil Set",
			input: struct {
				a *WorkspaceIDSet
				b *WorkspaceIDSet
			}{
				a: &WorkspaceIDSet{
					m: map[WorkspaceID]struct{}{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
					s: []WorkspaceID{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: nil,
			},
			expected: &WorkspaceIDSet{
				m: map[WorkspaceID]struct{}{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []WorkspaceID{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "Empty Set",
			input: struct {
				a *WorkspaceIDSet
				b *WorkspaceIDSet
			}{
				a: &WorkspaceIDSet{},
				b: &WorkspaceIDSet{},
			},
			expected: &WorkspaceIDSet{},
		},
		{
			name: "1 Empty Set",
			input: struct {
				a *WorkspaceIDSet
				b *WorkspaceIDSet
			}{
				a: &WorkspaceIDSet{
					m: map[WorkspaceID]struct{}{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
					s: []WorkspaceID{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &WorkspaceIDSet{},
			},
			expected: &WorkspaceIDSet{
				m: map[WorkspaceID]struct{}{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []WorkspaceID{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "2 non Empty Set",
			input: struct {
				a *WorkspaceIDSet
				b *WorkspaceIDSet
			}{
				a: &WorkspaceIDSet{
					m: map[WorkspaceID]struct{}{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
					s: []WorkspaceID{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &WorkspaceIDSet{
					m: map[WorkspaceID]struct{}{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2"): {}},
					s: []WorkspaceID{MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2")},
				},
			},
			expected: &WorkspaceIDSet{
				m: map[WorkspaceID]struct{}{
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2"): {},
				},
				s: []WorkspaceID{
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustWorkspaceID("01f3zhcaq35403zdjnd6dcm0t2"),
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, tc.input.a.Merge(tc.input.b))
		})
	}
}
