package github

import (
	"errors"
	"fmt"
	"net/http"
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

// todo injecter le client http pour les TU
func IsAvailable(username string) (bool, error) {
	url := "https://github.com/" + username
	resp, err := http.Get(url)
	if err != nil {
		return false, errors.New("unattended error")
	}
	defer resp.Body.Close() //important pour fermer la connexion en fin de fonction
	//fmt.Println(url, "=>", resp.StatusCode)

	switch resp.StatusCode {
	case http.StatusOK:
		return false, nil
	case http.StatusNotFound:
		return true, nil
	default:
		return false, errors.New("unattended error")
	}
}
