package controller

import (
	"log/slog"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/thebravebyte/findr/internals"
	"github.com/thebravebyte/findr/internals/mongodb"
	"github.com/thebravebyte/findr/pkg/res"
)

// LinkerApp struct holds the database and logger instances for the LinkerApp service.
type FindrApp struct {
	DB     mongodb.FindrStore
	Logger *slog.Logger
}

// Linker is a function that returns a new instance of the LinkerApp struct.
func NewFindr(db *mongo.Client, logger *slog.Logger) internals.FindrAppService {
	return &FindrApp{
		DB:     mongodb.NewFindrDB(db),
		Logger: logger,
	}
}

func (fr *FindrApp) FindrHomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res.Writer(w, http.StatusOK, map[string]interface{}{"message": "Welcome to the FindrApp service"})
	}
}
