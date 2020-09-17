package http

import (
	"fmt"
	"net/http"
	"os"
)

// GetReturns2xx tests performs a health check against an URL
func GetReturns2xx(URL string) bool {
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http error: %v\n", err)
		return false
	}
	return resp.StatusCode <= 200 && resp.StatusCode < 299
}
