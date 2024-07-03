package query

import (
	"context"
	"time"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
)

type Folder struct {
	ctx       context.Context
	CreatedAt time.Time `json:"created_at"`
	*resourcemanager.Folder
}

func NewFolder(ctx context.Context, folder *resourcemanager.Folder) *Folder {
	return &Folder{ctx: WithFolderID(ctx, folder.Id), Folder: folder, CreatedAt: folder.CreatedAt.AsTime()}
}

func (f Folder) Network(name string) *Network {
	n, err := GetSdk(f.ctx).VPC().Network().List(f.ctx, &vpc.ListNetworksRequest{
		FolderId: GetFolderID(f.ctx),
		Filter:   `name = '` + name + `'`,
	})

	if err != nil {
		panic(err)
	}

	if len(n.Networks) == 0 {
		return nil
	}

	return NewNetwork(f.ctx, n.Networks[0])
}

func (f Folder) Networks() []Network {
	n, err := GetSdk(f.ctx).VPC().Network().List(f.ctx, &vpc.ListNetworksRequest{
		FolderId: GetFolderID(f.ctx),
	})

	if err != nil {
		panic(err)
	}

	networks := make([]Network, 0, len(n.Networks))
	for _, network := range n.Networks {
		networks = append(networks, *NewNetwork(f.ctx, network))
	}

	return networks
}

func (f Folder) ServiceAccounts() []ServiceAccount {
	sa, err := GetSdk(f.ctx).IAM().ServiceAccount().List(f.ctx, &iam.ListServiceAccountsRequest{
		FolderId: GetFolderID(f.ctx),
	})

	if err != nil {
		panic(err)
	}

	accounts := make([]ServiceAccount, 0, len(sa.ServiceAccounts))
	for _, account := range sa.ServiceAccounts {
		accounts = append(accounts, *NewServiceAccount(f.ctx, account))
	}
	return accounts
}

func (f Folder) Instances() []Instance {
	ins, err := GetSdk(f.ctx).Compute().Instance().List(f.ctx, &compute.ListInstancesRequest{
		FolderId: GetFolderID(f.ctx),
	})

	if err != nil {
		panic(err)
	}

	instances := make([]Instance, 0, len(ins.Instances))
	for _, instance := range ins.Instances {
		instances = append(instances, *NewInstance(f.ctx, instance))
	}
	return instances
}
