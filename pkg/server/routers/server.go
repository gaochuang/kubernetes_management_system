package routers

import (
	"github.com/gin-gonic/gin"
	"kubernetes_management_system/common"
	"kubernetes_management_system/pkg/server/middleware"
)

func InitServer() {
	r := gin.Default()

	r.Use(middleware.Cores())
	r.Use(gin.LoggerWithFormatter(middleware.LogsFormatDefine))

	publicGroup := r.Group("/api/v1")
	{
		User(publicGroup)
	}

	privateGroup := r.Group("/api/v1")
	privateGroup.Use(middleware.AuthMiddleware())
	{
		InitUserInfo(privateGroup)
		InitClusterRouter(privateGroup)
	}

	common.LOG.Info("server addr: " + common.CONFIG.System.Addr)
	r.Run(common.CONFIG.System.Addr)
}
