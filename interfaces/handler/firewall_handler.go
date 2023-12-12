package handler

import (
	"opendoor/config"
	"opendoor/interfaces/response"
	"opendoor/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// FirewallHandler is the handler for the firewall.
type FirewallHandler struct {
	Servers []config.Server
}

// NewFirewallHandler creates a new FirewallHandler.
func NewFirewallHandler(servers []config.Server) *FirewallHandler {
	return &FirewallHandler{
		Servers: servers,
	}
}

// Create 添加新的防火墙规则.
func (h *FirewallHandler) Create(c *gin.Context) {
	clientIP := c.ClientIP()
	log.Info("收到新请求", zap.String("ip", clientIP))
	response.Success(c, gin.H{"ip": clientIP})

	////firewall.CreateRule(h.hkClient)
	//fmt.Fprintf(w, "您的 IP 地址是: %s\n", userIP)
}
