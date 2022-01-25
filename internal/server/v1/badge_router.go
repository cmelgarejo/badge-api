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
	var p badge.Badge
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = pr.Repository.Create(ctx, &p)
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), p.ID))
	_ = response.JSON(w, r, http.StatusCreated, response.Map{"badge": p})
}

// GetAllHandler response all the badges.
func (pr *BadgeRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	badges, err := pr.Repository.GetAll(ctx)
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	_ = response.JSON(w, r, http.StatusOK, response.Map{"badges": badges})
}

// GetOneHandler response one badge by id.
func (pr *BadgeRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	p, err := pr.Repository.GetOne(ctx, uint(id))
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	_ = response.JSON(w, r, http.StatusOK, response.Map{"badge": p})
}

// UpdateHandler update a stored badge by id.
func (pr *BadgeRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
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
	err = pr.Repository.Update(ctx, uint(id), p)
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	_ = response.JSON(w, r, http.StatusOK, nil)
}

// DeleteHandler Remove a badge by ID.
func (pr *BadgeRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = pr.Repository.Delete(ctx, uint(id))
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	_ = response.JSON(w, r, http.StatusOK, response.Map{})
}

// GetByUserHandler response badges by user id.
func (pr *BadgeRouter) GetByUserHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userId")

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

// Routes returns badge router with each endpoint.
func (pr *BadgeRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Authorize)

	r.Get("/user/{user_id}", pr.GetByUserHandler)

	r.Get("/", pr.GetAllHandler)

	r.Post("/", pr.CreateHandler)

	r.Get("/{id}", pr.GetOneHandler)

	r.Put("/{id}", pr.UpdateHandler)

	r.Delete("/{id}", pr.DeleteHandler)

	return r
}
