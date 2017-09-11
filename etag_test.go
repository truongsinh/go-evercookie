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
	r.Header.Set("Cookie", "evercookie_etag=<script>111</script>")
	r.Header.Set("If-None-Match", "<script>112</script>")
	handler(w, r, noop)
	expectCode(t, w, 200)
	expectHeader(t, w, "Etag", "&lt;script&gt;111&lt;/script&gt;")
	expectBody(t, w, "&lt;script&gt;111&lt;/script&gt;")
}

func TestEtagAndCookie(t *testing.T) {
	r, err := http.NewRequest("GET", "http://smarp.smh.re/evercookie_etag.php", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.Header.Set("Cookie", "evercookie_etag=<script>121</script>")
	r.Header.Set("If-None-Match", "<script>121</script>")
	handler(w, r, noop)
	expectCode(t, w, 200)
	expectHeader(t, w, "Etag", "&lt;script&gt;121&lt;/script&gt;")
	expectBody(t, w, "&lt;script&gt;121&lt;/script&gt;")
}

func TestEtagNoCookie(t *testing.T) {
	r, err := http.NewRequest("GET", "http://smarp.smh.re/evercookie_etag.php", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.Header.Set("If-None-Match", "<script>131</script>")
	handler(w, r, noop)
	expectCode(t, w, 304)
	expectHeader(t, w, "Etag", "&lt;script&gt;131&lt;/script&gt;")
}

func TestEtagOnlyCookie(t *testing.T) {
	r, err := http.NewRequest("GET", "http://smarp.smh.re/evercookie_etag.php", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.Header.Set("Cookie", "evercookie_etag=<script>141</script>")
	handler(w, r, noop)
	expectCode(t, w, 200)
	expectHeader(t, w, "Etag", "&lt;script&gt;141&lt;/script&gt;")
	expectBody(t, w, "&lt;script&gt;141&lt;/script&gt;")
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
