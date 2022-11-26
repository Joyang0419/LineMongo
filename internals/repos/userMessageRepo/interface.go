package userMessageRepo

import "LineMongo/internals/models"

type InterfaceUserMessageRepo interface {
	Get() []models.UserMessageDTO
	Create(dto models.UserMessageDTO) string
}
