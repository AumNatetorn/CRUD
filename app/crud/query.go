package crud

import (
	"CRUD/app"
	"context"
	"errors"

	"gorm.io/gorm"
)

type customerStore struct {
	db *gorm.DB
}

func NewCustomerStore(db gorm.DB) *customerStore {
	return &customerStore{db: &db}
}

func (c *customerStore) InsertOne(ctx context.Context, req Request) error {
	if err := c.db.WithContext(ctx).Create(&req).Error; err != nil {
		return err
	}

	return nil
}

func (c *customerStore) UpdateOne(ctx context.Context, req Request) error {
	var existingRecord Request
	if err := c.db.WithContext(ctx).First(&existingRecord, req.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return app.ErrRecordNotFound
		}
		return err
	}

	if err := c.db.WithContext(ctx).Model(&Request{}).Where("id = ?", req.ID).Updates(req).Error; err != nil {
		return err
	}

	return nil
}

func (c *customerStore) DeleteOne(ctx context.Context, id int) error {
	if err := c.db.WithContext(ctx).Where("id = ?", id).Delete(&Request{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *customerStore) FindOne(ctx context.Context, id int) (ent *Request, err error) {
	if err := r.db.WithContext(ctx).Model(ent).Where("id = ?", id).First(&ent).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, app.ErrRecordNotFound
		}
		return nil, err
	}

	return ent, nil
}
