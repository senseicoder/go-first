package main

import (
	"fmt"
	"os"

	"plcoder.net/namecheck/github"
)

func main() {
	// paramètre et validation de l'argument
	if len(os.Args) > 1 {
		firstArg := os.Args[1]
		res, err := github.IsValid(firstArg)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		if res {
			fmt.Println(firstArg)
			fmt.Println(github.IsAvailable(firstArg))
		}
	} else {
		fmt.Println("Aucun paramètre fourni.")
	}
}
