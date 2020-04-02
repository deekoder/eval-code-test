package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:     "https://167d2c66df3e48da866b47713560a465@sentry.io/5186918",
		Release: "dee-go-server@1.0.0",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("It works!")
	http.HandleFunc("/", HelloServer)

	http.ListenAndServe(":8080", nil)
}

// HelloServer : A simple dummy handler
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	sentry.CaptureMessage("opening file...")
	file, err := os.Open("/tmp/nonexistant-file")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", file)
}
