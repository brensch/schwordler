package main

import (
	"fmt"
	"testing"
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
