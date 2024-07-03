package query

import (
	"context"
	"time"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
)

type Network struct {
	ctx       context.Context
	CreatedAt time.Time `json:"created_at"`
	*vpc.Network
}

func NewNetwork(ctx context.Context, network *vpc.Network) *Network {
	return &Network{ctx: WithNetworkID(ctx, network.Id), Network: network, CreatedAt: network.CreatedAt.AsTime()}
}

type Subnet struct {
	ctx       context.Context
	CreatedAt time.Time `json:"created_at"`
	*vpc.Subnet
}

func NewSubnet(ctx context.Context, subnet *vpc.Subnet) *Subnet {
	return &Subnet{ctx: WithSubnetID(ctx, subnet.Id), Subnet: subnet, CreatedAt: subnet.CreatedAt.AsTime()}
}

func (n Network) Subnet(name string) *Subnet {
	s, err := GetSdk(n.ctx).VPC().Subnet().List(n.ctx, &vpc.ListSubnetsRequest{
		FolderId: GetFolderID(n.ctx),
		Filter:   `name = '` + name + `'`,
	})

	if err != nil {
		panic(err)
	}

	if len(s.Subnets) == 0 {
		return nil
	}

	return NewSubnet(n.ctx, s.Subnets[0])
}

func (n Network) Subnets() []Subnet {
	s, err := GetSdk(n.ctx).VPC().Subnet().List(n.ctx, &vpc.ListSubnetsRequest{
		FolderId: GetFolderID(n.ctx),
	})

	if err != nil {
		panic(err)
	}

	subnets := make([]Subnet, 0, len(s.Subnets))
	for _, subnet := range s.Subnets {
		if subnet.NetworkId != n.Id {
			continue
		}
		subnets = append(subnets, *NewSubnet(n.ctx, subnet))
	}

	return subnets
}
