package adapter

import (
	"context"
	"github.com/google/uuid"
	"order-pet/domain"
	"time"
)

type Adapter interface {
	CreateOrder(ctx context.Context, username string, orderId uuid.UUID, productId uuid.UUID, count int, date time.Time) error
	CancelOrder(ctx context.Context, username string, orderId uuid.UUID) (returnAmount int, err error)
	GetOrders(ctx context.Context, username string) (*[]domain.Order, error)
}
