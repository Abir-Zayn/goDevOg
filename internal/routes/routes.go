package routes

//Chi is a lightweight, idiomatic, and composable router for building Go HTTP services 
//that implements standard library interfaces and allows 
//for easy routing and middleware handling

import (
	"github.com/Abir-Zayn/goDevOg.git/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", app.HealthCheck)
	r.Get("/workouts/{id}",app.WorkoutHandler.HandleGetWorkoutByID)
	r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkOut)
	return r
}
