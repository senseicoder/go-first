package main

import (
	"fmt"
	"net/http"
	"os"

	"plcoder.net/namecheck/bluesky"
	"plcoder.net/namecheck/github"
)

func main() {
	// paramètre et validation de l'argument
	if len(os.Args) > 1 {
		github := github.Github{Client: http.DefaultClient}
		bluesky := bluesky.Bluesky{Client: http.DefaultClient}

		firstArg := os.Args[1]
		res, err := github.IsValid(firstArg)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		if res {
			fmt.Println("github: ", firstArg)
			fmt.Println(github.IsAvailable(firstArg))
		}

		res2, err2 := bluesky.IsValid(firstArg)
		if err2 != nil {
			fmt.Println("Error: ", err2)
		}
		if res2 {
			fmt.Println("bluesky: ", firstArg)
			fmt.Println(bluesky.IsAvailable(firstArg))
		}
	} else {
		fmt.Fprintln(os.Stderr, "Aucun paramètre fourni.")
	}
}
