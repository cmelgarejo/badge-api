package points

import (
	"context"
	"time"
)

// Repository handle the CRUD operations with Badges.
type Repository interface {
	GetPoints(ctx context.Context, userID uint, start *time.Time, end *time.Time) ([]Point, error)
	AssignPoints(ctx context.Context, p *Point) error
}
