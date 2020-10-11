package main

import (
	"fmt"
	"net/http"
	"playground/route"
)

func main() {

	// register routes
	http.HandleFunc("/ping", route.Ping)

	// init and start server
	srv := &http.Server{Addr: ":1080"}
	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("cannot start server (%v)", err)
	}
	fmt.Print("exiting")
}
