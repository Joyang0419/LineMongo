package implements

import "LineMongo/internals/controllers/lineController"

var (
	LineController lineController.InterfaceLineController = lineController.NewV1LineController(LineService)
)
