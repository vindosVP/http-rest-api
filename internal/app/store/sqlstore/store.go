package sqlstore

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/vindosVp/http-rest-api/internal/app/store"
)

type Store struct {
	db             *sql.DB
	UserRepository *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
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
