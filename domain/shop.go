package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Category struct {
	Name string
	Desc string
	Id   uuid.UUID
}

type Item struct {
	CategoryName string
	Name         string
	Sex          string
	Price        int
	Count        int
	Id           uuid.UUID
}

type Order struct {
	OId       uuid.UUID
	Username  string
	ProductId uuid.UUID
	Count     int
	Date      time.Time
}

type ShopUsecase interface {
	RegisterCategory(ctx context.Context, name string, desc string) (id uuid.UUID, err error)
	RegisterItem(ctx context.Context, categoryName string, name string, sex string, price int, count int) (id uuid.UUID, err error)
	AddToItem(ctx context.Context, Id uuid.UUID, count int) error
	GetItem(ctx context.Context, Id uuid.UUID) (*Item, error)
	RegisterOrder(ctx context.Context, id uuid.UUID, username string, productId uuid.UUID, date time.Time, count int) (uuid.UUID, error)
	CancelOrder(ctx context.Context, id uuid.UUID) error
	GetOrderByUser(ctx context.Context, username string) (*[]Order, error)
}

type ShopRepository interface {
	CreateCategory(ctx context.Context, name string, desc string, id uuid.UUID) error
	CreateItem(ctx context.Context, categoryName string, name string, sex string, price int, count int, id uuid.UUID) error
	UpdateItemCount(ctx context.Context, id uuid.UUID, newCount int) error
	GetItem(ctx context.Context, Id uuid.UUID) (*Item, error)
	GetCategoryByName(ctx context.Context, name string) (*Category, error)
	CreateOrder(ctx context.Context, id uuid.UUID, username string, productId uuid.UUID, date time.Time, count int) (uuid.UUID, error)
	CancelOrder(ctx context.Context, id uuid.UUID) error
	GetOrderByUser(ctx context.Context, username string) (*[]Order, error)
}
