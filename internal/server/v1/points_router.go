package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/cmelgarejo/badge-api/internal/middleware"
	"github.com/cmelgarejo/badge-api/pkg/points"
	"github.com/cmelgarejo/badge-api/pkg/response"
	"github.com/go-chi/chi"
)

// PointsRouter is the router of the badges.
type PointsRouter struct {
	Repository points.Repository
}

// AssignPointsHandler assign new points to user
func (pr *PointsRouter) AssignPointsHandler(w http.ResponseWriter, r *http.Request) {
	var p points.Point
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = pr.Repository.AssignPoints(ctx, &p)
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), p.ID))
	_ = response.JSON(w, r, http.StatusCreated, response.Map{"badge": p})
}

// GetByUserHandler response all the badges.
func (pr *PointsRouter) GetByUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var startDate, endDate *time.Time
	idStr := chi.URLParam(r, "user_id")
	startStr := chi.URLParam(r, "start_date")
	endStr := chi.URLParam(r, "start_date")
	fmt.Printf("get points: %v - %v - %v\n\n", idStr, startStr, endStr)
	if startStr != "" {
		start, err := time.Parse(time.RFC3339, startStr)
		if err != nil {
			_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
			return
		}
		startDate = &start
	}
	if endStr != "" {
		end, err := time.Parse(time.RFC3339, endStr)
		if err != nil {
			_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
			return
		}
		endDate = &end
	}
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
			return
		}

		points, err := pr.Repository.GetPoints(ctx, uint(id), startDate, endDate)
		if err != nil {
			_ = response.HTTPError(w, r, http.StatusNotFound, err.Error())
			return
		}
		_ = response.JSON(w, r, http.StatusOK, response.Map{"points": points})
	}

}

// Routes returns badge router with each endpoint.
func (pr *PointsRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Authorize)

	r.Get("/{user_id}", pr.GetByUserHandler)

	r.Post("/", pr.AssignPointsHandler)

	return r
}
