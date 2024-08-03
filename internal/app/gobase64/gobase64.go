package gobase64

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	services "github.com/Nortvest/goBase64/internal/services/base64"
)

type server struct {
	Config *Config
	Logger *slog.Logger
	Router *mux.Router
}

func New(config *Config) *server {
	return &server{
		Config: config,
		Logger: slog.New(slog.NewTextHandler(os.Stderr, nil)),
		Router: mux.NewRouter(),
	}
}

func (s *server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.Logger.Info("Server started!")

	return http.ListenAndServe(s.Config.BindAdrr, s.Router)
}

func (s *server) configureLogger() error {
	level, err := services.ParseLogLevel(s.Config.LogLevel)
	if err != nil {
		return err
	}

	s.Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))

	return nil
}

func (s *server) configureRouter() {
	s.Router.HandleFunc("/api/v1/encode", s.HandlerEncode()).Methods("GET")
	s.Router.HandleFunc("/api/v1/decode", s.HandlerDecode()).Methods("GET")
}

func (s *server) HandlerEncode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.HandlerCodeHelp(w, r, services.EncodeBase64)
	}
}

func (s *server) HandlerDecode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.HandlerCodeHelp(w, r, services.DecodeBase64)
	}
}

func (s *server) HandlerCodeHelp(w http.ResponseWriter, r *http.Request, coder func(text string) (string, error)) {
	query := r.URL.Query()
	text, ok := query["text"]

	if !ok {
		s.error(w, r, http.StatusBadRequest, errors.New("text is required"))
		return
	}

	codeText, err := coder(text[0])
	if err != nil {
		s.error(w, r, http.StatusBadRequest, errors.New("incorrect text"))
		return
	}
	s.respond(w, r, http.StatusOK, map[string]string{"text": text[0], "result": codeText})
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
