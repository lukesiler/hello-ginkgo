package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/lukesiler/hello-ginkgo/types"
)

// ServeAddress returns address and port that HTTP server listens on
func ServeAddress() string {
	if os.Getenv("PORT") == "" {
		// could also just be ":8080" but using full form
		return "0.0.0.0:8080"
	}

	return "0.0.0.0:" + os.Getenv("PORT")
}

// ServeAPI serves a simple REST API - can be invoked via main entrypoint of process or by test suite in-process
func ServeAPI() (httpServer *http.Server, gracefulStopChan chan os.Signal, err error) {
	logger := log.New(os.Stdout, "", 0)

	gracefulStopChan = make(chan os.Signal, 1)
	signal.Notify(gracefulStopChan, os.Interrupt, syscall.SIGTERM)

	addr := ServeAddress()

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/book", bookHandler)
	mux.HandleFunc("/books", booksHandler)
	httpServer = &http.Server{Addr: addr, Handler: mux}

	go func() {
		logger.Println("Listening on http://", addr)

		if err = httpServer.ListenAndServe(); err != nil {
			logger.Fatal(err)
		}
	}()

	return
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	b := types.Book{Title: "Nothing Special", Author: "Luke Siler", PageCount: 527}
	json, err := json.Marshal(b)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	b1 := types.Book{Title: "Nothing Special", Author: "Luke Siler", PageCount: 527}
	b2 := types.Book{Title: "Special", PageCount: 152}
	coll := types.Books{Items: []types.Book{b1, b2}}
	json, err := json.Marshal(coll)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}
