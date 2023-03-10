package models

import (
	"log"

	"github.com/sflewis2970/guessing-game-api/config"
	"github.com/sflewis2970/guessing-game-api/messaging"
)

type GuessingGameModel struct {
	cfgData    *config.CfgData
	redisModel *RedisModel
}

var guessGameModel *GuessingGameModel

func (ggm *GuessingGameModel) AddItem(generatedNbr int, ggData messaging.GuessingGameData) {
	log.Print("Entering models.AddItem...")

	// Add generated number to request before inserting data
	ggData.GeneratedNbr = generatedNbr

	log.Print("Generated Number: ", ggData.GeneratedNbr)
	log.Print("game ID: ", ggData.GameID)
	log.Print("generated number: ", ggData.GeneratedNbr)
	log.Print("number of attempts: ", ggData.NbrOfAttempts)

	insertErr := ggm.redisModel.Insert(ggData)
	if insertErr != nil {
		log.Print("Error inserting item...")
		return
	} else {
		log.Print("item successfully inserted...")
	}

	log.Print("Exiting models.AddItem...")
	log.Print("")
}

func (ggm *GuessingGameModel) FindItem(id string) (messaging.GuessingGameData, error) {
	log.Print("Entering models.FindItem...")

	ggRequest, getErr := ggm.redisModel.Get(id)
	if getErr != nil {
		log.Print("Error getting item...")
		return messaging.GuessingGameData{}, getErr
	} else {
		log.Print("item successfully retrieved...")
	}

	log.Print("Generated Number: ", ggRequest.GeneratedNbr)
	log.Print("game ID: ", ggRequest.GameID)
	log.Print("generated number: ", ggRequest.GeneratedNbr)
	log.Print("number of attempts: ", ggRequest.NbrOfAttempts)

	log.Print("Exiting models.FindItem...")
	log.Print("")

	return ggRequest, nil
}

func (ggm *GuessingGameModel) DeleteItem(id string) error {
	log.Print("Entering models.DeleteItem...")

	ggRequest, getErr := ggm.FindItem(id)
	if getErr != nil {
		log.Print("Error getting item...")
		return getErr
	}

	delErr := ggm.DeleteItem(ggRequest.GameID)
	if delErr != nil {
		log.Print("Error deleting item...")
		return delErr
	}

	log.Print("Exiting models.DeleteItem...")
	log.Print("")

	return nil
}

func NewGuessingGameModel() *GuessingGameModel {
	log.Print("Creating model object...")
	guessGameModel := new(GuessingGameModel)

	// Get config data
	guessGameModel.cfgData = config.NewConfig().LoadCfgData()

	// New model (cacheModel)
	guessGameModel.redisModel = NewRedisModel()

	return guessGameModel
}
