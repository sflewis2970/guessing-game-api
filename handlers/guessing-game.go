package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/sflewis2970/guessing-game-api/game"
	"github.com/sflewis2970/guessing-game-api/messaging"
	"github.com/sflewis2970/guessing-game-api/models"
)

type GuessingGameHandler struct {
	guessingGame      *game.GuessingGame
	guessingGameModel *models.GuessingGameModel
}

var ggHandler *GuessingGameHandler

func (ggh *GuessingGameHandler) StartGame(rw http.ResponseWriter, r *http.Request) {
	log.Print("Entering handlers.StartGame...")

	var ggData messaging.GuessingGameData
	var generatedNbr int

	// Start game and return request and generated number
	generatedNbr, ggData = ggh.guessingGame.StartGame()

	log.Print("Generated Number: ", generatedNbr)

	// Add item to model
	ggh.guessingGameModel.AddItem(generatedNbr, ggData)

	// Set content-type
	rw.Header().Add("Content-Type", "application/json")

	// Get item that has been saved
	ggNewData, getErr := ggh.guessingGameModel.FindItem(ggData.GameID)
	if getErr != nil {
		log.Print("Error getting item...")

		// Send OK status
		rw.WriteHeader(http.StatusInternalServerError)

		// Encode response
		encodeGGData(rw, ggNewData)
		return
	}

	log.Print("game ID: ", ggNewData.GameID)
	log.Print("generated number: ", ggNewData.GeneratedNbr)
	log.Print("number of attempts: ", ggNewData.NbrOfAttempts)
	log.Print("")

	// Send OK status
	rw.WriteHeader(http.StatusOK)

	// Encode response
	encodeGGData(rw, ggNewData)

	// Display a log message
	log.Print("Sending response to client...")
	log.Print("")
}

func (ggh *GuessingGameHandler) GuessNumber(rw http.ResponseWriter, r *http.Request) {
	log.Print("Entering handlers.GuessNumber...")

	// Get game ID from query parameter
	gameID := r.URL.Query().Get("game_id")

	// Get guessed number from query parameter
	guessNbrStr := r.URL.Query().Get("guess")

	log.Print("Query Parameter, game_id...:", gameID)
	log.Print("Query Parameter, guess...:", guessNbrStr)

	// Set content-type
	rw.Header().Add("Content-Type", "application/json")

	var ggData messaging.GuessingGameData

	guessNbr, convErr := strconv.Atoi(guessNbrStr)
	if convErr != nil {
		log.Print("Error converting string to int...: ", convErr)

		// Update AnswerResponse
		ggData.Message = convErr.Error()

		// Update HTTP Header
		rw.WriteHeader(http.StatusInternalServerError)

		// Write JSON to stream
		encodeGGData(rw, ggData)
		return
	}

	// Check guessed number
	var getErr error

	log.Print("Sending gameID to model...")
	ggData, getErr = ggh.guessingGameModel.FindItem(gameID)
	if getErr != nil {
		log.Print("Error converting string to int...: ", convErr)
		// Update AnswerResponse
		ggData.Message = getErr.Error()

		// Update HTTP Header
		rw.WriteHeader(http.StatusInternalServerError)

		// Write JSON to stream
		encodeGGData(rw, ggData)
		return
	}

	log.Print("game ID: ", ggData.GameID)
	log.Print("generated number: ", ggData.GeneratedNbr)
	log.Print("number of attempts: ", ggData.NbrOfAttempts)

	log.Print("Sending guessNbr and ggRequest to guessing game...")
	ggData = ggh.guessingGame.CheckSecretNumber(guessNbr, ggData)

	if ggData.GameOver {
		delErr := ggh.guessingGameModel.DeleteItem(gameID)
		if delErr != nil {
			ggData.Message = "Error deleting item"

			// Update HTTP Header
			rw.WriteHeader(http.StatusInternalServerError)

			// Write JSON to stream
			encodeGGData(rw, ggData)
			return
		}
	}

	// Send OK status
	rw.WriteHeader(http.StatusOK)

	// Encode response
	encodeGGData(rw, ggData)

	// Display a log message
	log.Print("Sending response to client...")
	log.Print("")
}

func encodeGGData(rw http.ResponseWriter, ggData messaging.GuessingGameData) {
	// Write JSON to stream
	encodeErr := json.NewEncoder(rw).Encode(ggData)
	if encodeErr != nil {
		log.Print("Error encoding json...:", encodeErr)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func NewGuessingGameHandler() *GuessingGameHandler {
	ggHandler = new(GuessingGameHandler)

	ggHandler.guessingGameModel = models.NewGuessingGameModel()
	ggHandler.guessingGame = game.NewGuessingGame()

	return ggHandler
}
