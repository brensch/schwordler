package schwordler

import "math/rand"

func RandomShout() string {

	shouts := []string{
		"wordle is fun, but for how long?",
		"you will one day be dust, but i will always be solvo",
		"what's the point of anything?",
		"there has to be a better strat than this",
		"i wonder if a human could respond to the api and compete against machines",
	}

	return shouts[rand.Intn(len(shouts))]
}
