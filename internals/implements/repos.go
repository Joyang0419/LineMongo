package implements

import (
	"LineMongo/internals/repos/userMessageRepo"
)

var (
	UserMessageRepo userMessageRepo.InterfaceUserMessageRepo = userMessageRepo.NewMongoUserMessageRepo(MongoDBManager)
)
