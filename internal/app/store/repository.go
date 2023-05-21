package store

import (
	"github.com/google/uuid"
	"github.com/vindosVp/http-rest-api/internal/app/model"
)

type UserRepository interface {
	Create(*model.User) error
	Find(uuid uuid.UUID) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
