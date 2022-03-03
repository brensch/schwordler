package main

import (
	"fmt"
	"testing"

	"github.com/brensch/battleword"
)

type wordPossibleCase struct {
	new        string
	prev       string
	result     []int
	isPossible bool
}

var (
	wordPossibleCases = []wordPossibleCase{
		{"beast", "beast", []int{0, 2, 2, 2, 2}, false},
		{"beast", "beast", []int{2, 2, 2, 2, 2}, true},
		{"digit", "beast", []int{2, 2, 2, 2, 2}, false},
		{"pbliy", "beast", []int{1, 0, 0, 0, 0}, true},
		{"eefts", "beast", []int{0, 1, 0, 1, 0}, false},
		{"effff", "beest", []int{0, 1, 0, 0, 0}, true},
		{"effef", "beest", []int{0, 1, 0, 0, 0}, false},
		{"iouuu", "beast", []int{0, 1, 0, 0, 0}, true},
		{"feast", "beast", []int{0, 2, 2, 2, 2}, true},
		{"feest", "fstee", []int{2, 1, 1, 1, 1}, true},
		{"maybj", "fstee", []int{2, 0, 0, 0, 0}, false},
	}
)

func TestWordPossible(t *testing.T) {

	for _, wordPossibleCase := range wordPossibleCases {
		possible := WordPossible(wordPossibleCase.new, wordPossibleCase.prev, wordPossibleCase.result)
		fmt.Println(possible)
		if possible != wordPossibleCase.isPossible {
			t.Log("got wrong result", wordPossibleCase.new, wordPossibleCase.prev)
			t.Fail()
		}
	}

}

func BenchmarkWordPossible(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WordPossible("beast", "beast", []int{2, 2, 1, 1, 1})
	}
}

func TestGuessWord(t *testing.T) {
	s := &store{}

	prevGuesses := []string{
		"beast",
		"found",
		"laugh",
	}

	prevResults := [][]int{
		{0, 0, 1, 0, 0},
		{0, 0, 2, 0, 0},
		{1, 2, 2, 0, 0},
	}

	possibleWords, err := s.GetPossibleWords(prevGuesses, prevResults)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(possibleWords)
}

func TestGuessWord2(t *testing.T) {
	s := &store{}

	prevGuesses := []string{
		"beast",
	}

	prevResults := [][]int{
		{0, 0, 1, 0, 0},
	}

	guess, err := s.GuessWord2(prevGuesses, prevResults)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(guess)
}

func BenchmarkGuessWord(b *testing.B) {

	s := &store{}

	prevGuesses := []string{
		"beast",
		"found",
		"laugh",
	}

	prevResults := [][]int{
		{0, 0, 1, 0, 0},
		{0, 0, 2, 0, 0},
		{1, 2, 2, 0, 0},
	}
	for i := 0; i < b.N; i++ {
		s.GetPossibleWords(prevGuesses, prevResults)
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

	var prevGuesses []string
	var prevResults [][]int

	answer := "event"

	for {
		guess, err := s.GuessWord2(prevGuesses, prevResults)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		result := battleword.GetResult(guess, answer)

		prevGuesses = append(prevGuesses, guess)
		prevResults = append(prevResults, result)
		if guess == answer {
			t.Log("got the answer")
			break
		}

		if len(prevGuesses) == 6 {
			break
		}

	}

	t.Logf("finished in %d turns, %+v", len(prevGuesses), prevGuesses)
	for i, j := range prevResults {
		fmt.Println(prevGuesses[i])
		fmt.Println(j)
	}

}

func TestGetDistributionExpectedRemainingAnswers(t *testing.T) {

	s := &store{}

	word := "adieu"
	distribution := s.GetWordDistribution(word, AllWords)

	fmt.Println(s.GetDistributionExpectedRemainingAnswers(len(AllWords), distribution))
}
