package teststore

import (
	"github.com/google/uuid"
	"github.com/vindosVp/http-rest-api/internal/app/model"
	"github.com/vindosVp/http-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func (ur *UserRepository) Create(u *model.User) error {

	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	newuuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	u.UUID = newuuid

	ur.users[u.Email] = u

	return nil
}

func (ur *UserRepository) FindByEmail(email string) (*model.User, error) {

	u, ok := ur.users[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}
