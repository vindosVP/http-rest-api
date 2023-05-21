package apiserver

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"github.com/vindosVp/http-rest-api/internal/app/config"
	"github.com/vindosVp/http-rest-api/internal/app/logger"
	"github.com/vindosVp/http-rest-api/internal/app/store/sqlstore"
	"net/http"
)

func Start(conf *config.Config) error {
	logger.GetLogger().Info("Starting server...")
	db, err := newDB(conf.DB)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.GetLogger().Fatal(err)
		}
	}(db)
	store := sqlstore.New(db)
	sessionsStore := sessions.NewCookieStore([]byte(conf.SessionKey))
	s := newServer(store, sessionsStore)
	logger.GetLogger().Info(fmt.Sprintf("Server listening on %s", conf.Sever.BindAddr))
	return http.ListenAndServe(conf.Sever.BindAddr, s)
}

func newDB(dbConf config.DBConfig) (*sql.DB, error) {
	logger.GetLogger().Info("Connecting to database...")

	cStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.DBName)
	db, err := sql.Open("postgres", cStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	logger.GetLogger().Info("Connected successful...")

	return db, nil
}
