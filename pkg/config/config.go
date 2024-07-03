package config

import (
	"os"
	"time"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
	"github.com/yandex-cloud/go-sdk/iamkey"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gopkg.in/yaml.v3"
)

type ServiceAccountKey struct {
	Id               string    `json:"id"`
	ServiceAccountId string    `json:"service_account_id"`
	CreatedAt        time.Time `json:"created_at"`
	KeyAlgorithm     string    `json:"key_algorithm"`
	PublicKey        string    `json:"public_key"`
	PrivateKey       string    `json:"private_key"`
}

func (k ServiceAccountKey) IamKey() *iamkey.Key {
	return &iamkey.Key{
		Id:           k.Id,
		Subject:      &iamkey.Key_ServiceAccountId{ServiceAccountId: k.ServiceAccountId},
		CreatedAt:    timestamppb.New(k.CreatedAt),
		Description:  "",
		KeyAlgorithm: iam.Key_Algorithm(iam.Key_Algorithm_value[k.KeyAlgorithm]),
		PublicKey:    k.PublicKey,
		PrivateKey:   k.PrivateKey,
	}
}

type Profile struct {
	ServiceAccountKey  ServiceAccountKey `json:"service-account-key"`
	Token              string            `json:"token"`
	CloudId            string            `json:"cloud-id"`
	FolderId           string            `json:"folder-id"`
	ComputeDefaultZone string            `json:"compute-default-zone"`
}

type Config struct {
	Current  string             `json:"current"`
	Profiles map[string]Profile `json:"profiles"`
}

func (c Config) CurrentProfile() Profile {
	return c.Profiles[c.Current]
}

func Parse(filename string) (*Config, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config

	if err = yaml.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
