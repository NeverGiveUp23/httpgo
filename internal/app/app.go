package app

import (
	"fmt"
	"github.com/nevergiveup23/httpgo/internal/api"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*Application, error) {
	// logger instead of println
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	// stores will go here

	// handlers will go here
	workoutHandler := api.NewWorkoutHandler()

	// reference to application
	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "status is available\n")
}
