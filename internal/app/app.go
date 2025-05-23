package app

import (
	"database/sql"
	"fmt"
	"github.com/nevergiveup23/httpgo/internal/api"
	"github.com/nevergiveup23/httpgo/internal/store"
	"github.com/nevergiveup23/httpgo/migrations"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}
	// logger instead of println
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	// stores will go here
	workoutStore := store.NewPostgresWorkoutStore(pgDB)
	// handlers will go here
	workoutHandler := api.NewWorkoutHandler(workoutStore)

	// reference to application
	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDB,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "status is available\n")
}
