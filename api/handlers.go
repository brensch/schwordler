package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/brensch/battleword"
	"github.com/brensch/schwordler"
)

func (api *api) HandleDoGuess(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		return
	}

	var gameState battleword.PlayerGameState
	err := json.NewDecoder(r.Body).Decode(&gameState)
	if err != nil {
		api.s.Log.WithError(err).Error("could not decode prevGuesses from engine")
		return
	}

	gameStateBytes, _ := json.Marshal(gameState)
	fmt.Println(string(gameStateBytes))

	word, err := api.s.GuessWord(gameState.GuessResults)
	if err != nil {
		api.s.Log.WithError(err).Error("problem guessing word")
		return
	}
	fmt.Println(word)

	guess := battleword.Guess{
		Guess: word,
		Shout: schwordler.RandomShout(),
	}

	err = json.NewEncoder(w).Encode(guess)
	if err != nil {
		log.Println(err)
		return
	}
}

func (api *api) HandleReceiveResults(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		return
	}

	var finalState battleword.PlayerMatchResults
	err := json.NewDecoder(r.Body).Decode(&finalState)
	if err != nil {
		log.Println(err)
		return
	}

	// var us *battleword.Player
	// for _, player := range finalState.Resultapi.s.Players {
	// 	if player.ID == finalState.PlayerID {
	// 		us = player
	// 	}
	// }

	// if us == nil {
	// 	log.Println("we weren't in the resultapi.s. strange")
	// 	return
	// }

	// log.Println("the game concluded, and the engine sent me the final state for all players:", string(finalStateJSON))
	log.Println("our final statistics were:")
	// log.Printf("accuracy: %f%%", 100*float64(uapi.s.Summary.GamesWon)/float64(len(finalState.Resultapi.s.Games)))
	// log.Printf("speed: %s", uapi.s.Summary.TotalTime)
	// log.Printf("average guesses: %f", float64(uapi.s.Summary.TotalGuesses)/float64(len(finalState.Resultapi.s.Games)))

}

func (api *api) HandleDoPing(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		return
	}

	log.Println("received ping")

	definition := &battleword.PlayerDefinition{
		Name:        "schwordler",
		Description: "the brave",
	}

	err := json.NewEncoder(w).Encode(definition)
	if err != nil {
		log.Println(err)
		return
	}
}
