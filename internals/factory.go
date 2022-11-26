package internals

import (
	_ "LineMongo/swaggerDocs" // 匯入swagger docs
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewServer(mode string, routes []func(router *gin.Engine), middlewares ...gin.HandlerFunc) (server *gin.Engine) {
	gin.SetMode(mode)
	server = gin.New()
	server.Use(middlewares...)
	for _, route := range routes {
		route(server)
	}

	if mode := gin.Mode(); mode == gin.DebugMode {
		server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return server
}
