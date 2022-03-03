package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/brensch/battleword"
)

func (s *store) HandleDoGuess(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		return
	}

	var prevGuesses battleword.PlayerGameState
	err := json.NewDecoder(r.Body).Decode(&prevGuesses)
	if err != nil {
		s.log.WithError(err).Error("could not decode prevGuesses from engine")
		return
	}

	word, err := s.GuessWord2(prevGuesses.Guesses, prevGuesses.Results)
	if err != nil {
		s.log.WithError(err).Error("problem guessing word")
		return
	}

	guess := battleword.Guess{
		Guess: word,
		Shout: RandomShout(),
	}

	err = json.NewEncoder(w).Encode(guess)
	if err != nil {
		log.Println(err)
		return
	}
}

func (s *store) HandleReceiveResults(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		return
	}

	var finalState battleword.PlayerMatchResults
	err := json.NewDecoder(r.Body).Decode(&finalState)
	if err != nil {
		log.Println(err)
		return
	}

	var us *battleword.Player
	for _, player := range finalState.Results.Players {
		if player.ID == finalState.PlayerID {
			us = player
		}
	}

	if us == nil {
		log.Println("we weren't in the results. strange")
		return
	}

	// log.Println("the game concluded, and the engine sent me the final state for all players:", string(finalStateJSON))
	log.Println("our final statistics were:")
	log.Printf("accuracy: %f%%", 100*float64(us.Summary.GamesWon)/float64(len(finalState.Results.Games)))
	log.Printf("speed: %s", us.Summary.TotalTime)
	log.Printf("average guesses: %f", float64(us.Summary.TotalGuesses)/float64(len(finalState.Results.Games)))

}

func (s *store) HandleDoPing(w http.ResponseWriter, r *http.Request) {

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
