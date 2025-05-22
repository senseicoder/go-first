package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"

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

		for _, t := range []any{
			(*github.Github)(nil),
			(*bluesky.Bluesky)(nil),
		} {
			network := reflect.New(reflect.TypeOf(t).Elem()).Interface().(SocialNetworker)
			network.SetClient(http.DefaultClient)

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
