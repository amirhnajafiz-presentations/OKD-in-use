package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/amirhnajafiz-presentations/snapp-OKD-in-use/project/internal/db"
)

func page(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/index.html")
}

func healthy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handler(w http.ResponseWriter, r *http.Request) {
	req := new(db.Config)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		log.Println(fmt.Errorf("failed to parse request: %w", err))

		return
	}

	log.Println(fmt.Sprintf("resolve request:\n\t%s", req.DNS()))

	err := db.Database{}.Connect(*req)
	if err != nil {
		log.Println(fmt.Errorf("failed to connect to db: %w", err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	port, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		log.Println(fmt.Errorf("failed to read HTTP_PORT: %w", err))

		port = 8080
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/readyz", healthy)
	mux.HandleFunc("/healthz", healthy)

	mux.HandleFunc("/", page)
	mux.HandleFunc("/api/resolve", handler)

	log.Println(fmt.Sprintf("server started on %d ...", port))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		panic(err)
	}
}
