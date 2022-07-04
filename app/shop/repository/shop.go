package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"shop-pet/domain"
	"time"
)

type shopRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ShopRepository {
	return &shopRepository{
		db: db,
	}
}

func (s *shopRepository) CreateCategory(ctx context.Context, name string, desc string, id uuid.UUID) error {
	category := Category{
		Id:   id,
		Name: name,
		Desc: desc,
	}
	err := s.db.WithContext(ctx).Model(&category).Create(&category).Error

	return err
}

func (s *shopRepository) CreateItem(ctx context.Context, categoryName string, name string, sex string, price int, count int, id uuid.UUID) error {
	cat, err := s.GetCategoryByName(ctx, categoryName)
	fmt.Println(cat, "kk ", err)
	if err != nil {
		return err
	}
	if cat == nil {
		return errors.New("category does not exist")
	}
	item := Item{
		Id:         id,
		CategoryId: cat.Id,
		Name:       name,
		Sex:        sex,
		Price:      price,
		Count:      count,
	}
	err = s.db.WithContext(ctx).Model(&item).Create(&item).Error
	return err
}

func (s *shopRepository) UpdateItemCount(ctx context.Context, id uuid.UUID, newCount int) error {

	var item Item
	err := s.db.WithContext(ctx).Model(&item).Where("id = ?", id).Update("count", newCount).Error

	return err
}

func (s *shopRepository) GetItem(ctx context.Context, Id uuid.UUID) (*domain.Item, error) {
	var item Item
	err := s.db.WithContext(ctx).Model(&item).Where("id = ?", Id).Find(&item).Error
	if err != nil {
		return nil, err
	}
	dto := item.toDomain()

	return dto, nil
}

func (s *shopRepository) GetCategoryByName(ctx context.Context, name string) (*domain.Category, error) {
	var category Category
	err := s.db.WithContext(ctx).Model(&category).Where("name = ?", name).Find(&category).Error
	if err != nil {
		return nil, err
	}
	dto := category.toDomain()

	return dto, nil
}

func (s *shopRepository) CreateOrder(ctx context.Context, id uuid.UUID, username string, productId uuid.UUID, date time.Time, count int) (uuid.UUID, error) {
	order := Order{
		OId:       id,
		Username:  username,
		ProductId: productId,
		Count:     count,
		Date:      date,
	}
	err := s.db.WithContext(ctx).Model(&order).Create(&order).Error
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (s *shopRepository) CancelOrder(ctx context.Context, id uuid.UUID) error {
	var order Order
	err := s.db.WithContext(ctx).Model(&order).Where("id = ?", id).Delete(&order).Error
	return err
}

func (s *shopRepository) GetOrderByUser(ctx context.Context, username string) (*[]domain.Order, error) {
	var daos []Order
	var order Order
	err := s.db.WithContext(ctx).Model(&order).Where("username = ?", username).Find(&daos).Error
	if err != nil {
		return nil, err
	}
	dtos := make([]domain.Order, len(daos))
	for idx, val := range daos {
		dto := val.toDomain()
		dtos[idx] = *dto
	}
	return &dtos, nil
}
