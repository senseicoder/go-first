package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"plcoder.net/namecheck/bluesky"
	"plcoder.net/namecheck/github"
	"plcoder.net/namecheck/interfaces"
)

func main() {
	if len(os.Args) > 1 {
		firstArg := os.Args[1]

		resultCh := make(chan interfaces.Result)

		networks := make([]interfaces.SocialNetworker, 0, 40) // Pré-allouer pour 40 éléments

		// Créer 20 instances de Github
		for range 20 {
			networks = append(networks, &github.Github{Client: http.DefaultClient})
		}

		// Créer 20 instances de Bluesky
		for range 20 {
			networks = append(networks, &bluesky.Bluesky{Client: http.DefaultClient})
		}

		var wg sync.WaitGroup
		for _, network := range networks {
			wg.Add(1)
			go interfaces.ExecTasks(network, firstArg, &wg, resultCh)
		}
		go func() {
			wg.Wait()
			close(resultCh) // Fermer le canal après que toutes les goroutines aient terminé
		}()

		nb := 0
		for result := range resultCh {
			fmt.Println("agrégé ", nb, ": ", result.Platform, ": ", result.Valid, " : ", result.Available)
			nb++
		}

		wg.Wait()
	} else {
		fmt.Fprintln(os.Stderr, "Aucun paramètre fourni.")
	}
}
