package rest

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestNewLogService_WhenAddLog_WithValidString_Returns200(t *testing.T) {
	r := SetUpRouter()
	h, _ := NewLogsService()
	r.POST("/ingest", h.AddLogs)
	//deviceId := "TESTxx"
	log := RequestBody{Entry: []string{"YYARKx|2022-03-22T21:42:02.362Z|-8"}}

	jsonValue, _ := json.Marshal(log)
	req, _ := http.NewRequest("POST", "/ingest", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestNewLogService_WhenAddLog_WithInValidString_Returns500(t *testing.T) {
	r := SetUpRouter()
	h, _ := NewLogsService()
	r.POST("/ingest", h.AddLogs)
	log := RequestBody{Entry: []string{"YYARKx"}} //Bad log string

	jsonValue, _ := json.Marshal(log)
	req, _ := http.NewRequest("POST", "/ingest", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
