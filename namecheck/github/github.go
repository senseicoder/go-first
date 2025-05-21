package github

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Variable globale initialisée par init
var re *regexp.Regexp

// Fonction init exécutée automatiquement lors de l'import
func init() {
	fmt.Println("github: init function executed")
	re = regexp.MustCompile(`^[A-Za-z0-9-]{3,39}$`)
}

func IsValid(username string) (bool, error) {
	var error error
	res := true

	if strings.HasPrefix(username, "-") {
		res = false
		error = errors.New("username cannot start with a hyphen")
	}

	if strings.Contains(username, "--") {
		res = false
		error = errors.New("username cannot contain consecutive hyphens")
	}

	reresult := re.MatchString(username)
	if !reresult {
		res = false
		error = errors.New("username must be between 3 and 39 characters long and can only contain letters, numbers, and hyphens")
	}

	return res, error
}
