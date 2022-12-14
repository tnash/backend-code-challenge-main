package models

import (
	"time"
)

type Log struct {
	EventId       uint       `gorm:"column:event_id; primaryKey; autoincrement"`
	DeviceId      string     `gorm:"column:device_id;not null;size:6" json:"deviceId" sql:"-"`
	TempFarenheit int64      `gorm:"column:temp_farenheit" json:"averageTemperature"`
	EventDate     *time.Time `gorm:"column:event_date;not null" json:"mostRecentLogDate"`
}

func (Log) TableName() string {
	return "logs"
}
