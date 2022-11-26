package implements

import "LineMongo/internals/services/lineService"

var (
	LineService lineService.InterfaceLineService = lineService.NewLineService(UserMessageRepo, LineBotManager)
)
