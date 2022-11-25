package implements

import (
	"LineMongo"
	"LineMongo/internals/tools"
)

var (
	MongoDBManager tools.InterfaceManger = tools.NewMongoDBManager(LineMongo.Config.MongoDBSetting)
	LineBotManager tools.InterfaceManger = tools.NewLineBotManager(LineMongo.Config.LineBotSetting)
)
