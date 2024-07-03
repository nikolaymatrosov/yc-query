package query

import (
	"context"
	"time"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
)

type Cloud struct {
	ctx       context.Context
	CreatedAt time.Time `json:"created_at"`
	*resourcemanager.Cloud
}

func NewCloud(ctx context.Context, cloud *resourcemanager.Cloud) *Cloud {
	return &Cloud{ctx: WithCloudID(ctx, cloud.Id), Cloud: cloud, CreatedAt: cloud.CreatedAt.AsTime()}
}

func (c Cloud) Folder(name string) *Folder {
	f, err := GetSdk(c.ctx).ResourceManager().Folder().List(c.ctx, &resourcemanager.ListFoldersRequest{
		CloudId: GetCloudID(c.ctx),
		Filter:  `name = '` + name + `'`,
	})

	if err != nil {
		panic(err)
	}

	if len(f.Folders) == 0 {
		return nil
	}

	return NewFolder(c.ctx, f.Folders[0])
}

func WithFolderID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, "folderID", id)
}

func GetFolderID(ctx context.Context) string {
	return ctx.Value("folderID").(string)
}

func (c Cloud) Folders() []Folder {
	f, err := GetSdk(c.ctx).ResourceManager().Folder().List(c.ctx, &resourcemanager.ListFoldersRequest{
		CloudId: GetCloudID(c.ctx),
	})

	if err != nil {
		panic(err)
	}

	folders := make([]Folder, 0, len(f.Folders))
	for _, folder := range f.Folders {
		folders = append(folders, *NewFolder(c.ctx, folder))
	}
	return folders
}
