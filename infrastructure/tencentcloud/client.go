package tencentcloud

import (
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

// CreateClient 创建腾讯云轻量服务调用 Client
func CreateClient(secretId string, secretKey string, region string, endpoint string) (*lighthouse.Client, error) {
	credential := common.NewCredential(
		secretId,
		secretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = endpoint
	client, err := lighthouse.NewClient(credential, region, cpf)
	if err != nil {
		return nil, fmt.Errorf("CreateClient: error creating client with secretId %s: %w", secretId, err)
	}

	return client, nil
}
