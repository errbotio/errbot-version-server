// This is a version server for errbot (errbot.io)
package version

import (
	"fmt"
	"github.com/blang/semver"
	"net/http"
	"net/url"
)

const LAST_PY2 = "4.2.2"
const LAST_PY3 = "6.1.2"

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	m, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, "Invalid request.", http.StatusBadRequest)
		return
	}
	if val, ok := m["python"]; ok {
		pyver, err := semver.Make(val[0])
		if err != nil {
			http.Error(w, "Invalid request.", http.StatusBadRequest)
			return
		}

		if pyver.Major == 2 {
			fmt.Fprint(w, LAST_PY2)
			return
		}
		fmt.Fprint(w, LAST_PY3)
		return
	}
	// Old version checker, assume PY2 until the user upgrades to 4.2.2 and we can know its real python version.
	fmt.Fprint(w, LAST_PY2)
}
