package query

import (
	"context"

	"github.com/yandex-cloud/go-sdk"
)

const (
	cloudIDContextKey          = "cloudID"
	diskIDContextKey           = "diskID"
	groupIDContextKey          = "groupID"
	instanceIDContextKey       = "instanceID"
	networkIDContextKey        = "networkID"
	organizationIDContextKey   = "organizationID"
	serviceAccountIDContextKey = "serviceAccountID"
	subnetIDContextKey         = "subnetID"

	sdkContextKey = "sdk"
)

func WithServiceAccountID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, serviceAccountIDContextKey, id)
}

func GetServiceAccountID(ctx context.Context) string {
	return ctx.Value(serviceAccountIDContextKey).(string)
}

func WithNetworkID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, networkIDContextKey, id)
}

func GetNetworkID(ctx context.Context) string {
	return ctx.Value(networkIDContextKey).(string)
}

func WithSubnetID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, subnetIDContextKey, id)
}

func WithSdk(ctx context.Context, sdk *ycsdk.SDK) context.Context {
	return context.WithValue(ctx, sdkContextKey, sdk)
}

func GetSdk(ctx context.Context) *ycsdk.SDK {
	return ctx.Value(sdkContextKey).(*ycsdk.SDK)
}

func WithCloudID(ctx context.Context, cloudID string) context.Context {
	return context.WithValue(ctx, cloudIDContextKey, cloudID)
}

func GetCloudID(ctx context.Context) string {
	return ctx.Value(cloudIDContextKey).(string)
}

func WithInstanceID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, instanceIDContextKey, id)
}

func GetInstanceID(ctx context.Context) string {
	return ctx.Value(instanceIDContextKey).(string)
}

func WithDiskID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, diskIDContextKey, id)
}

func GetDiskID(ctx context.Context) string {
	return ctx.Value(diskIDContextKey).(string)
}

func WithOrganizationID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, organizationIDContextKey, id)
}

func GetOrganizationID(ctx context.Context) string {
	return ctx.Value(organizationIDContextKey).(string)
}

func WithGroupID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, groupIDContextKey, id)
}

func GetGroupID(ctx context.Context) string {
	return ctx.Value(groupIDContextKey).(string)
}
