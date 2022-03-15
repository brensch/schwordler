package schwordler

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/brensch/battleword"
)

type wordPossibleCase struct {
	new        string
	prevWord   string
	prevResult []int
	// guessResult battleword.GuessResult
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
		{"iouuu", "beast", []int{0, 1, 0, 0, 0}, false},
		{"feast", "beast", []int{0, 2, 2, 2, 2}, true},
		{"feest", "fstee", []int{2, 1, 1, 1, 1}, true},
		{"maybj", "fstee", []int{2, 0, 0, 0, 0}, false},
		{"dopey", "boney", []int{0, 2, 1, 2, 2}, false},
	}
)

func TestWordPossible(t *testing.T) {

	for _, wordPossibleCase := range wordPossibleCases {
		guessResult := battleword.GuessResult{
			Guess:  wordPossibleCase.prevWord,
			Result: wordPossibleCase.prevResult,
		}
		possible := WordPossible(wordPossibleCase.new, guessResult)
		fmt.Println(possible)
		if possible != wordPossibleCase.isPossible {
			t.Log("got wrong result", wordPossibleCase.new, guessResult.Guess)
			t.Fail()
		}
	}

}

func BenchmarkWordPossible(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WordPossible("beast", battleword.GuessResult{"id", time.Time{}, time.Time{}, "beast", []int{2, 2, 1, 1, 1}})
	}
}

func TestGuessWord(t *testing.T) {
	s := &Store{}

	prevGuesses := []battleword.GuessResult{
		{"id", time.Time{}, time.Time{}, "beast", []int{0, 0, 1, 0, 0}},
		{"id", time.Time{}, time.Time{}, "found", []int{0, 0, 2, 0, 0}},
		{"id", time.Time{}, time.Time{}, "laugh", []int{1, 2, 2, 0, 0}},
	}

	possibleWords, err := s.GuessWord(prevGuesses)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(possibleWords)
}

func BenchmarkGuessWord(b *testing.B) {

	s := &Store{}

	prevGuesses := []battleword.GuessResult{
		{"id", time.Time{}, time.Time{}, "beast", []int{0, 0, 1, 0, 0}},
		{"id", time.Time{}, time.Time{}, "found", []int{0, 0, 2, 0, 0}},
		{"id", time.Time{}, time.Time{}, "laugh", []int{1, 2, 2, 0, 0}},
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
	s := &Store{}

	word := "beast"
	distribution := s.GetWordDistribution(word, AllWords)

	for i := range distribution {
		fmt.Printf("dist: %d, words: %+v\n", CodeToResult(i), len(distribution[i]))
	}

}

func BenchmarkGetWordDistribution(b *testing.B) {
	s := &Store{}

	word := "beast"

	for i := 0; i < b.N; i++ {
		s.GetWordDistribution(word, AllWords)

	}

}

func BenchmarkGetWordDistributionCount(b *testing.B) {
	s := &Store{}

	word := "beast"

	for i := 0; i < b.N; i++ {
		s.GetWordDistributionCount(word, AllWords)

	}

}

func TestGuessWordFull(t *testing.T) {

	s := &Store{}

	var prevGuessResults []battleword.GuessResult

	answer := "pound"

	for {
		guess, err := s.GuessWord(prevGuessResults)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
		fmt.Println("guessing", guess)

		result := battleword.GetResult(guess, answer)
		guessResult := battleword.GuessResult{
			Result: result,
			Guess:  guess,
		}

		prevGuessResults = append(prevGuessResults, guessResult)
		if guess == answer {
			t.Log("got the answer")
			break
		}

		prevResultsBytes, _ := json.Marshal(prevGuessResults)
		fmt.Println(string(prevResultsBytes))

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

	s := &Store{}

	word := "adieu"
	distribution := s.GetWordDistribution(word, AllWords)

	fmt.Println(s.GetDistributionExpectedRemainingAnswers(len(AllWords), distribution))
}

func TestSingleState(t *testing.T) {
	s := &Store{}

	raw := []byte(`
	[{"start":"0001-01-01T00:00:00Z","finish":"0001-01-01T00:00:00Z","guess":"crane","result":[0,0,0,1,1]},{"start":"0001-01-01T00:00:00Z","finish":"0001-01-01T00:00:00Z","guess":"islet","result":[0,1,0,2,0]},{"start":"0001-01-01T00:00:00Z","finish":"0001-01-01T00:00:00Z","guess":"women","result":[0,2,0,2,1]},{"start":"0001-01-01T00:00:00Z","finish":"0001-01-01T00:00:00Z","guess":"boney","result":[0,2,1,2,2]}]
`)

	var state []battleword.GuessResult
	err := json.Unmarshal(raw, &state)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	possibleAnswers, err := s.GetPossibleWords(state)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	fmt.Println(possibleAnswers)

}
