package apiserver

import (
	"golang-rest-api/internel/app/store"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config       *Config
	logger       *logrus.Logger
	router       *mux.Router
	storeFactori *store.StoreFactory
}

func New(c *Config) *APIServer {
	return &APIServer{
		config:       c,
		logger:       logrus.New(),
		router:       mux.NewRouter(),
		storeFactori: &store.StoreFactory{},
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return nil
	}

	s.configureRouter()

	if err := s.configureStoreFactory(); err != nil {
		return err
	}

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureStoreFactory() error {
	sf := store.New(s.config.Store)
	s.storeFactori = sf
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.hendleHello())
}

func (s *APIServer) hendleHello() http.HandlerFunc {
	//Local tape, variables and ...
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
