package usecase

import (
	"errors"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/integration"
	"github.com/reearth/reearth-cms/server/pkg/project"
	"github.com/reearth/reearth-cms/server/pkg/user"
)

type Operator struct {
	user               *user.ID
	integration        *integration.ID
	ReadableWorkspaces user.WorkspaceIDList
	WritableWorkspaces user.WorkspaceIDList
	OwningWorkspaces   user.WorkspaceIDList
	ReadableProjects   project.IDList
	WritableProjects   project.IDList
	OwningProjects     project.IDList
}

type UserOperator struct {
	Operator
}

type IntegrationOperator struct {
	Operator
}

func (o *Operator) UserOperator() (*UserOperator, error) {
	if o == nil || o.user == nil {
		return nil, errors.New("invalid operator")
	}
	return &UserOperator{*o}, nil
}

func NewUserOperator(u user.ID, rw, ww, ow user.WorkspaceIDList, rp, wp, op project.IDList) *Operator {
	return &Operator{
		user:               &u,
		ReadableWorkspaces: rw,
		WritableWorkspaces: ww,
		OwningWorkspaces:   ow,
		ReadableProjects:   rp,
		WritableProjects:   wp,
		OwningProjects:     op,
	}
}

func NewIntegrationOperator(i integration.ID, rw, ww, ow user.WorkspaceIDList, rp, wp, op project.IDList) *Operator {
	return &Operator{
		integration:        &i,
		ReadableWorkspaces: rw,
		WritableWorkspaces: ww,
		OwningWorkspaces:   ow,
		ReadableProjects:   rp,
		WritableProjects:   wp,
		OwningProjects:     op,
	}
}

func (uo *UserOperator) User() user.ID {
	return *uo.user
}

func (o *Operator) IntegrationOperator() (*IntegrationOperator, error) {
	if o == nil || o.integration == nil {
		return nil, errors.New("invalid operator")
	}
	return &IntegrationOperator{*o}, nil
}

func (io *IntegrationOperator) Integration() integration.ID {
	return *io.integration
}

func (o *Operator) Workspaces(r user.Role) []id.WorkspaceID {
	if o == nil {
		return nil
	}
	if r == user.RoleReader {
		return o.ReadableWorkspaces
	}
	if r == user.RoleWriter {
		return o.WritableWorkspaces
	}
	if r == user.RoleOwner {
		return o.OwningWorkspaces
	}
	return nil
}

func (o *Operator) AllReadableWorkspaces() user.WorkspaceIDList {
	return append(o.ReadableWorkspaces, o.AllWritableWorkspaces()...)
}

func (o *Operator) AllWritableWorkspaces() user.WorkspaceIDList {
	return append(o.WritableWorkspaces, o.AllOwningWorkspaces()...)
}

func (o *Operator) AllOwningWorkspaces() user.WorkspaceIDList {
	return o.OwningWorkspaces
}

func (o *Operator) IsReadableWorkspace(workspace ...id.WorkspaceID) bool {
	return o.AllReadableWorkspaces().Intersect(workspace).Len() > 0
}

func (o *Operator) IsWritableWorkspace(workspace ...id.WorkspaceID) bool {
	return o.AllWritableWorkspaces().Intersect(workspace).Len() > 0
}

func (o *Operator) IsOwningWorkspace(workspace ...id.WorkspaceID) bool {
	return o.AllOwningWorkspaces().Intersect(workspace).Len() > 0
}

func (o *Operator) AddNewWorkspace(workspace id.WorkspaceID) {
	o.OwningWorkspaces = append(o.OwningWorkspaces, workspace)
}

func (o *Operator) Projects(r user.Role) project.IDList {
	if o == nil {
		return nil
	}
	if r == user.RoleReader {
		return o.ReadableProjects
	}
	if r == user.RoleWriter {
		return o.WritableProjects
	}
	if r == user.RoleOwner {
		return o.OwningProjects
	}
	return nil
}

func (o *Operator) AllReadableProjects() project.IDList {
	return append(o.ReadableProjects, o.AllWritableProjects()...)
}

func (o *Operator) AllWritableProjects() project.IDList {
	return append(o.WritableProjects, o.AllOwningProjects()...)
}

func (o *Operator) AllOwningProjects() project.IDList {
	return o.OwningProjects
}

func (o *Operator) IsReadableProject(projects ...project.ID) bool {
	return o.AllReadableProjects().Intersect(projects).Len() > 0
}

func (o *Operator) IsWritableProject(projects ...project.ID) bool {
	return o.AllWritableProjects().Intersect(projects).Len() > 0
}

func (o *Operator) IsOwningProject(projects ...project.ID) bool {
	return o.AllOwningProjects().Intersect(projects).Len() > 0
}

func (o *Operator) AddNewProject(p project.ID) {
	o.OwningProjects = append(o.OwningProjects, p)
}
