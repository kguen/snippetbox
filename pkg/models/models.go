package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching records found")
	ErrInvalidCredentials = errors.New("model: invalid login credentials")
	ErrDuplicateEmail     = errors.New("model: duplicate email")
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type User struct {
	ID      int
	Name    string
	Email   string
	Created time.Time
}
