package lineService

import (
	"net/http"
)

type InterfaceLineService interface {
	ReceiveEvents(request *http.Request) error
	PushMessages(userID string, messages []string) error
	ReplyMessages(replyToken string, messages []string) error
	GetUserMessages(userID string) []map[string]any
}
