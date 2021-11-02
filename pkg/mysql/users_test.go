package mysql

import (
	"reflect"
	"testing"
	"time"

	"github.com/kguen/snippetbox/pkg/models"
)

func TestGet(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}
	tests := []struct {
		name      string
		userID    int
		wantUser  *models.User
		wantError error
	}{
		{
			name:   "Valid ID",
			userID: 1,
			wantUser: &models.User{
				ID:      1,
				Name:    "Khoa Nguyễn",
				Email:   "nanhkhoa460@gmail.com",
				Created: time.Date(2021, 11, 2, 14, 56, 21, 0, time.UTC),
			},
			wantError: nil,
		},
		{
			name:      "Zero ID",
			userID:    0,
			wantUser:  nil,
			wantError: models.ErrNoRecord,
		},
		{
			name:      "Non-existent ID",
			userID:    2,
			wantUser:  nil,
			wantError: models.ErrNoRecord,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := &UserModel{db}
			u, err := m.Get(tt.userID)

			if err != tt.wantError {
				t.Errorf("want %q; got %q", tt.wantError, err)
			}
			if !reflect.DeepEqual(u, tt.wantUser) {
				t.Errorf("want %v; got %v", tt.wantUser, u)
			}
		})
	}
}
