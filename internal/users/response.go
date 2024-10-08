package users

import "gis/internal/models"

// EmptyResponse adalah struktur untuk respons kosong
type EmptyResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Users []models.Users `json:"users"`
	} `json:"data"`
}
