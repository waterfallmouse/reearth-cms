package group

import (
	"github.com/reearth/reearthx/util"
	"golang.org/x/exp/slices"
)

type List []*Group

func (l List) SortByID() List {
	m := slices.Clone(l)
	slices.SortFunc(m, func(a, b *Group) bool {
		return a.ID().Compare(b.ID()) < 0
	})
	return m
}

func (l List) Clone() List {
	return util.Map(l, func(g *Group) *Group { return g.Clone() })
}
