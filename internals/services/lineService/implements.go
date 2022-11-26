package lineService

import (
	"LineMongo/internals/repos/userMessageRepo"
	"LineMongo/internals/tools"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"net/http"
)

type LineService struct {
	UserMessageRepo userMessageRepo.InterfaceUserMessageRepo
	LineBotManager  *tools.LineBotManager
}

func (service *LineService) ReceiveEvents(request *http.Request) (err error) {
	events, err := service.LineBotManager.Client.ParseRequest(request)
	if err != nil {
		return err
	}
	for _, event := range events {
		service.UserMessageRepo.Create(event)
	}
	return nil
}

func (service *LineService) PushMessages(userID string, messages []string) error {
	linebotMessages := make([]linebot.SendingMessage, 0, len(messages))
	for _, message := range messages {
		linebotMessages = append(linebotMessages, linebot.NewTextMessage(message))
	}

	_, err := service.LineBotManager.Client.PushMessage(userID, linebotMessages...).Do()
	if err != nil {
		return err
	}
	// todo
	// 先去找userID 存在
	return nil
}

func (service *LineService) ReplyMessages(replyToken string, messages []string) error {
	linebotMessages := make([]linebot.SendingMessage, 0, len(messages))
	for _, message := range messages {
		linebotMessages = append(linebotMessages, linebot.NewTextMessage(message))
	}

	_, err := service.LineBotManager.Client.ReplyMessage(replyToken, linebotMessages...).Do()
	if err != nil {
		return err
	}
	// todo
	// replytoken 存在
	return nil
}

func NewLineService(userMessageRepo userMessageRepo.InterfaceUserMessageRepo,
	LineBotManager *tools.LineBotManager) *LineService {
	return &LineService{UserMessageRepo: userMessageRepo, LineBotManager: LineBotManager}
}
