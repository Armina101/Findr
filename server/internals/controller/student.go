package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/thebravebyte/findr/internals"
	"github.com/thebravebyte/findr/pkg/enc"
	"github.com/thebravebyte/findr/pkg/req"
	"github.com/thebravebyte/findr/pkg/res"
)

// Todo: Endpoints implementation start here

func (fr *FindrApp) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var std internals.Students
		s := req.ReadAndValidate(w, r, &std)
		if r, ok := s.(*internals.Students); ok {
			std = *r
		}

		// encrypt the user password using argon21d package
		password, err := enc.CreateHash(std.Password)
		if err != nil {
			fr.Logger.Error(err.Error())
			_ = res.Writer(w, http.StatusBadRequest, map[string]interface{}{
				"error":   err.Error(),
				"message": "unable to secure your credentials",
			})
			return
		}

		std.Password = password
		std.UpdatedAt, std.CreatedAt = time.Now(), time.Now()

		// database queries to add new user
		userId, msg, err := fr.DB.AddNew(user)
		if err != nil {
			_ = res.Writer(wr, http.StatusInternalServerError, map[string]interface{}{
				"error":   err.Error(),
				"message": fmt.Sprintf("unable to add this user - %v", userId),
			})
			return

		}

		_ = res.Writer(w, http.StatusOK, map[string]interface{}{
			"message": msg,
			"user_id": userId,
		})
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
