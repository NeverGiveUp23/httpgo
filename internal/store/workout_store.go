package store

import "database/sql"

type Workout struct {
	ID              int            `json:"id"`
	Title           string         `json:"title"`
	Description     string         `json:"description"`
	DurationMinutes int            `json:"durationMinutes"`
	CaloriesBurned  int            `json:"calories_burned"`
	Entries         []WorkoutEntry `json:"entries"`
}

type WorkoutEntry struct {
	ID              int      `json:"id"`
	ExerciseName    string   `json:"exerciseName"`
	_Sets           int      `json:"_sets"`
	Reps            *int     `json:"reps"` // check to be nil or not
	DurationSeconds *int     `json:"durationSeconds"`
	Weight          *float64 `json:"weight"`
	Notes           string   `json:"notes"`
	OrderIndex      int      `json:"order_index"`
}

type PostgresWorkoutStore struct {
	db *sql.DB
}

func NewPostgresWorkoutStore(db *sql.DB) *PostgresWorkoutStore {
	return &PostgresWorkoutStore{
		db: db,
	}
}

// application layer to communicate with the postgres DB, in case we need to migrate to another DB, the interface will communicate between the two not the Postgres struct as above
type WorkoutStore interface {
	CreateWorkout(*Workout) (*Workout, error)
	GetWorkoutById(id int64) (*Workout, error)
	UpdateWorkout(*Workout) error
}

func (pg *PostgresWorkoutStore) CreateWorkout(workout *Workout) (*Workout, error) {
	tx, err := pg.db.Begin() // tx = transaction
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // causes a rollback in case going against ACID principles

	query := `INSERT INTO workouts (title, description, duration_minutes, calories_burned) VALUES ($1, $2, $3, $4) RETURNING id`

	err = tx.QueryRow(query, workout.Title, workout.Description, workout.DurationMinutes, workout.CaloriesBurned).Scan(&workout.ID)
	if err != nil {
		return nil, err
	}

	// insert entries
	for _, entry := range workout.Entries {
		query := `
		INSERT INTO workout_entries (workout_id, exercise_name, _sets, reps, duration_seconds, weight, notes, order_index) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id
		`
		err = tx.QueryRow(query, workout.ID, entry.ExerciseName, entry._Sets, entry.Reps, entry.DurationSeconds, entry.Weight, entry.Notes, entry.OrderIndex).Scan(&entry.ID)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return workout, nil
}

func (pg *PostgresWorkoutStore) GetWorkoutById(id int64) (*Workout, error) {
	workout := &Workout{}
	query := `
		SELECT id , title, description, duration_minutes, calories_burned FROM workouts WHERE id = $1`
	err := pg.db.QueryRow(query, id).Scan(&workout.ID, &workout.Title, &workout.Description, &workout.DurationMinutes, &workout.CaloriesBurned)
	if err != nil {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	// gets entries
	entryQuery := `
		SELECT id, exercise_name, _sets, reps, duration_seconds, weight, notes, order_index FROM workout_entries WHERE workout_id = $1 ORDER BY order_index
	`

	rows, err := pg.db.Query(entryQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var entry WorkoutEntry
		err = rows.Scan(
			&entry.ID,
			&entry.ExerciseName,
			&entry._Sets,
			&entry.Reps,
			&entry.DurationSeconds,
			&entry.Weight,
			&entry.Notes,
			&entry.OrderIndex,
		)
		if err != nil {
			return nil, err
		}
		workout.Entries = append(workout.Entries, entry)
	}

	return workout, nil
}
