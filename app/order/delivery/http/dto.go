package http

import "time"

type RegisterOrder struct {
	ProductId string    `json:"productId"`
	Count     int       `json:"count"`
	Date      time.Time `json:"date"`
}
