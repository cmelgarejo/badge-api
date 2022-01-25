package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cmelgarejo/badge-api/internal/middleware"
	"github.com/cmelgarejo/badge-api/pkg/badge"
	"github.com/cmelgarejo/badge-api/pkg/response"
	"github.com/go-chi/chi"
)

// BadgeRouter is the router of the badges.
type BadgeRouter struct {
	Repository badge.Repository
}

// CreateHandler Create a new badge.
func (pr *BadgeRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	orgIDStr := chi.URLParam(r, "org_id")

	orgID, err := strconv.Atoi(orgIDStr)
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	var p badge.Badge
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = pr.Repository.Create(ctx, uint(orgID), &p)
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), p.ID))
	_ = response.JSON(w, r, http.StatusCreated, response.Map{"badge": p})
}

// GetByUserHandler response badges by user id.
func (pr *BadgeRouter) GetByUserHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "user_id")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	badges, err := pr.Repository.GetByUser(ctx, uint(userID))
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	_ = response.JSON(w, r, http.StatusOK, response.Map{"badges": badges})
}

// GetByUserHandler response badges by user id.
func (br *BadgeRouter) GetByOrgHandler(w http.ResponseWriter, r *http.Request) {
	orgIDStr := chi.URLParam(r, "org_id")

	orgID, err := strconv.Atoi(orgIDStr)
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	badges, err := br.Repository.GetByOrg(ctx, uint(orgID))
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	_ = response.JSON(w, r, http.StatusOK, response.Map{"badges": badges})
}

// AssignBadgeHandler assign new badge to user
func (br *BadgeRouter) AssignBadgeHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	badgeIDStr := r.URL.Query().Get("badge_id")
	badgeID, err := strconv.Atoi(badgeIDStr)
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = br.Repository.AssignBadge(ctx, uint(userID), uint(badgeID))
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	_ = response.JSON(w, r, http.StatusCreated, response.Map{"badge": "Assigned"})
}

// Routes returns badge router with each endpoint.
func (pr *BadgeRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Authorize)

	r.Get("/user/{user_id}", pr.GetByUserHandler)

	r.Post("/user/{user_id}", pr.AssignBadgeHandler)

	r.Get("/org/{org_id}", pr.GetByOrgHandler)

	r.Post("/org/{org_id}", pr.CreateHandler)

	return r
}
