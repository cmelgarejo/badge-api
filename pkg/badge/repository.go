package badge

import "context"

// Repository handle the CRUD operations with Badges.
type Repository interface {
	GetAll(ctx context.Context) ([]Badge, error)
	GetOne(ctx context.Context, id uint) (Badge, error)
	GetByUser(ctx context.Context, userID uint) ([]Badge, error)
	Create(ctx context.Context, badge *Badge) error
	Update(ctx context.Context, id uint, badge Badge) error
	Delete(ctx context.Context, id uint) error
}
