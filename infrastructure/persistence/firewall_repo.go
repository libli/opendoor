package persistence

import (
	"context"
	"fmt"
	"strings"

	"opendoor/domain/entity"
	"opendoor/domain/repository"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

type FirewallRepo struct {
	clientFirewalls map[*lighthouse.Client]entity.Firewall
}

func NewFirewallRepo(clients map[*lighthouse.Client]entity.Firewall) *FirewallRepo {
	return &FirewallRepo{
		clientFirewalls: clients,
	}
}

var _ repository.IFirewallRepo = (*FirewallRepo)(nil)

func (f FirewallRepo) Create(ctx context.Context, ip string) error {
	var errors []string
	for client, firewall := range f.clientFirewalls {
		rule := &lighthouse.FirewallRule{
			Protocol:                common.StringPtr("TCP"),
			Port:                    common.StringPtr("ALL"),
			CidrBlock:               &ip,
			Action:                  common.StringPtr("ACCEPT"),
			FirewallRuleDescription: &firewall.Tag,
		}
		request := lighthouse.NewCreateFirewallRulesRequest()
		request.InstanceId = &firewall.InstanceID
		request.FirewallRules = []*lighthouse.FirewallRule{rule}

		_, err := client.CreateFirewallRules(request)
		if err != nil {
			errors = append(errors, fmt.Sprintf("instance %s: %v", firewall.InstanceID, err))
			continue
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("CreateRule: encountered errors: %s", strings.Join(errors, "; "))
	}
	return nil
}
