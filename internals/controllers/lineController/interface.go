package lineController

import "github.com/gin-gonic/gin"

type InterfaceLineController interface {
	Callback(c *gin.Context)
	PushMessages(c *gin.Context)
	ReplyMessages(c *gin.Context)
	GetUserMessages(c *gin.Context)
}
