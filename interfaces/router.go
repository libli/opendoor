package interfaces

import (
	"opendoor/application"
	"opendoor/infrastructure/persistence"
	"opendoor/interfaces/handler"
	"opendoor/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter(r *gin.Engine, repo *persistence.Repositories, token string) {
	healthHandler := handler.NewHealthHandler()
	r.GET("/healthz", healthHandler.Healthz)

	apiGroup := r.Group("/api", middleware.Auth(token))
	{
		firewallApp := application.NewFirewallApp(repo.Firewall)
		firewallHandler := handler.NewFirewallHandler(firewallApp)
		apiGroup.POST("/firewall", firewallHandler.Create)
	}
}
