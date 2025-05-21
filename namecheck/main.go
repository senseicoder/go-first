package main

import (
	"fmt"
	"os"

	"plcoder.net/namecheck/mypackage"
)

func inc(p *int) {
	*p++
}

func main() {

	// paramètre et validation de l'argument
	if len(os.Args) > 1 {
		firstArg := os.Args[1]
		res, err := mypackage.IsValid(firstArg)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		if res {
			fmt.Println(firstArg)
		}
	} else {
		fmt.Println("Aucun paramètre fourni.")
	}

	// quelques tests
	for _, s := range []string{
		"-cedric",
		"cedric--test",
		"ce",
		"cedriccedriccedriccedriccedriccedriccedriccedriccedriccedriccedriccedriccedriccedric",
		"cedric",
	} {
		res, err := mypackage.IsValid(s)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Println(s, " => ", res)
	}

	i := 42
	p := &i

	// pointers
	fmt.Println(p)  // adresse mémoire
	fmt.Println(*p) // valeur pointée
	*p++
	fmt.Println(*p) // valeur pointée
	inc(p)
	inc(&i)
	fmt.Println(*p)       // valeur pointée
	fmt.Printf("%T\n", p) // type de la valeur pointée
	fmt.Printf("%T\n", i) // type de la variable
}
