package route

import (
	"net/http"
)

// Ping replies with "pong"
func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{pong}"))
}
