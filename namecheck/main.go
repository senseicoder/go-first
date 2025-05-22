package main

import (
	"fmt"
	"net/http"
	"os"

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

		for _, network := range networks {
			res, err := network.IsValid(firstArg)
			if err != nil {
				fmt.Println("Error: ", err)
			}
			if res {
				fmt.Println(network, ": ", firstArg)
				fmt.Println(network.IsAvailable(firstArg))
			}
		}
	} else {
		fmt.Fprintln(os.Stderr, "Aucun paramètre fourni.")
	}
}
