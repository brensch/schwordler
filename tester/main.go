package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/brensch/battleword"
	"github.com/brensch/schwordler"
	"github.com/sirupsen/logrus"
)

type Result struct {
	Starter string        `json:"starter,omitempty"`
	Time    time.Duration `json:"time,omitempty"`
	Average float64       `json:"average,omitempty"`
}

func main() {
	s := schwordler.InitStore(logrus.New())

	concurrentSolverCHAN := make(chan struct{}, 6)
	resultCHAN := make(chan Result)
	var wgGenerator, wgListener sync.WaitGroup

	wgListener.Add(1)
	go func() {
		var allResults []Result
		lowestAverage := float64(6)
		var lowestAverageWord string

		defer wgListener.Done()
		for result := range resultCHAN {
			allResults = append(allResults, result)
			if result.Average < lowestAverage {
				lowestAverage = result.Average
				lowestAverageWord = result.Starter
			}
			fmt.Printf("finished %s in %f seconds with average %f\n", result.Starter, result.Time.Seconds(), result.Average)

		}

		fs, err := os.Create("./result.json")
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

		fmt.Println("best word is", lowestAverageWord)

	}()

	for _, starter := range battleword.CommonWords {
		wgGenerator.Add(1)
		go func(starter string) {
			allResults := make([][]battleword.GuessResult, len(battleword.CommonWords))
			defer wgGenerator.Done()

			concurrentSolverCHAN <- struct{}{}
			defer func() { <-concurrentSolverCHAN }()

			fmt.Printf("starting %s\n", starter)
			start := time.Now()

			for i, answer := range battleword.CommonWords {
				// do the first result from the current start word.
				var results []battleword.GuessResult

				for {
					guess, err := s.GuessWord(results, starter)
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
						// fmt.Println("failed", answer)
						break
					}
				}
			}

			end := time.Now()

			resultCHAN <- Result{
				Starter: starter,
				Time:    end.Sub(start),
				Average: scoreList(allResults),
			}

			fs, err := os.Create(fmt.Sprintf("./deepdive/%s.json.gz", starter))
			if err != nil {
				fmt.Println(err)
				return
			}
			defer fs.Close()

			gz := gzip.NewWriter(fs)
			defer gz.Close()

			err = json.NewEncoder(gz).Encode(allResults)
			if err != nil {
				fmt.Println(err)
				return
			}
		}(starter)
	}

	wgGenerator.Wait()
	close(resultCHAN)
	wgListener.Wait()

}

func scoreList(input [][]battleword.GuessResult) float64 {
	var length int
	for _, guessResult := range input {
		length += len(guessResult)
	}

	return float64(length) / float64(len(input))
}
