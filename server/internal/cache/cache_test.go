package cache

import (
	"context"
	"testing"

	"github.com/llm-operator/rbac-manager/server/internal/config"
	uv1 "github.com/llm-operator/user-manager/api/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestCache(t *testing.T) {
	l := &fakeUserInfoLister{
		apikeys: &uv1.ListAPIKeysResponse{
			Data: []*uv1.APIKey{
				{
					Id:           "id0",
					Secret:       "s0",
					User:         &uv1.User{Id: "u0"},
					Organization: &uv1.Organization{Id: "o0"},
					Project:      &uv1.Project{Id: "p0"},
				},
				{
					Id:           "id1",
					Secret:       "s1",
					User:         &uv1.User{Id: "u1"},
					Organization: &uv1.Organization{Id: "o1"},
					Project:      &uv1.Project{Id: "p1"},
				},
			},
		},
		orgs: &uv1.ListOrganizationsResponse{
			Organizations: []*uv1.Organization{
				{
					Id: "o0",
				},
				{
					Id: "o1",
				},
			},
		},
		orgusers: &uv1.ListOrganizationUsersResponse{
			Users: []*uv1.OrganizationUser{
				{
					UserId:         "u0",
					OrganizationId: "o0",
					Role:           uv1.OrganizationRole_ORGANIZATION_ROLE_OWNER,
				},
				{
					UserId:         "u0",
					OrganizationId: "o1",
					Role:           uv1.OrganizationRole_ORGANIZATION_ROLE_READER,
				},
			},
		},
		projects: &uv1.ListProjectsResponse{
			Projects: []*uv1.Project{
				{
					Id:                  "p0",
					OrganizationId:      "o0",
					KubernetesNamespace: "ns0",
				},
				{
					Id:                  "p1",
					OrganizationId:      "o1",
					KubernetesNamespace: "ns1",
				},
			},
		},
		projectusers: &uv1.ListProjectUsersResponse{
			Users: []*uv1.ProjectUser{
				{
					UserId:    "u0",
					ProjectId: "p0",
					Role:      uv1.ProjectRole_PROJECT_ROLE_OWNER,
				},
				{
					UserId:    "u0",
					ProjectId: "p1",
					Role:      uv1.ProjectRole_PROJECT_ROLE_MEMBER,
				},
			},
		},
	}
	c := NewStore(
		&fakeUserInfoListerFactory{l},
		&config.DebugConfig{
			APIKeyRole: "role",
		},
	)

	err := c.updateCache(context.Background())
	assert.NoError(t, err)

	wantKeys := map[string]*K{
		"s0": {
			Role:           "role",
			UserID:         "u0",
			OrganizationID: "o0",
		},
		"s1": {
			Role:           "role",
			UserID:         "u1",
			OrganizationID: "o1",
		},
		"s2": nil,
	}

	for k, v := range wantKeys {
		got, ok := c.GetAPIKeyBySecret(k)
		if v == nil {
			assert.False(t, ok)
			continue
		}

		assert.True(t, ok)
		assert.Equal(t, v.Role, got.Role)
	}

	wantOrgs := map[string]*O{
		"o0": {
			ID: "o0",
		},
		"o1": {
			ID: "o1",
		},
	}
	for id, want := range wantOrgs {
		got, ok := c.GetOrganizationByID(id)
		assert.True(t, ok)
		assert.Equal(t, want, got)
	}

	userorgs := c.GetOrganizationsByUserID("u0")
	assert.Len(t, userorgs, 2)
	userorgsByOrg := map[string]*OU{}
	for _, uo := range userorgs {
		userorgsByOrg[uo.OrganizationID] = &uo
	}
	wantOUs := map[string]*OU{
		"o0": {
			Role:           uv1.OrganizationRole_ORGANIZATION_ROLE_OWNER,
			OrganizationID: "o0",
		},
		"o1": {
			Role:           uv1.OrganizationRole_ORGANIZATION_ROLE_READER,
			OrganizationID: "o1",
		},
	}
	for orgID, want := range wantOUs {
		got, ok := userorgsByOrg[orgID]
		assert.True(t, ok)
		assert.Equal(t, want, got)
	}

	userorgs = c.GetOrganizationsByUserID("u1")
	assert.Empty(t, userorgs)

	wantProjects := map[string]*P{
		"p0": {
			ID:                  "p0",
			OrganizationID:      "o0",
			KubernetesNamespace: "ns0",
		},
		"p1": {
			ID:                  "p1",
			OrganizationID:      "o1",
			KubernetesNamespace: "ns1",
		},
	}
	for id, want := range wantProjects {
		got, ok := c.GetProjectByID(id)
		assert.True(t, ok)
		assert.Equal(t, want, got)

		gots := c.GetProjectsByOrganizationID(want.OrganizationID)
		assert.Len(t, gots, 1)
		assert.Equal(t, *want, gots[0])
	}

	userprojects := c.GetProjectsByUserID("u0")
	assert.Len(t, userprojects, 2)
	userprojectsByProject := map[string]*PU{}
	for _, up := range userprojects {
		userprojectsByProject[up.ProjectID] = &up
	}
	wantPUs := map[string]*PU{
		"p0": {
			Role:      uv1.ProjectRole_PROJECT_ROLE_OWNER,
			ProjectID: "p0",
		},
		"p1": {
			Role:      uv1.ProjectRole_PROJECT_ROLE_MEMBER,
			ProjectID: "p1",
		},
	}
	for projectID, want := range wantPUs {
		got, ok := userprojectsByProject[projectID]
		assert.True(t, ok)
		assert.Equal(t, want, got)
	}
}

type fakeUserInfoListerFactory struct {
	l *fakeUserInfoLister
}

func (f *fakeUserInfoListerFactory) Create() (UserInfoLister, error) {
	return f.l, nil
}

type fakeUserInfoLister struct {
	apikeys      *uv1.ListAPIKeysResponse
	orgs         *uv1.ListOrganizationsResponse
	orgusers     *uv1.ListOrganizationUsersResponse
	projects     *uv1.ListProjectsResponse
	projectusers *uv1.ListProjectUsersResponse
}

func (l *fakeUserInfoLister) ListAPIKeys(ctx context.Context, in *uv1.ListAPIKeysRequest, opts ...grpc.CallOption) (*uv1.ListAPIKeysResponse, error) {
	return l.apikeys, nil
}

func (l *fakeUserInfoLister) ListOrganizations(ctx context.Context, in *uv1.ListOrganizationsRequest, opts ...grpc.CallOption) (*uv1.ListOrganizationsResponse, error) {
	return l.orgs, nil
}

func (l *fakeUserInfoLister) ListOrganizationUsers(ctx context.Context, in *uv1.ListOrganizationUsersRequest, opts ...grpc.CallOption) (*uv1.ListOrganizationUsersResponse, error) {
	return l.orgusers, nil
}

func (l *fakeUserInfoLister) ListProjects(ctx context.Context, in *uv1.ListProjectsRequest, opts ...grpc.CallOption) (*uv1.ListProjectsResponse, error) {
	return l.projects, nil
}
func (l *fakeUserInfoLister) ListProjectUsers(ctx context.Context, in *uv1.ListProjectUsersRequest, opts ...grpc.CallOption) (*uv1.ListProjectUsersResponse, error) {
	return l.projectusers, nil
}
