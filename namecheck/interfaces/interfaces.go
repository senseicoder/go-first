package interfaces

import (
	"fmt"
	"net/http"
	"sync"
)

type Getter interface {
	Get(url string) (*http.Response, error)
}

type Result struct {
	Platform  string `json:"platform"`
	Valid     bool   `json:"valid"`
	Available bool   `json:"available"`
	Err       error  `json:"error,omitempty"`
}

type SocialNetworker interface {
	IsValid(username string) (bool, error)
	IsAvailable(username string) (bool, error)
	SetClient(client Getter)
	//String() string peut être remplacée par l'import d'une interface
	fmt.Stringer
}

func ExecTasks(network SocialNetworker, username string, wg *sync.WaitGroup, resultCh chan<- Result) {
	defer wg.Done()

	res, err := network.IsValid(username)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	if res {
		if available, err := network.IsAvailable(username); err == nil {
			//fmt.Println(network, ": ", username, " : ", available)
			resultCh <- Result{Platform: network.String(), Valid: res, Available: available, Err: err}
		}
	}
}
