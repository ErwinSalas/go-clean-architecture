package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHttpHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
	accounts   *AccountHandler
}

func NewAPIServer(listenAddr string, accountHandler *AccountHandler) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		accounts:   accountHandler,
	}
}

func (s *APIServer) Run() {
	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Use(middleware.RequestID)
		r.Use(middleware.Recoverer, middleware.URLFormat)
		r.Use(render.SetContentType(render.ContentTypeJSON))
		r.Use(cors.Handler(cors.Options{
			AllowOriginFunc:  allowOriginFunc,
			AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}))

		r.Get("/account", makeHttpHandler(s.accounts.handleGetAccount))
	})
	certFile := "server.crt"
	keyFile := "server.key"

	// Start the HTTPS server
	err := http.ListenAndServeTLS(s.listenAddr, certFile, keyFile, router)
	fmt.Print(err.Error())
}

func allowOriginFunc(r *http.Request, origin string) bool {
	// TODO: Replace this with slices.Contains after upgrading Go version to 1.18 or newer.
	allowedOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin || allowedOrigin == "*" {
			return true
		}
	}
	return false
}
