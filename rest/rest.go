package rest

import "github.com/gin-gonic/gin"

func RunAPI(address string) error {
	//Get gin's default engine
	r := gin.Default()
	//Define a handler
	h, _ := NewLogsService()
	// Add log entries
	r.POST("/ingest", h.AddLogs)
	// Get logs by device id
	r.GET("/device/:deviceId", h.GetLogsByDeviceId)
	return r.Run(address)
}
