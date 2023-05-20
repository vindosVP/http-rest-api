package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/vindosVp/http-rest-api/internal/app/config"
)

type Store struct {
	config         *config.Config
	db             *sql.DB
	logger         *logrus.Logger
	UserRepository *UserRepository
}

func New(config *config.Config, logger *logrus.Logger) *Store {
	return &Store{
		config: config,
		logger: logger,
	}
}

func (s *Store) User() *UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}
	s.UserRepository = &UserRepository{
		store: s,
	}
	return s.UserRepository
}

func (s *Store) Open() error {
	s.logger.Info("Connecting to database...")
	cStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		s.config.DB.Host, s.config.DB.Port, s.config.DB.User, s.config.DB.Password, s.config.DB.DBName)
	db, err := sql.Open("postgres", cStr)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	s.logger.Info("Connected successful..")
	return nil
}

func (s *Store) Close() error {
	err := s.db.Close()
	if err != nil {
		return err
	}
	return nil
}
