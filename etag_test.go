package evercookie

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

/**
"with If-None-Match header responds id both in Etag header and cookie, favors cookie value"
*/
func TestEtagVsCookie(t *testing.T) {
	r, err := http.NewRequest("GET", "http://smarp.smh.re/evercookie_etag.php", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.Header.Set("Cookie", "evercookie_etag=111")
	r.Header.Set("If-None-Match", "112")
	handler(w, r, noop)
	expectCode(t, w, 200)
	expectHeader(t, w, "Etag", "111")
	expectBody(t, w, "111")
}

func TestEtagAndCookie(t *testing.T) {
	r, err := http.NewRequest("GET", "http://smarp.smh.re/evercookie_etag.php", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.Header.Set("Cookie", "evercookie_etag=121")
	r.Header.Set("If-None-Match", "121")
	handler(w, r, noop)
	expectCode(t, w, 200)
	expectHeader(t, w, "Etag", "121")
	expectBody(t, w, "121")
}

func TestEtagNoCookie(t *testing.T) {
	r, err := http.NewRequest("GET", "http://smarp.smh.re/evercookie_etag.php", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.Header.Set("If-None-Match", "131")
	handler(w, r, noop)
	expectCode(t, w, 304)
}

func TestEtagOnlyCookie(t *testing.T) {
	r, err := http.NewRequest("GET", "http://smarp.smh.re/evercookie_etag.php", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.Header.Set("Cookie", "evercookie_etag=141")
	handler(w, r, noop)
	expectCode(t, w, 200)
	expectHeader(t, w, "Etag", "141")
	expectBody(t, w, "141")
}

func TestEtagNothing(t *testing.T) {
	r, err := http.NewRequest("GET", "http://smarp.smh.re/evercookie_etag.php", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	handler(w, r, noop)
	expectCode(t, w, 304)
}
