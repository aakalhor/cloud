package repository

import (
	"github.com/google/uuid"
	"shop-pet/domain"
	"time"
)

type Category struct {
	Id   uuid.UUID `gorm:"primary_key;auto_increment;not_null"`
	Name string    `gorm:"primaryKey"`
	Desc string
}

type Item struct {
	Id         uuid.UUID `gorm:"primary_key;auto_increment;not_null"`
	CategoryId uuid.UUID
	Name       string
	Sex        string
	Price      int
	Count      int
}

type Order struct {
	OId       uuid.UUID `gorm:"primary_key;auto_increment;not_null"`
	Username  string
	ProductId uuid.UUID
	Count     int
	Date      time.Time
}

func (item *Item) toDomain() *domain.Item {
	return &domain.Item{
		CategoryName: item.CategoryId.String(),
		Name:         item.Name,
		Sex:          item.Sex,
		Price:        item.Price,
		Count:        item.Count,
		Id:           item.Id,
	}
}

func (category *Category) toDomain() *domain.Category {
	return &domain.Category{
		Name: category.Name,
		Desc: category.Desc,
		Id:   category.Id,
	}
}

func (order *Order) toDomain() *domain.Order {
	return &domain.Order{
		OId:       order.OId,
		Username:  order.Username,
		ProductId: order.ProductId,
		Count:     order.Count,
		Date:      order.Date,
	}
}
