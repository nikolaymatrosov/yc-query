package query

import (
	"context"
	"fmt"
	"time"

	"github.com/sosodev/duration"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type Env struct {
	ctx context.Context
}

func NewEnv(ctx context.Context, sdk *ycsdk.SDK) Env {
	return Env{
		ctx: WithSdk(ctx, sdk),
	}
}

// Cloud finds a cloud by name.
func (e Env) Cloud(name string) Cloud {
	c, err := GetSdk(e.ctx).ResourceManager().Cloud().List(e.ctx, &resourcemanager.ListCloudsRequest{
		Filter: `name = '` + name + `'`,
	})

	if err != nil {
		panic(err)
	}

	if len(c.Clouds) == 0 {
		return Cloud{}
	}

	return Cloud{
		ctx:   WithCloudID(e.ctx, c.Clouds[0].Id),
		Cloud: c.Clouds[0],
	}
}

// CloudById finds a cloud by ID.
func (e Env) CloudById(id string) Cloud {
	c, err := GetSdk(e.ctx).ResourceManager().Cloud().Get(e.ctx, &resourcemanager.GetCloudRequest{
		CloudId: id,
	})

	if err != nil {
		panic(err)
	}

	return Cloud{
		ctx:   WithCloudID(e.ctx, c.Id),
		Cloud: c,
	}
}

// Clouds returns all clouds.
func (e Env) Clouds() []Cloud {
	c, err := GetSdk(e.ctx).ResourceManager().Cloud().List(e.ctx, &resourcemanager.ListCloudsRequest{})

	if err != nil {
		panic(err)
	}

	clouds := make([]Cloud, 0, len(c.Clouds))
	for _, cloud := range c.Clouds {
		clouds = append(clouds, Cloud{
			ctx:   WithCloudID(e.ctx, cloud.Id),
			Cloud: cloud,
		})
	}
	return clouds
}

// Organization finds an organization by name.
func (e Env) Organization(name string) Organization {
	o, err := GetSdk(e.ctx).OrganizationManager().Organization().List(e.ctx, &organizationmanager.ListOrganizationsRequest{
		Filter: `name = '` + name + `'`,
	})

	if err != nil {
		panic(err)
	}

	if len(o.Organizations) == 0 {
		return Organization{}
	}

	return Organization{
		ctx:          WithOrganizationID(e.ctx, o.Organizations[0].Id),
		Organization: o.Organizations[0],
	}
}

// OrganizationById finds an organization by ID.
func (e Env) OrganizationById(id string) Organization {
	o, err := GetSdk(e.ctx).OrganizationManager().Organization().Get(e.ctx, &organizationmanager.GetOrganizationRequest{
		OrganizationId: id,
	})

	if err != nil {
		panic(err)
	}

	return Organization{
		ctx:          WithOrganizationID(e.ctx, o.Id),
		Organization: o,
	}
}

// FormatSize converts size to human-readable format.
func (e Env) FormatSize(size int) string {
	// Convert size to human-readable format
	if size < 1024 {
		return fmt.Sprintf("%d B", size)
	}
	size /= 1024
	if size < 1024 {
		return fmt.Sprintf("%d KB", size)
	}
	size /= 1024
	if size < 1024 {
		return fmt.Sprintf("%d MB", size)
	}
	size /= 1024
	return fmt.Sprintf("%d GB", size)
}

func (e Env) Duration(in string) time.Duration {
	d, err := duration.Parse(in)
	if err != nil {
		return 0
	}
	return d.ToTimeDuration()
}
