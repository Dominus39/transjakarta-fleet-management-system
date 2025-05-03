package api

import (
	"log"
	"net/http"
	"strconv"

	"TJ-project/internal/db"

	"github.com/gin-gonic/gin"
)

func GetLatestLocation(c *gin.Context) {
	id := c.Param("vehicle_id")
	loc, err := db.GetLatestLocation(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errr": "not found"})
		return
	}
	c.JSON(http.StatusOK, loc)
}

func GetHistory(c *gin.Context) {
	id := c.Param("vehicle_id")

	startInt, err1 := strconv.ParseInt(c.Query("start"), 10, 64)
	endInt, err2 := strconv.ParseInt(c.Query("end"), 10, 64)

	if err1 != nil || err2 != nil {
		log.Println("Start:", startInt, "End:", endInt)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start or end timestamp"})
		return
	}

	locs, err := db.GetLocationHistory(id, startInt, endInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch history"})
		return
	}

	c.JSON(http.StatusOK, locs)
}
