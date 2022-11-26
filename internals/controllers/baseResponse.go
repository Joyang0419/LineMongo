package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseRouteResponse struct{}

func (base *BaseRouteResponse) GinSuccessResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

func (base *BaseRouteResponse) GinErrorResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}
