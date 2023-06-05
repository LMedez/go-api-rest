package entity

import (
	"time"
)

// Album represents an album record.
type Album struct {
	ID        string    `firestore:"id"`
	Name      string    `firestore:"name"`
	CreatedAt time.Time `firestore:"created_at"`
	UpdatedAt time.Time `firestore:"updated_at"`
}
