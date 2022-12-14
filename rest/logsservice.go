package rest

import (
	"backend-code-challenge-main/models"
	"backend-code-challenge-main/repository"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type LogsServiceInterface interface {
	AddLogs(c *gin.Context)
	GetLogsByDeviceId(c *gin.Context)
}

type LogsService struct {
	db *repository.DBORM
}

type MockLogsService struct {
	db *repository.MockLogsRepository
}

type RequestBody struct {
	// json tag to de-serialize json body
	Entry []string `json:"data"`
}

type DeviceLogs struct {
	AverageTemperature float64    `json:"averageTemperature"`
	MostRecentLogDate  *time.Time `json:"mostRecentLogDate"`
	Logs               []Log      `json:"logs"`
}

type Log struct {
	LogDate     *time.Time `json:"logDate"`
	Temperature int64      `json:"temperature"`
	Humidity    float32    `json:"humidity"`
}

func NewLogsService() (LogsServiceInterface, error) {
	db, err := repository.NewORM("localhost", "pguser", "pgpassword", "code_challenge", "5432")
	if err != nil {
		return nil, err
	}
	return &LogsService{
		db: db,
	}, nil
}

func (h *LogsService) GetLogsByDeviceId(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	deviceId := c.Param("deviceId")
	logs, err := h.db.GetLogsByDeviceId(deviceId)
	deviceLogs := BuildDeviceLog(logs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Found %d logs\n", len(logs))
	c.JSON(http.StatusOK, deviceLogs)
}

func (h *LogsService) AddLogs(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}

	var body RequestBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.ParseRequestAndSaveLogs(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *LogsService) ParseRequestAndSaveLogs(body RequestBody) error {
	for _, deviceEntry := range body.Entry {
		tokens := strings.SplitN(deviceEntry, "|", 3)
		if len(tokens) == 3 {
			deviceId := tokens[0]
			eventDate, err := ExtractEntryDate(tokens[1])
			if err != nil {
				return err
			}
			tempF, err := ExtractTemperature(tokens[2])
			if err != nil {
				return err
			}
			logRecord := models.Log{EventId: 0, DeviceId: deviceId, TempFarenheit: tempF, EventDate: &eventDate}
			_, err = h.db.AddLog(logRecord)
		} else {
			return errors.New("Error: log entry did not match expected number of elements.")
		}
	}
	return nil
}

func ExtractEntryDate(dateString string) (time.Time, error) {
	eventDate, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		return eventDate, errors.New("Error: event date not of expected datetime value.")
	}
	return eventDate, err
}

func ExtractTemperature(tempString string) (int64, error) {
	tempC, err := strconv.ParseInt(tempString, 10, 64)
	if err != nil {
		return tempC, errors.New("Error: temperature value cannot be converted to integer.")
	}
	tempF := CelciusToFarenheit(tempC)
	return tempF, err
}

func CelciusToFarenheit(tempCelcius int64) int64 {
	return (2 * tempCelcius) + 30
}

func BuildDeviceLog(logs []models.Log) DeviceLogs {
	if len(logs) > 0 {
		var totalTemperature int64 = 0
		var averageTemperature float64
		mostRecentLogDate := time.Now().UTC()
		logRecords := make([]Log, len(logs))
		for index, logEntry := range logs {
			totalTemperature += logEntry.TempFarenheit
			if logEntry.EventDate.Before(mostRecentLogDate) {
				mostRecentLogDate = logEntry.EventDate.UTC()
			}
			logRecords[index] = Log{LogDate: logEntry.EventDate, Temperature: logEntry.TempFarenheit}
		}
		averageTemperature = float64(totalTemperature / int64(len(logs)))

		return DeviceLogs{AverageTemperature: averageTemperature, MostRecentLogDate: &mostRecentLogDate, Logs: logRecords}
	} else {
		return DeviceLogs{}
	}
}
