package models

// Mcdonald adalah model untuk tabel mcdonald
type Mcdonald struct {
    ID        int     `json:"id"`
    Name      string  `json:"name"`
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
}
