package album

import (
	"gis/internal/models"
)

type EmptyResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Albums []models.Album `json:"albums"`
	} `json:"data"`
}
