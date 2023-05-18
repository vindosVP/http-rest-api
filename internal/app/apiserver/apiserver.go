package apiserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

type ApiServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *ApiServer {
	return &ApiServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *ApiServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *ApiServer) configureRouter() {
	s.router.HandleFunc("/heartbeat", s.handleHeartbeat())
}

func (s *ApiServer) handleHeartbeat() http.HandlerFunc {
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

func (s *ApiServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()
	s.logger.Info(fmt.Sprintf("Server listening on %s", s.config.BindAddr))

	return http.ListenAndServe(s.config.BindAddr, s.router)
}
