package lineRouter

import (
	"LineMongo/internals/implements"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	lineGroup := router.Group("line")
	lineGroup.POST("callback", implements.LineController.Callback)
	lineGroup.POST("pushMessages", implements.LineController.PushMessages)
	lineGroup.POST("replyMessages", implements.LineController.ReplyMessages)
}
