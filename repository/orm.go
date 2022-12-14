package repository

import (
	"backend-code-challenge-main/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(host, user, password, dbname, port string) (*DBORM, error) {
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port +
		" sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return &DBORM{
		DB: db,
	}, err
}

func (db *DBORM) GetLogsByDeviceId(deviceId string) (logs []models.Log, err error) {
	return logs, db.Where(&models.Log{DeviceId: deviceId}).Order("event_date desc").Find(&logs).Error
}

func (db *DBORM) AddLog(log models.Log) (models.Log, error) {
	err := db.Create(&log).Error
	return log, err
}
