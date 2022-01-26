package points

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type JSONB map[interface{}]interface{}

// Points for users.
type Point struct {
	ID         uint      `json:"id,omitempty"`
	UserID     uint      `json:"-"`
	Points     float32   `json:"points,omitempty"`
	Metadata   *string   `json:"metadata,omitempty"`
	Level      int       `json:"level,omitempty"`
	AssignedAt time.Time `json:"assigned_at,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

// Make the JSONB struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Make the JSONB struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (a *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
