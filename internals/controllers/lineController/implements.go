package lineController

import (
	"LineMongo/internals/services/lineService"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type V1 struct {
	LineService lineService.InterfaceLineService
}

func NewV1LineController(LineService lineService.InterfaceLineService) *V1 {
	return &V1{
		LineService: LineService,
	}
}

// Callback
// @Router /line/callback [get]
// @Accept  json
// @Produce  json
// @Success 200 {object} CallbackSuccessResponse
// @Failure 400 {object} CallbackErrorResponse
func (controller V1) Callback(c *gin.Context) {
	err := controller.LineService.ReceiveEvents(c.Request)
	if err != nil {
		log.Error("Please Provide Line signature")
		errResponse := CallbackErrorResponse{StatusCode: http.StatusBadRequest, Data: "Please Provide Line signature"}
		errResponse.GinErrorResponse(c, errResponse)
		return
	}

	successResponse := CallbackSuccessResponse{StatusCode: http.StatusOK, Data: "Line web socket Success"}
	successResponse.GinSuccessResponse(c, successResponse)
}

// PushMessages
// @Router /line/pushMessages [post]
// @Accept  json
// @Produce  json
// @param PushMessagesBody body PushMessagesBody true "PushMessagesRequired"
// @Success 200 {object} PushMessagesSuccessResponse
// @Failure 400 {object} PushMessagesErrorResponse
func (controller V1) PushMessages(c *gin.Context) {
	var (
		body *PushMessagesBody
		err  error
	)
	err = c.BindJSON(&body)
	if err != nil {
		errResponse := PushMessagesErrorResponse{StatusCode: http.StatusBadRequest, Data: "Wrong body"}
		errResponse.GinErrorResponse(c, errResponse)
		return
	}
	err = controller.LineService.PushMessages(body.UserID, body.Messages)
	if err != nil {
		errResponse := PushMessagesErrorResponse{StatusCode: http.StatusBadRequest, Data: "Wrong userID"}
		errResponse.GinErrorResponse(c, errResponse)
		return
	}
	successResponse := PushMessagesSuccessResponse{StatusCode: http.StatusOK, Data: body}
	successResponse.GinSuccessResponse(c, successResponse)
}

// ReplyMessages
// @Router /line/replyMessages [post]
// @Accept  json
// @Produce  json
// @param ReplyMessagesBody body ReplyMessagesBody true "ReplyMessagesRequired"
// @Success 200 {object} ReplyMessagesSuccessResponse
// @Failure 400 {object} ReplyMessagesErrorResponse
func (controller V1) ReplyMessages(c *gin.Context) {
	var (
		body *ReplyMessagesBody
		err  error
	)
	err = c.BindJSON(&body)
	if err != nil {
		errResponse := ReplyMessagesErrorResponse{StatusCode: http.StatusBadRequest, Data: "Wrong body"}
		errResponse.GinErrorResponse(c, errResponse)
		return
	}
	err = controller.LineService.ReplyMessages(body.ReplyToken, body.Messages)
	if err != nil {
		errResponse := ReplyMessagesErrorResponse{StatusCode: http.StatusBadRequest, Data: "Wrong replyToken"}
		errResponse.GinErrorResponse(c, errResponse)
		return
	}
	successResponse := ReplyMessagesSuccessResponse{StatusCode: http.StatusOK, Data: body}
	successResponse.GinSuccessResponse(c, successResponse)
}

// GetUserMessages
// @Router /line/userMessages [get]
// @Accept  json
// @Produce  json
// @param userID query string true "用戶ID" default("Ubbd98d38fbfc1fe0e7db93a2d8bc9c34")
// @Success 200 {object} GetUserMessagesSuccessResponse
// @Failure 400 {object} GetUserMessagesErrorResponse
func (controller V1) GetUserMessages(c *gin.Context) {
	_ = c
	//userid
	//messages
}
