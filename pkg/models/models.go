package models

import (
	"errors"
	"time"
)

// ErrNoRecord is error for record not found.
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet is created by user that has expire timestamp.
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
