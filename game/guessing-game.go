package game

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sflewis2970/guessing-game-api/common"
	"github.com/sflewis2970/guessing-game-api/messaging"
)

const (
	// Redis server settings
	NBR_RANGE          int = 100
	NBR_STARTING_VALUE int = 1
	NBR_MAX_ATTEMPTS   int = 5
)

type GuessingGame struct {
	NumberRange         int
	MinimumVal          int
	MaxNumberOfAttempts int
}

var guessingGame *GuessingGame

func (gg *GuessingGame) generateSecretNumber() int {
	return common.GenerateFloat(float64(gg.NumberRange), float64(gg.MinimumVal))
}

func (gg *GuessingGame) StartGame() (int, messaging.GuessingGameData) {
	log.Print("Entering game.StartGame...")

	var ggData messaging.GuessingGameData
	ggData.GameID = uuid.New().String()
	ggData.Timestamp = common.FormattedTime(time.Now(), "Mon Jan 2 15:04:05 2006")
	generatedNbr := gg.generateSecretNumber()

	log.Print("Generated secret number: ", generatedNbr)
	log.Print("Exiting game.StartGame...")
	log.Print("")

	return generatedNbr, ggData
}

func (gg *GuessingGame) CheckSecretNumber(guessNbr int, ggData messaging.GuessingGameData) messaging.GuessingGameData {
	log.Print("Entering game.CheckSecretNumber...")

	log.Print("guess number: ", guessNbr)
	log.Print("game ID: ", ggData.GameID)
	log.Print("generated number: ", ggData.GeneratedNbr)
	log.Print("number of attempts: ", ggData.NbrOfAttempts)

	// Get data from model
	ggData.NbrOfAttempts++
	ggData.GuessNbr = guessNbr
	ggData.GameOver = false
	ggData.Timestamp = common.FormattedTime(time.Now(), "Mon Jan 2 15:04:05 2006")
	if guessNbr < ggData.GeneratedNbr {
		ggData.Message = "Guess is too low"
		return ggData
	} else if guessNbr > ggData.GeneratedNbr {
		ggData.Message = "Guess is too high"
		return ggData
	}

	ggData.GameOver = true
	ggData.Message = "Congrats! Guess is correct!"

	log.Print("Exiting game.CheckSecretNumber...")
	log.Print("")

	return ggData
}

func (gg *GuessingGame) SetGameSettings(nbrRange int, minVal int, maxNbrOfAttempts int) {
	log.Print("Entering game.SetGameSettings...")

	guessingGame.NumberRange = nbrRange
	guessingGame.MinimumVal = minVal
	guessingGame.MaxNumberOfAttempts = maxNbrOfAttempts

	log.Print("Exiting game.SetGameSettings...")
	log.Print("")
}

func NewGuessingGame() *GuessingGame {
	guessingGame = new(GuessingGame)

	guessingGame.SetGameSettings(NBR_RANGE, NBR_STARTING_VALUE, NBR_MAX_ATTEMPTS)

	return guessingGame
}
