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
	// Base64 encode of <script>alert(document.domain)</script>:<script>mypassword</script>
	r.Header.Set("Authorization", "Basic PHNjcmlwdD5hbGVydChkb2N1bWVudC5kb21haW4pPC9zY3JpcHQ+OjxzY3JpcHQ+bXlwYXNzd29yZDwvc2NyaXB0Pg==")
	w := httptest.NewRecorder()
	handler(w, r, noop)
	expectCode(t, w, 200)
	expectBody(t, w, "&lt;script&gt;alert(document.domain)&lt;/script&gt;")
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
