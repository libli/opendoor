package interfaces

import (
	"opendoor/config"
	"opendoor/interfaces/handler"

	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter(r *gin.Engine, cfg config.Config) {
	healthHandler := handler.NewHealthHandler()
	r.GET("/healthz", healthHandler.Healthz)

	firewallHandler := handler.NewFirewallHandler(cfg.Token, cfg.Servers)
	r.GET("/firewall", firewallHandler.Update)
}
