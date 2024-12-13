package main

import (
	"net/http"
	"strconv"

	"github.com/codepnw/social/internal/store"
	"github.com/go-chi/chi/v5"
)

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
	if err != nil {
		app.errorBadRequest(w, r, err)
		return
	}

	ctx := r.Context()

	user, err := app.store.Users.GetByID(ctx, userID)
	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.errorNotFound(w, r, err)
			return
		default:
			app.errorInternalServer(w, r, err)
			return
		}
	}

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.errorInternalServer(w, r, err)
	}
}