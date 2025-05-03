package api

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/vehicles/:vehicle_id/location", GetLatestLocation)
	r.GET("/vehicles/:vehicle_id/history", GetHistory)

	return r
}
