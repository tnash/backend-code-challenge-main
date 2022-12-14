package repository

import (
	"backend-code-challenge-main/models"
	"encoding/json"
	"fmt"
)

type MockLogsRepository struct {
	err  error
	logs []models.Log
}

func NewMockLogsRepository(logs []models.Log) *MockLogsRepository {
	return &MockLogsRepository{
		logs: logs,
	}
}

func NewMockLogsRepositoryWithData() *MockLogsRepository {
	LOGS := `[
        {
			"EventId": 21
            "EventDate": "2022-03-22 21:42:02.362000 +00:00",
            "TempFarenheit": 14,
            "DeviceId": "YYARKx"
        },
        {
			"EventId": 22
            "EventDate": "2022-03-22 21:42:04.372000 +00:00",
            "TempFarenheit": 28,
            "DeviceId": "YYARKx"
        },
        {
			"EventId": 23
            "EventDate": "2022-03-22 21:42:50.572000 +00:00",
            "TempFarenheit": 44,
            "DeviceId": "YYARKx"
        },
        {
			"EventId": 24
            "EventDate": "2022-03-23 21:42:02.362000 +00:00",
            "TempFarenheit": -6,
            "DeviceId": "XXARKx"
        },
        {
			"EventId": 25
            "EventDate": "2022-03-22 21:42:04.362000 +00:00",
            "TempFarenheit": 10,
            "DeviceId": "XXARKx"
        },
        {
			"EventId": 26
            "EventDate": "2022-03-20 21:42:50.572000 +00:00",
            "TempFarenheit": 48,
            "DeviceId": "XXARKx"
        }

    ]`
	var logs []models.Log
	err := json.Unmarshal([]byte(LOGS), &logs)
	if err != nil {
		fmt.Println(err)
	}
	return NewMockLogsRepository(logs)
}

func (mock *MockLogsRepository) GetMockLogData() []models.Log {
	return mock.logs
}

func (mock *MockLogsRepository) GetLogsByDeviceId(deviceId string) ([]models.Log, error) {
	//Should we return an error?
	if mock.err != nil {
		return nil, mock.err
	}
	//build logs array matching deviceId
	var logs = []models.Log{}
	for _, log := range mock.logs {
		if log.DeviceId == deviceId {
			logs = append(logs, log)
		}
	}
	return logs, nil
}

func (mock *MockLogsRepository) AddLog(log models.Log) (models.Log, error) {
	if mock.err != nil {
		return models.Log{}, mock.err
	}
	mock.logs = append(mock.logs, log)
	return log, nil
}

func (mock *MockLogsRepository) SetError(err error) {
	mock.err = err
}
