package handler

import (
	"opendoor/application"
	"opendoor/interfaces/response"
	"opendoor/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// FirewallHandler is the handler for the firewall.
type FirewallHandler struct {
	firewallApp application.IFirewallApp
}

// NewFirewallHandler creates a new FirewallHandler.
func NewFirewallHandler(app application.IFirewallApp) *FirewallHandler {
	return &FirewallHandler{
		firewallApp: app,
	}
}

// Create 添加新的防火墙规则.
func (h *FirewallHandler) Create(c *gin.Context) {
	clientIP := c.ClientIP()

	if err := h.firewallApp.Add(c, clientIP); err != nil {
		log.Error("添加防火墙规则失败", zap.Error(err))
		response.Error(c, err)
		return
	}

	log.Info("添加防火墙规则成功", zap.String("ip", clientIP))
	response.Success(c, nil)
}
