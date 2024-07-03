package query

import (
	"context"
	"time"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
)

type ServiceAccount struct {
	ctx       context.Context
	CreatedAt time.Time `json:"created_at"`
	*iam.ServiceAccount
}

func NewServiceAccount(ctx context.Context, sa *iam.ServiceAccount) *ServiceAccount {
	return &ServiceAccount{ctx: WithServiceAccountID(ctx, sa.Id), ServiceAccount: sa, CreatedAt: sa.CreatedAt.AsTime()}
}

type AccessBinding struct {
	ctx         context.Context
	RoleId      string
	SubjectId   string
	SubjectType string
}

func (sa ServiceAccount) FolderAccessBindings() []AccessBinding {
	b, err := GetSdk(sa.ctx).ResourceManager().Folder().ListAccessBindings(sa.ctx, &access.ListAccessBindingsRequest{
		ResourceId: GetFolderID(sa.ctx),
	})

	if err != nil {
		panic(err)
	}

	var accessBindings []AccessBinding
	for _, binding := range b.AccessBindings {
		if binding.Subject.Type != "serviceAccount" || binding.Subject.Id != sa.Id {
			continue
		}
		accessBindings = append(accessBindings, AccessBinding{
			ctx:         sa.ctx,
			RoleId:      binding.RoleId,
			SubjectId:   binding.Subject.Id,
			SubjectType: binding.Subject.Type,
		})
	}

	return accessBindings
}

func (sa ServiceAccount) CloudAccessBindings() []AccessBinding {
	b, err := GetSdk(sa.ctx).ResourceManager().Cloud().ListAccessBindings(sa.ctx, &access.ListAccessBindingsRequest{
		ResourceId: GetCloudID(sa.ctx),
	})

	if err != nil {
		panic(err)
	}

	var accessBindings []AccessBinding
	for _, binding := range b.AccessBindings {
		if binding.Subject.Type != "serviceAccount" || binding.Subject.Id != sa.Id {
			continue
		}
		accessBindings = append(accessBindings, AccessBinding{
			ctx:         sa.ctx,
			RoleId:      binding.RoleId,
			SubjectId:   binding.Subject.Id,
			SubjectType: binding.Subject.Type,
		})
	}

	return accessBindings
}
