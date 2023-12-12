package helper

import (
	"net"
	"net/http"
)

// handleCORSRequest handles the CORS request.
func handleCORSRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(http.StatusOK)
		return
	}
}

// authRequest 检查请求中的 token 参数是否与给定的 token 匹配
func authRequest(r *http.Request, token string) bool {
	// 从请求中解析查询参数
	queryToken := r.URL.Query().Get("token")

	// 比较查询参数中的 token 与给定的 token 是否相同
	return queryToken == token
}

// getUserIP 用于从 http 请求中提取用户的 IP 地址
func getUserIP(r *http.Request) (string, string) {
	var userIP, ipSource string

	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		userIP = xForwardedFor
		ipSource = "X-Forwarded-For"
	} else {
		xRealIP := r.Header.Get("X-Real-IP")
		if xRealIP != "" {
			userIP = xRealIP
			ipSource = "X-Real-IP"
		} else {
			userIP = r.RemoteAddr
			ipSource = "RemoteAddr"
		}
	}

	userIP, _, _ = net.SplitHostPort(userIP)
	return userIP, ipSource
}
