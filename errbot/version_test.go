package version

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var versionTests = []struct {
	request  string
	expected string
}{
	{"http://version.errbot.io/", LAST_PY2},
	{"http://version.errbot.io/?4.1.3", LAST_PY2},
	{"http://version.errbot.io/?python=2.7.1&errbot=4.0.0", LAST_PY2},
	{"http://version.errbot.io/?python=3.4.3&errbot=4.3.0", LAST_PY3},
	{"http://version.errbot.io/?python=3.4.3&errbot=4.3.0", LAST_PY3},
}

func TestDefaultVersion(t *testing.T) {
	for _, vt := range versionTests {
		req := httptest.NewRequest("GET", vt.request, nil)
		w := httptest.NewRecorder()
		handler(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Home page didn't return %v", http.StatusOK)
		}
		actual := w.Body.String()
		if actual != vt.expected {
			t.Errorf("Version server should have returned %s for %s but %s", actual, vt.expected, vt.request)
		}
	}
}
