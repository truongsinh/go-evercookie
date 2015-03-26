package evercookie

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthPositive(t *testing.T) {
	r, err := http.NewRequest("GET", "http://smarp.smh.re/evercookie_auth.php", nil)
	if err != nil {
		t.Error(err)
	}
	r.Header.Set("Authorization", "Basic NTg5Og==")
	w := httptest.NewRecorder()
	handler(w, r, noop)
	expectCode(t, w, 200)
	expectBody(t, w, "589")
}

func TestAuthNegative(t *testing.T) {
	r, err := http.NewRequest("GET", "http://smarp.smh.re/evercookie_auth.php", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	handler(w, r, noop)
	expectCode(t, w, 401)
	expectHeader(t, w, "WWW-Authenticate", "Basic")
}
