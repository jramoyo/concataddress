package concataddress

import (
	"encoding/json"
	"net/http"

	"github.com/jramoyo/concataddress/internal/pkg/concataddress"
)

// Handler ...
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		var input concataddress.Input
		err := decoder.Decode(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		js, err := json.Marshal(concataddress.Process(input))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	default:
		// Give an error message.
	}
}
