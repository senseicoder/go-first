package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func validateUsername(input string) bool {
	if strings.HasPrefix(input, "-") {
		log.Fatal("Invalid username format (begin by hyphen)")
		return false
	}

	if strings.Contains(input, "--") {
		log.Fatal("Invalid username format (double hyphen)")
		return false
	}

	if re := regexp.MustCompile(`^[A-Za-z0-9-]{3,39}$`); !re.MatchString(input) {
		log.Fatal("Invalid username format")
		return false
	}

	return true
}

func main() {
	if len(os.Args) > 1 {
		firstArg := os.Args[1]
		if validateUsername(firstArg) {
			fmt.Println(firstArg)
		}
	} else {
		fmt.Println("Aucun paramètre fourni.")
	}

}
