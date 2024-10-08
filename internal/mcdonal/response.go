package mcdonal

import (
	"gis/internal/models"
)

// EmptyResponse adalah struktur untuk respons kosong
type EmptyResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		McDonalds []models.Mcdonald `json:"mcdonalds"`
	} `json:"data"`
}
