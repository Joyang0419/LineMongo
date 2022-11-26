package main

import (
	"LineMongo/internals"
	"LineMongo/internals/middleware"
	"LineMongo/internals/routers/lineRouter"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// @title LineMongo
// @version 1.0
// @description LineMongo swagger
// @host 0.0.0.0:8080
// schemes http
func main() {
	middlewares := []gin.HandlerFunc{
		middleware.CorsMiddleware(),
		gin.Recovery(), // Recovery middleware recovers from any panics and writes a 500 if there was one.
		middleware.JSONLogMiddleware(true, log.DebugLevel),
	}

	server := internals.NewServer(
		gin.DebugMode, []func(router *gin.Engine){lineRouter.Route}, middlewares...)
	err := server.Run()
	if err != nil {
		panic(err)
	}
}
