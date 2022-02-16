package main

import (
	"fmt"

	"github.com/brensch/battleword"
)

func (s *store) GuessWord(prevGuesses []string, prevResults [][]int) (string, error) {

	words, err := s.GetPossibleWords(prevGuesses, prevResults)
	if err != nil {
		return "", err
	}

	if len(words) == 0 {
		return "", fmt.Errorf("got no possible words. something is up.")
	}

	return words[0], nil
}

func (s *store) GetPossibleWords(prevGuesses []string, prevResults [][]int) ([]string, error) {

	possibleWords := CommonWords

	for i, prevGuess := range prevGuesses {

		var newPossibleWords []string
		for _, newGuess := range possibleWords {
			if WordPossible(newGuess, prevGuess, prevResults[i]) {
				// fmt.Println(newGuess, prevGuess, prevResults[i])
				newPossibleWords = append(newPossibleWords, newGuess)
			}

		}
		possibleWords = newPossibleWords
	}

	return possibleWords, nil
}

func WordPossible(newGuess, prevGuess string, prevResult []int) bool {

	// i figured this out by looking at all the results. kinda cool. plz don't steal.
	newResult := battleword.GetResult(prevGuess, newGuess)
	// fmt.Println(newGuess, prevGuess)
	// fmt.Println(newResult)
	// fmt.Println(prevResult)
	for i := 0; i < len(newGuess); i++ {
		if newResult[i] > prevResult[i] {
			return false
		}

		if prevResult[i] == 2 && newResult[i] < 2 {
			return false
		}
	}

	return true

}
