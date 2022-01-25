package data

import (
	"context"

	"github.com/cmelgarejo/badge-api/pkg/badge"
)

// BadgeRepository manages the operations with the database that
// correspond to the badge model.
type BadgeRepository struct {
	Data *Data
}

/**
`/v1/badge/org_id`

GET - Get all badges available to an organization
POST - params (name, image) Create a badge

`/v1/badge/user/_id`

GET - Get badges users received with (id, name, and image)
POST - params (badge id) - Add a badge to the user
*/

// Create adds a new badge.
func (br *BadgeRepository) Create(ctx context.Context, orgID uint, b *badge.Badge) error {
	q := `INSERT INTO badges (name, image, org_id) VALUES ($1, $2, $3) RETURNING id;`

	stmt, err := br.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, b.Name, b.Image, orgID)

	err = row.Scan(&b.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetByOrg returns all orgs badges.
func (br *BadgeRepository) GetByOrg(ctx context.Context, orgID uint) ([]badge.Badge, error) {
	q := `SELECT id, name, image, created_at, updated_at FROM badges WHERE org_id = $1;`

	rows, err := br.Data.DB.QueryContext(ctx, q, orgID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var badges []badge.Badge
	for rows.Next() {
		var p badge.Badge
		_ = rows.Scan(&p.ID, &p.Image, &p.Name, &p.CreatedAt, &p.UpdatedAt)
		badges = append(badges, p)
	}

	return badges, nil
}

// GetByUser returns all user badges.
func (br *BadgeRepository) GetByUser(ctx context.Context, userID uint) ([]badge.Badge, error) {
	q := `SELECT b.id, b.name, b.image, b.created_at, b.updated_at
		FROM badges b
		JOIN user_badges ON b.id = user_badges.badge_id
		WHERE user_id = $1;`

	rows, err := br.Data.DB.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var badges []badge.Badge
	for rows.Next() {
		var p badge.Badge
		_ = rows.Scan(&p.ID, &p.Image, &p.Name, &p.CreatedAt, &p.UpdatedAt)
		badges = append(badges, p)
	}

	return badges, nil
}

// AssignBadge assign a badge to a given user
func (br *BadgeRepository) AssignBadge(ctx context.Context, userID uint, badgeID uint) error {
	q := `INSERT INTO user_badges (badge_id, user_id) VALUES ($1, $2);`

	stmt, err := br.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, userID, badgeID)

	err = row.Err()
	if err != nil {
		return err
	}

	return nil
}
