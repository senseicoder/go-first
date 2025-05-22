package interfaces

import "net/http"

type Getter interface {
	Get(url string) (*http.Response, error)
}
