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
	http.ServeFile(w, r, "/web/index.html")
}

func handler(w http.ResponseWriter, r *http.Request) {
	req := new(db.Config)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		log.Println(fmt.Errorf("failed to parse request: %w", err))

		return
	}

	err := db.Database{}.Connect(*req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	port, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))

	mux := http.NewServeMux()

	mux.HandleFunc("/", page)
	mux.HandleFunc("/api/resolve", handler)

	log.Println(fmt.Sprintf("app server started on %d ...", port))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		panic(err)
	}
}
