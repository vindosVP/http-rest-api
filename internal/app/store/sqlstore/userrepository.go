package sqlstore

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/vindosVp/http-rest-api/internal/app/model"
	"github.com/vindosVp/http-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {

	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow("INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING uuid", u.Email, u.EncryptedPassword).Scan(&u.UUID)
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {

	u := &model.User{}

	if err := r.store.db.QueryRow("SELECT uuid, email, encrypted_password FROM users WHERE email = $1", email).Scan(&u.UUID, &u.Email, &u.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) Find(uuid uuid.UUID) (*model.User, error) {

	u := &model.User{}

	if err := r.store.db.QueryRow("SELECT uuid, email, encrypted_password FROM users WHERE uuid = $1", uuid).Scan(&u.UUID, &u.Email, &u.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}
