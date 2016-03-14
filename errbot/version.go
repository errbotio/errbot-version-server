// This is a version server for errbot (errbot.io)
package version

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "3.2.4")
}
