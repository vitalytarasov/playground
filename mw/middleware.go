package mw

import (
	"fmt"
	"net/http"
)

// Example is an example of HTTP middleware
func Example(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("i am middleware")
		next(w, r)
	}
}
