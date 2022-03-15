package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/brensch/battleword"
	"github.com/brensch/schwordler"
	"github.com/sirupsen/logrus"
)

func main() {
	s := schwordler.InitStore(logrus.New())
	numberGuesses := len(battleword.CommonWords)
	allResults := make([][]battleword.GuessResult, numberGuesses)

	start := time.Now()
	j := 0
	for i, answer := range battleword.CommonWords {
		j++
		if j > numberGuesses {
			break
		}
		var results []battleword.GuessResult

		for {

			guess, err := s.GuessWord(results)
			if err != nil {
				fmt.Println(err)
				break
			}

			result := battleword.GetResult(guess, answer)
			results = append(results, battleword.GuessResult{
				Guess:  guess,
				Result: result,
			})

			if guess == answer {
				allResults[i] = results
				// fmt.Println("guessed it", answer)
				break
			}

			if len(results) > 5 {
				allResults[i] = append(results, battleword.GuessResult{ID: "failed", Guess: answer})
				fmt.Println("failed", answer)
				break
			}
		}
	}

	end := time.Now()

	fmt.Printf("finished in %f seconds\n", end.Sub(start).Seconds())
	fs, err := os.Create("guess.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fs.Close()

	err = json.NewEncoder(fs).Encode(allResults)
	if err != nil {
		fmt.Println(err)
		return
	}

	scoreList(allResults)
}

func scoreList(input [][]battleword.GuessResult) {
	var length int
	for _, guessResult := range input {
		length += len(guessResult)
	}
	fmt.Println("total length", length)
	fmt.Println("total input length", len(input))
	fmt.Println("average length", float64(length)/float64(len(input)))
}
