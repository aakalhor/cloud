package usecase

import (
	"context"
	"github.com/google/uuid"
	"shop-pet/app/shop/adapter"
	"shop-pet/domain"
	"time"
)

type shopUsecase struct {
	a adapter.Adapter
	r domain.ShopRepository
}

func New(r domain.ShopRepository) domain.ShopUsecase {
	return &shopUsecase{
		r: r,
	}
}

func (s *shopUsecase) RegisterCategory(ctx context.Context, name string, desc string) (id uuid.UUID, err error) {
	id = uuid.New()
	return id, s.r.CreateCategory(ctx, name, desc, id)
}

func (s *shopUsecase) RegisterItem(ctx context.Context, categoryName string, name string, sex string, price int, count int) (id uuid.UUID, err error) {
	id = uuid.New()
	return id, s.r.CreateItem(ctx, categoryName, name, sex, price, count, id)
}

func (s *shopUsecase) AddToItem(ctx context.Context, Id uuid.UUID, count int) error {
	item, err := s.GetItem(ctx, Id)
	if err != nil {
		return err
	}

	newCount := item.Count + count
	return s.r.UpdateItemCount(ctx, Id, newCount)
}

func (s *shopUsecase) GetItem(ctx context.Context, Id uuid.UUID) (*domain.Item, error) {
	return s.r.GetItem(ctx, Id)
}

func (s *shopUsecase) RegisterOrder(ctx context.Context, id uuid.UUID, username string, productId uuid.UUID, date time.Time, count int) (uuid.UUID, error) {
	return s.r.CreateOrder(ctx, id, username, productId, date, count)
}

func (s *shopUsecase) CancelOrder(ctx context.Context, id uuid.UUID) error {
	return s.r.CancelOrder(ctx, id)
}

func (s *shopUsecase) GetOrderByUser(ctx context.Context, username string) (*[]domain.Order, error) {
	return s.r.GetOrderByUser(ctx, username)
}
