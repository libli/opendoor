package handler

import (
	"opendoor/interfaces/response"

	"github.com/gin-gonic/gin"
)

// HealthHandler is the handler for the health check.
type HealthHandler struct {
}

// NewHealthHandler creates a new HealthHandler.
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Healthz is the handler for the health check.
func (h *HealthHandler) Healthz(c *gin.Context) {
	response.Success(c, nil)
}
