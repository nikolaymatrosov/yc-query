package query

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/nikolaymatrosov/yc-query/pkg/filter"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
)

type Instance struct {
	ctx       context.Context
	CreatedAt time.Time `json:"created_at"`
	*compute.Instance
}

func NewInstance(ctx context.Context, instance *compute.Instance) *Instance {
	return &Instance{ctx: WithInstanceID(ctx, instance.Id), Instance: instance, CreatedAt: instance.CreatedAt.AsTime()}
}

type Disk struct {
	ctx       context.Context
	CreatedAt time.Time `json:"created_at"`
	*compute.Disk
}

func NewDisk(ctx context.Context, disk *compute.Disk) *Disk {
	return &Disk{ctx: WithDiskID(ctx, disk.Id), Disk: disk, CreatedAt: disk.CreatedAt.AsTime()}
}

func (f Folder) Instance(name string) *Instance {
	resp, err := GetSdk(f.ctx).Compute().Instance().List(f.ctx, &compute.ListInstancesRequest{
		FolderId: GetFolderID(f.ctx),
		Filter:   `name = '` + name + `'`,
	})

	if err != nil {
		panic(err)
	}

	if len(resp.Instances) == 0 {
		return nil
	}

	i := resp.Instances[0]

	return NewInstance(f.ctx, i)
}

func (i Instance) BootDisk() *Disk {
	d, err := GetSdk(i.ctx).Compute().Disk().Get(i.ctx, &compute.GetDiskRequest{
		DiskId: i.Instance.BootDisk.DiskId,
	})

	if err != nil {
		panic(fmt.Errorf("boot disk: %w", err))
	}

	return NewDisk(i.ctx, d)
}

func (i Instance) SecondaryDisks() []Disk {
	if len(i.Instance.SecondaryDisks) == 0 {
		return []Disk{}
	}

	var diskIds []string
	for _, d := range i.Instance.SecondaryDisks {
		diskIds = append(diskIds, d.DiskId)
	}

	d, err := GetSdk(i.ctx).Compute().Disk().List(i.ctx, &compute.ListDisksRequest{
		FolderId: GetFolderID(i.ctx),
		Filter:   filter.In("id", diskIds...),
	})

	if err != nil {
		panic(err)
	}

	disks := make([]Disk, 0, len(d.Disks))
	for _, disk := range d.Disks {
		disks = append(disks, *NewDisk(i.ctx, disk))
	}
	return disks
}

func (i Instance) Disk(name string) *Disk {
	d, err := GetSdk(i.ctx).Compute().Disk().List(i.ctx, &compute.ListDisksRequest{
		Filter: `name = '` + name + `'`,
	})

	if err != nil {
		panic(err)
	}

	if len(d.Disks) == 0 {
		return nil
	}

	return NewDisk(i.ctx, d.Disks[0])
}

func (i Instance) Disks() []Disk {
	d, err := GetSdk(i.ctx).Compute().Disk().List(i.ctx, &compute.ListDisksRequest{
		FolderId: GetFolderID(i.ctx),
	})

	if err != nil {
		panic(err)
	}

	disks := make([]Disk, 0, len(d.Disks))
	for _, disk := range d.Disks {
		if !slices.Contains(disk.InstanceIds, i.Id) {
			continue
		}
		disks = append(disks, *NewDisk(i.ctx, disk))
	}

	return disks
}

func (i Instance) ServiceAccount() *ServiceAccount {
	sa, err := GetSdk(i.ctx).IAM().ServiceAccount().Get(i.ctx, &iam.GetServiceAccountRequest{
		ServiceAccountId: i.ServiceAccountId,
	})

	if err != nil {
		panic(err)
	}

	return NewServiceAccount(i.ctx, sa)
}
