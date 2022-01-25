package data

import (
	"context"
	"time"

	"github.com/cmelgarejo/badge-api/pkg/badge"
)

// BadgeRepository manages the operations with the database that
// correspond to the badge model.
type BadgeRepository struct {
	Data *Data
}

// GetAll returns all badges.
func (pr *BadgeRepository) GetAll(ctx context.Context) ([]badge.Badge, error) {
	q := `
	SELECT id, body, user_id, created_at, updated_at
		FROM badges;
	`

	rows, err := pr.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var badges []badge.Badge
	for rows.Next() {
		var p badge.Badge
		_ = rows.Scan(&p.ID, &p.Body, &p.UserID, &p.CreatedAt, &p.UpdatedAt)
		badges = append(badges, p)
	}

	return badges, nil
}

// GetOne returns one badge by id.
func (pr *BadgeRepository) GetOne(ctx context.Context, id uint) (badge.Badge, error) {
	q := `
	SELECT id, body, user_id, created_at, updated_at
		FROM badges WHERE id = $1;
	`

	row := pr.Data.DB.QueryRowContext(ctx, q, id)

	var p badge.Badge
	err := row.Scan(&p.ID, &p.Body, &p.UserID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return badge.Badge{}, err
	}

	return p, nil
}

// GetByUser returns all user badges.
func (pr *BadgeRepository) GetByUser(ctx context.Context, userID uint) ([]badge.Badge, error) {
	q := `
	SELECT id, body, user_id, created_at, updated_at
		FROM badges
		WHERE user_id = $1;
	`

	rows, err := pr.Data.DB.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var badges []badge.Badge
	for rows.Next() {
		var p badge.Badge
		_ = rows.Scan(&p.ID, &p.Body, &p.UserID, &p.CreatedAt, &p.UpdatedAt)
		badges = append(badges, p)
	}

	return badges, nil
}

// Create adds a new badge.
func (pr *BadgeRepository) Create(ctx context.Context, p *badge.Badge) error {
	q := `
	INSERT INTO badges (body, user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, p.Body, p.UserID, time.Now(), time.Now())

	err = row.Scan(&p.ID)
	if err != nil {
		return err
	}

	return nil
}

// Update updates a badge by id.
func (pr *BadgeRepository) Update(ctx context.Context, id uint, p badge.Badge) error {
	q := `
	UPDATE badges set body=$1, updated_at=$2
		WHERE id=$3;
	`

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, p.Body, time.Now(), id,
	)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a badge by id.
func (pr *BadgeRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM badges WHERE id=$1;`

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
