package persistence

import (
	"opendoor/config"
	"opendoor/domain/entity"
	"opendoor/domain/repository"
	"opendoor/infrastructure/tencentcloud"

	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

type Repositories struct {
	Firewall        repository.IFirewallRepo
	clientFirewalls map[*lighthouse.Client]entity.Firewall
}

func NewRepositories(servers []config.Server) (*Repositories, error) {
	repos := &Repositories{
		clientFirewalls: make(map[*lighthouse.Client]entity.Firewall),
	}

	for _, server := range servers {
		client, err := tencentcloud.CreateClient(server.SecretID, server.SecretKey, server.Region, "lighthouse.tencentcloudapi.com")
		if err != nil {
			return nil, err
		}
		firewall := entity.Firewall{
			InstanceID: server.InstanceID,
			Tag:        server.RuleTag,
		}
		repos.clientFirewalls[client] = firewall
	}
	repos.Firewall = NewFirewallRepo(repos.clientFirewalls)
	return repos, nil
}
