package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/jub0bs/cors"
	"plcoder.net/namecheck/bluesky"
	"plcoder.net/namecheck/github"
	"plcoder.net/namecheck/interfaces"
)

type Response struct {
	Username string              `json:"username"`
	Results  []interfaces.Result `json:"results,omitempty"`
}

func handleCheck(w http.ResponseWriter, r *http.Request) {
	// Parse the username from query parameters
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	resultCh := make(chan interfaces.Result)

	networks := make([]interfaces.SocialNetworker, 0, 40) // Pré-allouer pour 40 éléments

	// Créer 20 instances de Github
	for i := 0; i < 20; i++ {
		networks = append(networks, &github.Github{Client: http.DefaultClient})
	}

	// Créer 20 instances de Bluesky
	for i := 0; i < 20; i++ {
		networks = append(networks, &bluesky.Bluesky{Client: http.DefaultClient})
	}

	var wg sync.WaitGroup
	for _, network := range networks {
		wg.Add(1)
		go interfaces.ExecTasks(network, username, &wg, resultCh)
	}
	go func() {
		wg.Wait()
		close(resultCh) // Fermer le canal après que toutes les goroutines aient terminé
	}()

	w.Header().Set("Content-Type", "application/json")
	var sliceReponse = make([]interfaces.Result, 0, len(networks))

	for result := range resultCh {
		sliceReponse = append(sliceReponse, result)
	}

	var Response = Response{
		Username: username,
		Results:  sliceReponse,
	}

	// Encode the slice into JSON and send it as the response
	if err := json.NewEncoder(w).Encode(Response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/check", handleCheck)
	corsMw, err := cors.NewMiddleware(cors.Config{
		Origins: []string{"https://namecheck.jub0bs.dev"},
	})
	if err != nil {
		log.Fatal(err)
	}

	// apply your CORS middleware to your HTTP request multiplexer
	handler := corsMw.Wrap(mux)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":8080", nil)
}
