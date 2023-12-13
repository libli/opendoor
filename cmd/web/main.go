package main

import (
	"time"

	"opendoor/config"
	"opendoor/infrastructure/persistence"
	"opendoor/interfaces"
	"opendoor/log"

	"github.com/gin-contrib/requestid"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const configPath = "config.yaml"

func main() {
	defer log.Sync()
	// 读取配置文件
	cfg, err := config.Get(configPath)
	if err != nil {
		log.Fatal("read config error", zap.Error(err))
	}

	gin.SetMode(cfg.Gin.Mode)
	r := gin.New()

	// requestID 中间件
	r.Use(requestid.New())

	r.Use(ginzap.GinzapWithConfig(log.Logger(), &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        false,                                // 使用本地时间
		SkipPaths:  []string{"/healthz", "/favicon.ico"}, // 不记录日志的路径
		Context: func(c *gin.Context) []zapcore.Field {
			return []zapcore.Field{
				zap.String("request-id", requestid.Get(c)),
			}
		},
	}))
	r.Use(ginzap.RecoveryWithZap(log.Logger(), true))

	repo, err := persistence.NewRepositories(cfg.Servers)
	if err != nil {
		log.Fatal("create repositories error", zap.Error(err))
	}
	// 设置路由
	interfaces.SetupRouter(r, repo, cfg.Token)

	// 启动服务
	if err := r.Run(":" + cfg.Gin.Port); err != nil {
		log.Fatal("server start failed", zap.Error(err))
	}
}
