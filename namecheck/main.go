package main

import (
	"fmt"
	"net/http"
	"os"

	"plcoder.net/namecheck/bluesky"
	"plcoder.net/namecheck/github"
)

type SocialNetworker interface {
	IsValid(username string) (bool, error)
	IsAvailable(username string) (bool, error)
	Name() string
}

func main() {
	// paramètre et validation de l'argument
	if len(os.Args) > 1 {
		firstArg := os.Args[1]

		github := github.Github{Client: http.DefaultClient}
		bluesky := bluesky.Bluesky{Client: http.DefaultClient}

		for _, network := range []SocialNetworker{&github, &bluesky} {
			res, err := network.IsValid(firstArg)
			if err != nil {
				fmt.Println("Error: ", err)
			}
			if res {
				fmt.Println(network.Name(), ": ", firstArg)
				fmt.Println(network.IsAvailable(firstArg))
			}
		}
	} else {
		fmt.Fprintln(os.Stderr, "Aucun paramètre fourni.")
	}
}
