package badge

import "context"

// Repository handle the CRUD operations with Badges.
type Repository interface {
	GetByOrg(ctx context.Context, orgID uint) ([]Badge, error)
	GetByUser(ctx context.Context, userID uint) ([]Badge, error)
	Create(ctx context.Context, orgID uint, b *Badge) error
	AssignBadge(ctx context.Context, userID uint, badgeID uint) error
}
