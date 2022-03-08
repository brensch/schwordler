package main

import (
	"fmt"
	"testing"

	"github.com/brensch/battleword"
)

type wordPossibleCase struct {
	new         string
	guessResult battleword.GuessResult
	isPossible  bool
}

var (
	wordPossibleCases = []wordPossibleCase{
		{"beast", battleword.GuessResult{"beast", []int{0, 2, 2, 2, 2}}, false},
		{"beast", battleword.GuessResult{"beast", []int{2, 2, 2, 2, 2}}, true},
		{"digit", battleword.GuessResult{"beast", []int{2, 2, 2, 2, 2}}, false},
		{"pbliy", battleword.GuessResult{"beast", []int{1, 0, 0, 0, 0}}, true},
		{"eefts", battleword.GuessResult{"beast", []int{0, 1, 0, 1, 0}}, false},
		{"effff", battleword.GuessResult{"beest", []int{0, 1, 0, 0, 0}}, true},
		{"effef", battleword.GuessResult{"beest", []int{0, 1, 0, 0, 0}}, false},
		{"iouuu", battleword.GuessResult{"beast", []int{0, 1, 0, 0, 0}}, true},
		{"feast", battleword.GuessResult{"beast", []int{0, 2, 2, 2, 2}}, true},
		{"feest", battleword.GuessResult{"fstee", []int{2, 1, 1, 1, 1}}, true},
		{"maybj", battleword.GuessResult{"fstee", []int{2, 0, 0, 0, 0}}, false},
	}
)

func TestWordPossible(t *testing.T) {

	for _, wordPossibleCase := range wordPossibleCases {
		possible := WordPossible(wordPossibleCase.new, wordPossibleCase.guessResult)
		fmt.Println(possible)
		if possible != wordPossibleCase.isPossible {
			t.Log("got wrong result", wordPossibleCase.new, wordPossibleCase.guessResult.Guess)
			t.Fail()
		}
	}

}

func BenchmarkWordPossible(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WordPossible("beast", battleword.GuessResult{"beast", []int{2, 2, 1, 1, 1}})
	}
}

func TestGuessWord(t *testing.T) {
	s := &store{}

	prevGuesses := []battleword.GuessResult{
		{"beast", []int{0, 0, 1, 0, 0}},
		{"found", []int{0, 0, 2, 0, 0}},
		{"laugh", []int{1, 2, 2, 0, 0}},
	}

	possibleWords, err := s.GuessWord(prevGuesses)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(possibleWords)
}

func TestGuessWord2(t *testing.T) {
	s := &store{}

	prevGuesses := []battleword.GuessResult{
		{"beast", []int{0, 0, 1, 0, 0}},
		{"found", []int{0, 0, 2, 0, 0}},
		{"laugh", []int{1, 2, 2, 0, 0}},
	}

	possibleWords, err := s.GuessWord2(prevGuesses)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(possibleWords)
}

func BenchmarkGuessWord(b *testing.B) {

	s := &store{}

	prevGuesses := []battleword.GuessResult{
		{"beast", []int{0, 0, 1, 0, 0}},
		{"found", []int{0, 0, 2, 0, 0}},
		{"laugh", []int{1, 2, 2, 0, 0}},
	}
	for i := 0; i < b.N; i++ {
		s.GetPossibleWords(prevGuesses)
	}
}

func TestResultToInt(t *testing.T) {

	for i := 0; i < IntPow(3, 5); i++ {
		result := CodeToResult(i)
		code := ResultToCode(result)
		if code != i {
			t.Logf("got %d when expected %d", code, i)
			t.FailNow()
		}
	}

}

func TestGetWordDistribution(t *testing.T) {
	s := &store{}

	word := "beast"
	distribution := s.GetWordDistribution(word, AllWords)

	for i := range distribution {
		fmt.Printf("dist: %d, words: %+v\n", CodeToResult(i), len(distribution[i]))
	}

}

func BenchmarkGetWordDistribution(b *testing.B) {
	s := &store{}

	word := "beast"

	for i := 0; i < b.N; i++ {
		s.GetWordDistribution(word, AllWords)

	}

}

func BenchmarkGetWordDistributionCount(b *testing.B) {
	s := &store{}

	word := "beast"

	for i := 0; i < b.N; i++ {
		s.GetWordDistributionCount(word, AllWords)

	}

}

func TestGuessWordFull(t *testing.T) {

	s := &store{}

	var prevGuessResults []battleword.GuessResult

	answer := "event"

	for {
		guess, err := s.GuessWord2(prevGuessResults)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		result := battleword.GetResult(guess, answer)

		prevGuessResults = append(prevGuessResults, result)
		if guess == answer {
			t.Log("got the answer")
			break
		}

		if len(prevGuessResults) == 6 {
			break
		}

	}

	t.Logf("finished in %d turns, %+v", len(prevGuessResults), prevGuessResults)
	for _, guessResult := range prevGuessResults {
		fmt.Println(guessResult.Guess, guessResult.Result)
	}

}

func TestGetDistributionExpectedRemainingAnswers(t *testing.T) {

	s := &store{}

	word := "adieu"
	distribution := s.GetWordDistribution(word, AllWords)

	fmt.Println(s.GetDistributionExpectedRemainingAnswers(len(AllWords), distribution))
}
