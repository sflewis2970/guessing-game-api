package controllers

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/sflewis2970/guessing-game-api/handlers"
)

// Controller structure defines teh layout of the Controller
type Controller struct {
	Router              *mux.Router
	guessingGameHandler *handlers.GuessingGameHandler
}

// Package controllers object
var controller *Controller

func (c *Controller) setupRoutes() {
	// Display log message
	log.Print("Setting up service routes")

	// API routes
	c.Router.HandleFunc("/api/v1/guessing-game/start", c.guessingGameHandler.StartGame).Methods("GET")
	c.Router.HandleFunc("/api/v1/guessing-game/guess", c.guessingGameHandler.GuessNumber).Methods("POST")
}

// NewController function create a new Controller and initializes new Controller object
func NewController() *Controller {
	// Create controllers component
	log.Print("Creating controllers object...")
	controller = new(Controller)

	// Trivia handler
	controller.guessingGameHandler = handlers.NewGuessingGameHandler()

	// Set controllers routes
	controller.Router = mux.NewRouter()
	controller.setupRoutes()

	return controller
}
