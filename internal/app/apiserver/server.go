package apiserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/vindosVp/http-rest-api/internal/app/logger"
	"github.com/vindosVp/http-rest-api/internal/app/store"
	"io"
	"net/http"
	"time"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logger.GetLogger(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/heartbeat", s.handleHeartbeat()).Methods(http.MethodGet)
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods(http.MethodPost)
}

func (s *server) handleHeartbeat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		location, err := time.LoadLocation("Europe/Moscow")
		if err != nil {
			s.logger.Fatal(err)
		}
		timeNow := time.Now().UTC().In(location).Format("2006-01-02T15:04:05-0700")
		_, err = io.WriteString(w, fmt.Sprintf("[%v] server running", timeNow))
		if err != nil {
			s.logger.Error(err)
		}
	}
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
