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
	Token   string
	Servers []config.Server
}

// NewFirewallHandler creates a new FirewallHandler.
func NewFirewallHandler(token string, servers []config.Server) *FirewallHandler {
	return &FirewallHandler{
		Token:   token,
		Servers: servers,
	}
}

// Update is the handler for the firewall.
func (h *FirewallHandler) Update(c *gin.Context) {
	token := c.Query("token")
	if token != h.Token {
		response.Forbidden(c, "Unauthorized", nil)
		return
	}
	clientIP := c.ClientIP()
	log.Info("收到新请求", zap.String("ip", clientIP))
	response.Success(c, gin.H{"ip": clientIP})

	////firewall.CreateRule(h.hkClient)
	//fmt.Fprintf(w, "您的 IP 地址是: %s\n", userIP)
}
