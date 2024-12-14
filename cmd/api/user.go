package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/codepnw/social/internal/store"
	"github.com/go-chi/chi/v5"
)

type userKey string

const userKeyCtx userKey = "user"

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r)

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.errorInternalServer(w, r, err)
	}
}

// User Follower
type FollowUser struct {
	UserID int64 `json:"user_id"`
}

func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	followerUser := getUserFromContext(r)

	// TODO: revert back to auth userID
	var payload FollowUser
	if err := readJSON(w, r, &payload); err != nil {
		app.errorBadRequest(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Followers.Follow(ctx, followerUser.ID, payload.UserID); err != nil {
		switch err {
		case store.ErrConflict:
			app.errorConflict(w, r, err)
			return
		default:
			app.errorInternalServer(w, r, err)
			return
		}
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.errorInternalServer(w, r, err)
	}
}

func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	unfollowedUser := getUserFromContext(r)

	// TODO: revert back to auth userID
	var payload FollowUser
	if err := readJSON(w, r, &payload); err != nil {
		app.errorBadRequest(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Followers.Unfollow(ctx, unfollowedUser.ID, payload.UserID); err != nil {
		app.errorInternalServer(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.errorInternalServer(w, r, err)
	}
}

func (app *application) userContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		ctx = context.WithValue(ctx, userKeyCtx, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserFromContext(r *http.Request) *store.User {
	user, _ := r.Context().Value(userKeyCtx).(*store.User)
	return user
}
