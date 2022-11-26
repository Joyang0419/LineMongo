package lineController

import (
	"LineMongo/internals/controllers"
)

type CallbackSuccessResponse struct {
	StatusCode int    `example:"200"`
	Data       string `example:"Line web socket Success"`
	routers.BaseRouteResponse
}

type CallbackErrorResponse struct {
	StatusCode int    `example:"400"`
	Data       string `example:"Please Provide Line signature"`
	routers.BaseRouteResponse
}

type PushMessagesSuccessResponse struct {
	StatusCode int `example:"200"`
	Data       *PushMessagesBody
	routers.BaseRouteResponse
}

type PushMessagesErrorResponse struct {
	StatusCode int    `example:"400"`
	Data       string `example:"Wrong userID"`
	routers.BaseRouteResponse
}

type ReplyMessagesSuccessResponse struct {
	StatusCode int `example:"400"`
	Data       *ReplyMessagesBody
	routers.BaseRouteResponse
}

type ReplyMessagesErrorResponse struct {
	StatusCode int    `example:"400"`
	Data       string `example:"Wrong ReplyToken"`
	routers.BaseRouteResponse
}

type userMessagesData struct {
	CreatedAt string `example:"2022-11-26T14:28:58.526+00:00"`
	Message   string `example:"Hello World"`
}

type GetUserMessagesSuccessResponse struct {
	StatusCode int `example:"400"`
	Data       []userMessagesData
	routers.BaseRouteResponse
}

type GetUserMessagesErrorResponse struct {
	StatusCode int    `example:"400"`
	Data       string `example:"Please Provide userID"`
	routers.BaseRouteResponse
}
