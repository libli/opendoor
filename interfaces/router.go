package interfaces

import (
	"opendoor/config"
	"opendoor/interfaces/handler"
	"opendoor/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter(r *gin.Engine, cfg config.Config) {
	healthHandler := handler.NewHealthHandler()
	r.GET("/healthz", healthHandler.Healthz)

	apiGroup := r.Group("/api", middleware.Auth(cfg.Token))
	{
		firewallHandler := handler.NewFirewallHandler(cfg.Servers)
		apiGroup.POST("/firewall", firewallHandler.Create)
	}
}
