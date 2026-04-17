package station

import (
	"net/http"

	"github.com/Afdhal-io/MRT-Schedules/modules/station/common/client/response"
	"github.com/gin-gonic/gin"
)

func Initiate(router *gin.RouterGroup) {
	stationService := NewService()

	station := router.Group("/stations")

	station.GET("", func(c *gin.Context) {
		GetAllStation(c, stationService)
	})
}

func GetAllStation(c *gin.Context, service Service) {
	datas, err := service.GetAllStation()
	if err != nil {
		c.JSON(
			http.StatusBadGateway,
			response.APIResponse{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		response.APIResponse{
			Success: true,
			Message: "Successfully get all station",
			Data:    datas,
		},
	)
}
