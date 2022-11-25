package tools

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	log "github.com/sirupsen/logrus"
)

type LineBotSetting struct {
	OwnerID       string `mapstructure:"LineBotUserID"`
	ChannelSecret string `mapstructure:"LineBotChannelSecret"`
	ChannelToken  string `mapstructure:"LineBotChannelToken"`
}

type LineBotManager struct {
	Setting *LineBotSetting
	Client  *linebot.Client
}

func (Manager *LineBotManager) Init() {
	bot, err := linebot.New(Manager.Setting.ChannelSecret, Manager.Setting.ChannelToken)
	if err != nil {
		panic(err)
	}
	Manager.Client = bot
}

func (Manager *LineBotManager) IsConnected() bool {
	message := linebot.NewTextMessage("test")
	_, err := Manager.Client.PushMessage(Manager.Setting.OwnerID, message).Do()
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func (Manager *LineBotManager) ProvideDBConnection() any {
	return Manager.Client
}

func NewLineBotManager(Setting *LineBotSetting) *LineBotManager {
	Manager := LineBotManager{Setting: Setting}
	Manager.Init()
	return &Manager
}
