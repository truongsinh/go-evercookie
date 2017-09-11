package evercookie

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCacheWithCookie(t *testing.T) {
	r, err := http.NewRequest("GET", "http://smarp.smh.re/evercookie_cache.php", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.Header.Set("Cookie", "evercookie_cache=<script>alert(document.domain)</script>")
	handler(w, r, noop)
	expectCode(t, w, 200)
	expectHeader(t, w, "Content-Type", "text/html")
	expectHeader(t, w, "Expires", "Tue, 31 Dec 2030 23:30:45 GMT")
	expectHeader(t, w, "Cache-Control", "private, max-age=630720000")
	expectBody(t, w, "&lt;script&gt;alert(document.domain)&lt;/script&gt;")
}

func TestCacheNoCookie(t *testing.T) {
	r, err := http.NewRequest("GET", "http://smarp.smh.re/evercookie_cache.php", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	handler(w, r, noop)
	expectCode(t, w, 304)
}
