package application

import (
	"context"

	"opendoor/domain/repository"
)

// IFirewallApp 防火墙应用服务接口
type IFirewallApp interface {
	// Add 为所有服务器实例添加指定的防火墙规则
	Add(ctx context.Context, ip string) error
}

type FirewallApp struct {
	firewallRepo repository.IFirewallRepo
}

var _ IFirewallApp = (*FirewallApp)(nil)

func NewFirewallApp(firewallRepo repository.IFirewallRepo) *FirewallApp {
	return &FirewallApp{
		firewallRepo: firewallRepo,
	}
}

func (app *FirewallApp) Add(ctx context.Context, ip string) error {
	return app.firewallRepo.Create(ctx, ip)
}
