package api

import (
	"github.com/Kratos40-sba/data-service/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LastNMeasurementHandler(db *database.Connection) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// todo
	}
}
func LastMeasurementSinceT(db *database.Connection) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// todo
	}
}
func ExampleHandler(db *database.Connection) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		err := ctx.ShouldBindJSON(db.ExampleInflux())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		//ctx.JSON(http.StatusOK, db.ExampleInflux())
	}
}
