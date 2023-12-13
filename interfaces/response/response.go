package response

import (
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

// APIResponse API 响应
type APIResponse struct {
	RequestID string      `json:"request-id"`
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		RequestID: requestid.Get(c),
		Code:      0,
		Message:   "success",
		Data:      data,
	})
}

// Error 错误响应
func Error(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		RequestID: requestid.Get(c),
		Code:      1,
		Message:   "error",
		Data:      data,
	})
}

// Forbidden 禁止响应
func Forbidden(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusForbidden, APIResponse{
		RequestID: requestid.Get(c),
		Code:      403,
		Message:   message,
		Data:      data,
	})
}

// Unauthorized 未授权响应
func Unauthorized(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusUnauthorized, APIResponse{
		RequestID: requestid.Get(c),
		Code:      401,
		Message:   message,
		Data:      data,
	})
}
