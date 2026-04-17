package station

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Afdhal-io/MRT-Schedules/modules/station/common/client"
)

type Service interface {
	GetAllStation() (response []StationResponse, err error)
}

type service struct {
	client *http.Client
}

type StationAPIResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Id   string `json:"nid"`
		Name string `json:"title"`
	} `json:"data"`
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *service) GetAllStation() (response []StationResponse, err error) {
	url := "https://mrt-jakarta-api-production.up.railway.app/v1/stations"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return nil, err
	}

	var result StationAPIResponse
	err = json.Unmarshal(byteResponse, &result)
	if err != nil {
		return nil, err
	}

	for _, item := range result.Data {
		response = append(response, StationResponse{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	return response, nil
}
