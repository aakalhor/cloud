package usecase

import (
	"context"
	"github.com/google/uuid"
	"order-pet/app/order/adapter"
	"order-pet/domain"
	"time"
)

type orderUsecase struct {
	a adapter.Adapter
}

func New(a adapter.Adapter) domain.OrderUsecase {
	return &orderUsecase{
		a: a,
	}
}

func (o *orderUsecase) RegisterOrder(ctx context.Context, username string, productId uuid.UUID, date time.Time, count int) (*domain.Order, error) {
	orderId := uuid.New()
	err := o.a.CreateOrder(ctx, username, orderId, productId, count, date)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (o *orderUsecase) CancelOrder(ctx context.Context, username string, orderId uuid.UUID) (int, error) {
	return o.a.CancelOrder(ctx, username, orderId)
}

func (o *orderUsecase) GetOrdersByUsername(ctx context.Context, username string) (*[]domain.Order, error) {
	return o.a.GetOrders(ctx, username)
}
