package userMessageRepo

import "LineMongo/internals/models"

type Filter struct {
	Key   string
	Value any
}

type InterfaceUserMessageRepo interface {
	Get(...Filter) []map[string]any
	Create(dto models.UserMessageDTO) string
}
