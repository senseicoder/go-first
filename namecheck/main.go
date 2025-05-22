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

type SocialNetworker interface {
	IsValid(username string) (bool, error)
	IsAvailable(username string) (bool, error)
	SetClient(client interfaces.Getter)
	//String() string
	fmt.Stringer
}

func execTasks(network SocialNetworker, username string, wg *sync.WaitGroup) {
	defer wg.Done()

	res, err := network.IsValid(username)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	if res {
		if available, err := network.IsAvailable(username); err == nil {
			fmt.Println(network, ": ", username, " : ", available)
		}
	}
}

func main() {
	if len(os.Args) > 1 {
		firstArg := os.Args[1]

		networks := make([]SocialNetworker, 0, 40) // Pré-allouer pour 40 éléments

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
			go execTasks(network, firstArg, &wg)
		}

		wg.Wait()
	} else {
		fmt.Fprintln(os.Stderr, "Aucun paramètre fourni.")
	}
}
