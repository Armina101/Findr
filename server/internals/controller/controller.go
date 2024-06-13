package controller

import (
	"log/slog"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/thebravebyte/linker-app/internals"
)

// LinkerApp struct holds the database and logger instances for the LinkerApp service.
type FindrApp struct {
	DB     *mongo.Client
	Logger *slog.Logger
}

// Linker is a function that returns a new instance of the LinkerApp struct.
func NewFindr(db *mongo.Client, logger *slog.Logger) internals.FindrAppService {
	return &FindrApp{
		DB:     db,
		Logger: logger,
	}
}

// Todo: Endpoints implementation start here
func (lk *FindrApp) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		return
	}

}

func (lk *FindrApp) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		return
	}
}
func (lk *FindrApp) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		return
	}
}
