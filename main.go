package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port = "8080"
)

func init() {

	flag.StringVar(&port, "port", port, "port to listen for games on")

}

func main() {

	flag.Parse()

	s := initStore()

	s.log.Debug("starting")

	http.HandleFunc("/guess", s.HandleDoGuess)
	http.HandleFunc("/results", s.HandleReceiveResults)
	http.HandleFunc("/ping", s.HandleDoPing)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Println(err)
	}

}

// //go:embed words.csv
// var words embed.FS

// type Word struct {
// 	Text      string `json:"text,omitempty"`
// 	Frequency int    `json:"frequency,omitempty"`
// }

// func readWords() ([10][]Word, [10][]Word) {

// 	start := time.Now()

// 	f, err := words.Open("words.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	data, err := csv.NewReader(f).ReadAll()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var allWords, commonWords [10][]Word

// 	for _, row := range data {

// 		frequency, err := strconv.Atoi(row[1])
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}

// 		if len(row[0]) > 10 {
// 			continue
// 		}

// 		w := Word{
// 			Text:      row[0],
// 			Frequency: frequency,
// 		}

// 		allWords[len(row[0])-1] = append(allWords[len(row[0])-1], w)

// 		// TODO: figure out a better method to get this frequency cutoff
// 		if frequency > 3000000 {
// 			commonWords[len(row[0])-1] = append(commonWords[len(row[0])-1], w)

// 		}
// 	}

// 	log.Printf("finished loading words in %s", time.Since(start))

// 	return allWords, commonWords
// }
