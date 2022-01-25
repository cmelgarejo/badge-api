package data

import (
	"context"
	"time"

	"github.com/cmelgarejo/badge-api/internal/middleware"
	"github.com/cmelgarejo/badge-api/pkg/points"
)

// PointRepository manages the operations with the database that
// correspond to the point model.
type PointRepository struct {
	Data *Data
}

// GetPoints returns all points for a given user.
func (pr *PointRepository) GetPoints(ctx context.Context, userID uint, start *time.Time, end *time.Time) ([]points.Point, error) {
	q := `SELECT id, points, metadata, assigned_at, created_at, updated_at FROM points
	WHERE user_id = $1 AND assigned_at BETWEEN $2 AND $3;`

	if start == nil || end == nil {
		now := time.Now()
		start = &time.Time{}
		end = &now
	}
	rows, err := pr.Data.DB.QueryContext(ctx, q, userID, start, end)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var thepoints []points.Point
	for rows.Next() {
		var p points.Point
		_ = rows.Scan(&p.ID, &p.Points, &p.Metadata, &p.AssignedAt, &p.CreatedAt, &p.UpdatedAt)
		thepoints = append(thepoints, p)
	}

	return thepoints, nil
}

// AssignPoints adds a new points to the user
func (pr *PointRepository) AssignPoints(ctx context.Context, p *points.Point) error {
	q := `INSERT INTO points (points, metadata, user_id, assigned_at) VALUES ($1, $2, $3, $4) RETURNING id;`

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, p.Points, p.Metadata, ctx.Value(middleware.UserIDKey), p.AssignedAt)

	err = row.Scan(&p.ID)
	if err != nil {
		return err
	}

	return nil
}
