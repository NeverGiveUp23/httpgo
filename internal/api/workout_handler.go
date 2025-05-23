package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/nevergiveup23/httpgo/internal/store"
	"net/http"
	"strconv"
)

type WorkoutHandler struct {
	workoutStore store.WorkoutStore //
}

func NewWorkoutHandler() *WorkoutHandler {
	return &WorkoutHandler{}
}

func (wh *WorkoutHandler) HandleGetWorkoutByID(w http.ResponseWriter, r *http.Request) {
	// retrieve uuid
	paramsWourkoutID := chi.URLParam(r, "id")

	if paramsWourkoutID == "" {
		http.NotFound(w, r)
		return
	}
	workoutID, err := strconv.ParseInt(paramsWourkoutID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "this is the workout id %d\n", workoutID)
}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout store.Workout
	err := json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to create workout", http.StatusInternalServerError)
		return
	}

}
