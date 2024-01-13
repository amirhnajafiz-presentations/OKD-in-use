package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

}
