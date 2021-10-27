package apiserver

import (
	"io"
	"net/http"

	"github.com/Muhammad-D/http_rest_api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//APIServer...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

//New...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//Start...
func (s *APIServer) Start() error {

	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Logger successfully connected")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

//####################################   A Configure Section   ####################################

//configureLogger...
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

//configureRouter...
func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

//configureStore...
func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)

	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

//####################################   A Handlers Section   ####################################

//Handlers...
func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello this is IT")
	}
}
