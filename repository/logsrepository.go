package repository

import (
	"backend-code-challenge-main/models"
)

type LogsRepository interface {
	GetLogsByDeviceId(deviceId string) ([]models.Log, error)
	AddLog(models.Log) (models.Log, error)
}
