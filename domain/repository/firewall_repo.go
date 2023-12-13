package repository

import (
	"context"
)

type IFirewallRepo interface {
	Create(ctx context.Context, ip string) error
}
