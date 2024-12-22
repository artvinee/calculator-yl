package core

import (
	"encoding/json"
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/api/v1/calculate", func(w http.ResponseWriter, r *http.Request) {
		var req RequestBody
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		result, err := CalculateExpression(req.Expression)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ErrorBody{Error: err.Error()})
			return
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(ResultBody{Result: result})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
