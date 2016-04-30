package barnard

import (
	"fmt"
	"net/http"
)

func initializeApiMethods() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	return mux
}
