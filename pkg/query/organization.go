package query

import (
	"context"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
)

type Organization struct {
	ctx context.Context
	*organizationmanager.Organization
}

type Group struct {
	ctx context.Context
	*organizationmanager.Group
}

type Member struct {
	ctx context.Context
	*organizationmanager.GroupMember
}

func (o Organization) Cloud(name string) *Cloud {
	c, err := GetSdk(o.ctx).ResourceManager().Cloud().List(o.ctx, &resourcemanager.ListCloudsRequest{
		OrganizationId: GetOrganizationID(o.ctx),
		Filter:         `name = '` + name + `'`,
	})

	if err != nil {
		panic(err)
	}

	if len(c.Clouds) == 0 {
		return nil
	}

	return NewCloud(o.ctx, c.Clouds[0])
}

func (o Organization) Clouds() []Cloud {
	c, err := GetSdk(o.ctx).ResourceManager().Cloud().List(o.ctx, &resourcemanager.ListCloudsRequest{
		OrganizationId: GetOrganizationID(o.ctx),
	})

	if err != nil {
		panic(err)
	}

	clouds := make([]Cloud, 0, len(c.Clouds))
	for _, cloud := range c.Clouds {
		clouds = append(clouds, *NewCloud(o.ctx, cloud))
	}
	return clouds
}

func (o Organization) Group(name string) Group {
	g, err := GetSdk(o.ctx).OrganizationManager().Group().List(o.ctx, &organizationmanager.ListGroupsRequest{
		OrganizationId: GetOrganizationID(o.ctx),
		Filter:         `name = '` + name + `'`,
	})

	if err != nil {
		panic(err)
	}

	if len(g.Groups) == 0 {
		return Group{}
	}

	return Group{
		ctx:   WithGroupID(o.ctx, g.Groups[0].Id),
		Group: g.Groups[0],
	}
}

func (o Organization) Groups() []Group {
	g, err := GetSdk(o.ctx).OrganizationManager().Group().List(o.ctx, &organizationmanager.ListGroupsRequest{
		OrganizationId: GetOrganizationID(o.ctx),
	})

	if err != nil {
		panic(err)
	}

	groups := make([]Group, 0, len(g.Groups))
	for _, group := range g.Groups {
		groups = append(groups, Group{
			ctx:   WithGroupID(o.ctx, group.Id),
			Group: group,
		})
	}
	return groups
}

func (g Group) Members() {
	m, err := GetSdk(g.ctx).OrganizationManager().Group().ListMembers(g.ctx, &organizationmanager.ListGroupMembersRequest{
		GroupId: GetGroupID(g.ctx),
	})

	if err != nil {
		panic(err)
	}

	members := make([]Member, 0, len(m.Members))
	for _, member := range m.Members {
		members = append(members, Member{
			ctx:         g.ctx,
			GroupMember: member,
		})
	}
}
