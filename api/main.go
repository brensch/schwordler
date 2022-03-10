package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/brensch/schwordler"
	"github.com/sirupsen/logrus"
)

const (
	// cloud run specific env vars
	// the vendor lock in is real
	EnvVarService       = "K_SERVICE"
	EnvVarRevision      = "K_REVISION"
	EnvVarConfiguration = "K_CONFIGURATION"
	EnvVarPort          = "PORT"

	defaultPort = "8080"
)

type api struct {
	s *schwordler.Store
}

func main() {

	service := os.Getenv(EnvVarService)
	revision := os.Getenv(EnvVarRevision)
	configuration := os.Getenv(EnvVarConfiguration)
	port := os.Getenv(EnvVarPort)
	if port == "" {
		port = defaultPort
	}
	// EnvVarService should always be set when running in a cloud run instance
	onCloud := service != ""

	s := schwordler.InitStore()

	s.Log.WithFields(logrus.Fields{
		"revision":      revision,
		"configuration": configuration,
		"harambe":       onCloud,
	}).Info("starting")

	a := &api{
		s: s,
	}

	http.HandleFunc("/guess", a.HandleDoGuess)
	http.HandleFunc("/results", a.HandleReceiveResults)
	http.HandleFunc("/ping", a.HandleDoPing)

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
