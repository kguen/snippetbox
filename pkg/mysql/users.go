package mysql

import (
	"database/sql"
	"log"
	"strings"

	"github.com/alexedwards/argon2id"
	"github.com/go-sql-driver/mysql"
	"github.com/kguen/snippetbox/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}
	stmt := /* sql */ `
		INSERT INTO users (name, email, hashed_password, created)
		VALUES (?, ?, ?, UTC_TIMESTAMP())`

	_, err = m.DB.Exec(stmt, name, email, hashedPassword)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			// email unique constraint error
			if mysqlErr.Number == 1062 && strings.Contains(mysqlErr.Message, "users_uc_email") {
				return models.ErrDuplicateEmail
			}
		}
		return err
	}
	return err
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	stmt := /* sql */ `
		SELECT id, hashed_password FROM users
		WHERE email = ?`

	var id int
	var hashedPassword string
	err := m.DB.
		QueryRow(stmt, email).
		Scan(&id, &hashedPassword)

	if err == sql.ErrNoRows {
		return 0, models.ErrInvalidCredentials
	}
	if err != nil {
		return 0, err
	}
	match, err := argon2id.ComparePasswordAndHash(password, hashedPassword)
	if err != nil {
		return 0, err
	}
	if !match {
		return 0, models.ErrInvalidCredentials
	}
	return id, nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
	stmt := /* sql */ `
		SELECT id, name, email, created FROM users
		WHERE id = ?`

	u := new(models.User)
	err := m.DB.
		QueryRow(stmt, id).
		Scan(&u.ID, &u.Name, &u.Email, &u.Created)

	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	}
	if err != nil {
		return nil, err
	}
	return u, nil
}
