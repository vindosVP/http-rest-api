package apiserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/vindosVp/http-rest-api/internal/app/logger"
	"github.com/vindosVp/http-rest-api/internal/app/model"
	"github.com/vindosVp/http-rest-api/internal/app/store"
	"io"
	"net/http"
	"time"
)

type server struct {
	router        *mux.Router
	store         store.Store
	sessionsStore sessions.Store
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
)

const (
	sessionName = "session"
)

func newServer(store store.Store, sessionsStore sessions.Store) *server {
	s := &server{
		router:        mux.NewRouter(),
		store:         store,
		sessionsStore: sessionsStore,
	}

	s.configureRouter()

	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/heartbeat", s.handleHeartbeat()).Methods(http.MethodGet)
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods(http.MethodPost)
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods(http.MethodPost)
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}

func (s *server) handleHeartbeat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		location, err := time.LoadLocation("Europe/Moscow")
		if err != nil {
			logger.GetLogger().Fatal(err)
		}
		timeNow := time.Now().UTC().In(location).Format("2006-01-02T15:04:05-0700")
		_, err = io.WriteString(w, fmt.Sprintf("[%v] server running", timeNow))
		if err != nil {
			logger.GetLogger().Error(err)
		}
	}
}

func (s *server) handleUsersCreate() http.HandlerFunc {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)

	}
}

func (s *server) handleSessionsCreate() http.HandlerFunc {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePwd(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		session, err := s.sessionsStore.Get(r, sessionName)

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_uuid"] = u.UUID.String()
		if err := s.sessionsStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}
