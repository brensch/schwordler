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

	word, err := s.GuessWord(prevGuesses.Guesses, prevGuesses.Results)
	if err != nil {
		s.log.WithError(err).Error("problem guessing word")
		return
	}

	log.Println("based on previous state, i will make the completely random guess:", word)

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

	var finalState battleword.Match
	err := json.NewDecoder(r.Body).Decode(&finalState)
	if err != nil {
		log.Println(err)
		return
	}

	finalStateJSON, _ := json.Marshal(finalState)

	log.Println("the game concluded, and the engine sent me the final state for all players:", string(finalStateJSON))

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
