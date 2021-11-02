package mock

import (
	"time"

	"github.com/kguen/snippetbox/pkg/models"
)

type UserModel struct{}

var mockUser = &models.User{
	ID:      1,
	Name:    "Khoa",
	Email:   "khoa@example.com",
	Created: time.Now(),
}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "khoa@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	switch email {
	case "khoa@example.com":
		return 1, nil
	default:
		return 0, models.ErrInvalidCredentials
	}
}

func (m *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return nil, models.ErrNoRecord
	}
}
