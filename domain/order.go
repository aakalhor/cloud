package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Order struct {
	OId       string
	Username  string
	ProductId string
	Count     int
	Date      time.Time
}

type OrderUsecase interface {
	RegisterOrder(ctx context.Context, username string, productId uuid.UUID, date time.Time, count int) (*Order, error)
	CancelOrder(ctx context.Context, username string, orderId uuid.UUID) (int, error)
	GetOrdersByUsername(ctx context.Context, username string) (*[]Order, error)
}
