package route

import (
	"fmt"
	"net/http"
)

// Ping replies with "pong"
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println("i am route ping")
	w.Write([]byte("{pong}"))
}
