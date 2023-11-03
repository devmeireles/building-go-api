// main.go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Data struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		id := strings.TrimPrefix(r.URL.Path, "/")

		if r.Method == http.MethodGet {
			if len(id) > 0 {
				fmt.Fprintf(w, "GET request for ID: %s", id)
			} else {
				fmt.Fprintf(w, "GET request")
			}
		} else if r.Method == http.MethodPost {
			var data Data
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&data)
			if err != nil {
				http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
				return
			}

			fmt.Fprintf(w, "POST request for: %s", data.Message)
		} else if r.Method == http.MethodPatch {
			var data Data
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&data)
			if err != nil {
				http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
				return
			}

			fmt.Fprintf(w, "PATCH request for ID: %s with body: %s", id, data.Message)
		} else if r.Method == http.MethodDelete {
			fmt.Fprintf(w, "DELETE request for ID: %s", id)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
