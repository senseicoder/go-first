package bluesky

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"plcoder.net/namecheck/interfaces"
)

type Bluesky struct {
	Client interfaces.Getter
}

// Variable globale initialisée par init
var re *regexp.Regexp

// Fonction init exécutée automatiquement lors de l'import
func init() {
	fmt.Println("bluesky: init function executed")
	re = regexp.MustCompile(`^[A-Za-z0-9-]{3,39}$`)
}

func (E *Bluesky) String() string {
	return "bluesky"
}

func (E *Bluesky) SetClient(client interfaces.Getter) {
	E.Client = client
}

func (E *Bluesky) IsValid(username string) (bool, error) {
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

func (E *Bluesky) IsAvailable(username string) (bool, error) {
	return true, nil
}

func (E *Bluesky) IsAvailableAPI(username string) (bool, error) {
	url := "https://bsky.social/xrpc/com.atproto.identity.resolveHandle"

	// Construct proper handle format
	handle := username
	if !strings.HasSuffix(handle, ".bsky.social") {
		handle += ".bsky.social"
	}

	// Add handle as query parameter
	url += "?handle=" + handle

	resp, err := E.Client.Get(url)
	if err != nil {
		return false, fmt.Errorf("error checking availability: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return false, nil // Username is taken
	case http.StatusNotFound:
		return true, nil // Username is available
	default:
		return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
}
