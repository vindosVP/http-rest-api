package sqlstore

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/vindosVp/http-rest-api/internal/app/store"
)

type Store struct {
	db             *sql.DB
	logger         *logrus.Logger
	UserRepository *UserRepository
}

func New(db *sql.DB, logger *logrus.Logger) *Store {
	return &Store{
		db:     db,
		logger: logger,
	}
}

func (s *Store) User() store.UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}
	s.UserRepository = &UserRepository{
		store: s,
	}
	return s.UserRepository
}
