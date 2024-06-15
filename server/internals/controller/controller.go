package controller

import (
	"log/slog"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/thebravebyte/findr/internals"
	"github.com/thebravebyte/findr/pkg/req"
	"github.com/thebravebyte/findr/pkg/res"
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

func (fr *FindrApp) GetDB() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res.Writer(w, http.StatusOK, map[string]interface{}{"message": "Welcome to the FindrApp service"})
	}
}

// Todo: Endpoints implementation start here
func (fr *FindrApp) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var std internals.Students
		req.ReadAndValidate(w, r, nil)
	}

}

func (fr *FindrApp) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		return
	}
}
func (fr *FindrApp) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		return
	}
}
