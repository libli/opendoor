package logic

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

// CreateRule 创建防火墙规则
func CreateRule(client *lighthouse.Client, instanceID, tag, ip string) error {
	// 创建新规则
	rule := &lighthouse.FirewallRule{
		Protocol:                common.StringPtr("TCP"),
		Port:                    common.StringPtr("ALL"),
		CidrBlock:               &ip,
		Action:                  common.StringPtr("ACCEPT"),
		FirewallRuleDescription: &tag,
	}
	request := lighthouse.NewCreateFirewallRulesRequest()
	request.InstanceId = &instanceID
	request.FirewallRules = []*lighthouse.FirewallRule{rule}

	response, err := client.CreateFirewallRules(request)
	if err != nil {
		return fmt.Errorf("CreateRule: error creating firewall rule for instance %s: %w", instanceID, err)
	}
	if response.Response == nil {
		return fmt.Errorf("CreateRule: response is nil for instance %s", instanceID)
	}
	return nil
}
