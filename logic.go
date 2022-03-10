package schwordler

import (
	"fmt"

	"github.com/brensch/battleword"
)

func (s *Store) GuessWord(prevGuessResults []battleword.GuessResult) (string, error) {

	if len(prevGuessResults) == 0 {
		return "crane", nil
	}

	possibleAnswers, err := s.GetPossibleWords(prevGuessResults)
	if err != nil {
		return "", err
	}

	if len(possibleAnswers) == 1 {
		return possibleAnswers[0], nil
	}

	expectedRemainingWords := make([]float64, len(possibleAnswers))
	distributions := make([][][]string, len(possibleAnswers))
	for i, guess := range possibleAnswers {
		distribution := s.GetWordDistribution(guess, possibleAnswers)
		expectedRemainingWords[i] = s.GetDistributionExpectedRemainingAnswers(len(possibleAnswers), distribution)
		distributions[i] = distribution
	}

	var bestWord string
	bestWordExpectedRemaining := float64(len(possibleAnswers))
	for i, expectedRemainingWord := range expectedRemainingWords {
		if expectedRemainingWord <= bestWordExpectedRemaining {
			bestWord = possibleAnswers[i]
			bestWordExpectedRemaining = expectedRemainingWord
		}

	}

	return bestWord, nil
}

func (s *Store) GetPossibleWords(prevGuessResults []battleword.GuessResult) ([]string, error) {

	possibleWords := CommonWords

	for _, prevGuessResult := range prevGuessResults {

		var newPossibleWords []string
		for _, newGuess := range possibleWords {
			if WordPossible(newGuess, prevGuessResult) {
				newPossibleWords = append(newPossibleWords, newGuess)
			}

		}
		possibleWords = newPossibleWords
	}

	return possibleWords, nil
}

func ResultToCode(result []int) int {
	answer := 0
	for i, j := range result {
		base := IntPow(3, i)
		answer += base * j
	}
	return answer
}

func CodeToResult(code int) []int {
	var result []int
	i := 1
	for {
		base := IntPow(3, i)
		rem := code % base
		result = append(result, rem/IntPow(3, i-1))
		// fmt.Println(code, base, rem, rem/IntPow(3, i-1), result)
		code -= rem
		if code == 0 {
			return result
		}
		i++
		if i > 10 {
			fmt.Println(i)
			panic("too many iterations")
		}
	}

}

// 83

// IntPower return x^y but as ints for speed
func IntPow(x, y int) int {
	result := 1
	for i := 0; i < y; i++ {
		result = result * x
	}

	return result

}

type ResultOdds struct {
	Result []int
	Words  []string
}

func (s *Store) GetWordDistribution(word string, possibleAnswers []string) [][]string {

	distribution := make([][]string, IntPow(3, len(word)))
	for _, possibleAnswer := range possibleAnswers {
		result := battleword.GetResult(word, possibleAnswer)
		resultCode := ResultToCode(result.Result)
		distribution[resultCode] = append(distribution[resultCode], possibleAnswer)
	}

	return distribution
}

func (s *Store) GetWordDistributionCount(word string, possibleAnswers []string) []int {

	distribution := make([]int, IntPow(3, len(word)))
	for _, possibleAnswer := range possibleAnswers {
		result := battleword.GetResult(word, possibleAnswer)
		resultCode := ResultToCode(result.Result)
		distribution[resultCode]++
	}

	return distribution
}

func (s *Store) GetDistributionExpectedRemainingAnswers(wordCount int, distribution [][]string) float64 {

	expectedRemainingAnswer := float64(0)

	for _, wordList := range distribution {

		expectedRemainingAnswer += float64(len(wordList)) / float64(wordCount) * float64(len(wordList))

	}

	return expectedRemainingAnswer
}

func WordPossible(newGuess string, prevGuessResult battleword.GuessResult) bool {

	// i figured this out by looking at all the results. kinda cool. plz don't steal.
	newResult := battleword.GetResult(prevGuessResult.Guess, newGuess)
	// fmt.Println(newGuess, prevGuess)
	// fmt.Println(newResult)
	// fmt.Println(prevResult)
	for i := 0; i < len(newGuess); i++ {
		if newResult.Result[i] > prevGuessResult.Result[i] {
			return false
		}

		if prevGuessResult.Result[i] == 2 && newResult.Result[i] < 2 {
			return false
		}
	}

	return true

}
