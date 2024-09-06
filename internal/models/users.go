package models

// Users adalah model untuk tabel users
type Users struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
