package implements

import (
	"LineMongo"
	"LineMongo/internals/tools"
)

var (
	MongoDBManager = tools.NewMongoDBManager(LineMongo.Config.MongoDBSetting)
	LineBotManager = tools.NewLineBotManager(LineMongo.Config.LineBotSetting)
)
