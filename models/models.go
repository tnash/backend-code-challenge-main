package models

import (
	"time"
)

type Log struct {
	EventId       uint       `gorm:"column:event_id; primaryKey; autoincrement"`
	DeviceId      string     `gorm:"column:device_id;not null;size:6" json:"deviceId"`
	TempFarenheit int64      `gorm:"column:temp_farenheit"`
	EventDate     *time.Time `gorm:"column:event_date;not null" json:"logDate"`
}

func (Log) TableName() string {
	return "logs"
}
