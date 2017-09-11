package evercookie

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var handler func(http.ResponseWriter, *http.Request, func())
var noop = func() {}

func init() {
	handler = Evercookie(DefaultConfig())
}

func expectHeader(t *testing.T, w *httptest.ResponseRecorder, header, expected string) {
	value := w.Header().Get(header)
	if value != expected {
		t.Errorf("Header %q not %q but %q", header, expected, value)
	}
}
func expectCode(t *testing.T, w *httptest.ResponseRecorder, expected int) {
	value := w.Code
	if value != expected {
		t.Errorf("Code not %d but %d", expected, value)
	}
}
func expectBody(t *testing.T, w *httptest.ResponseRecorder, expected string) {
	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error(err)
	}
	if string(b) != expected {
		t.Errorf("Body not %q but %q", expected, b)
	}
}

func TestPng(t *testing.T) {
	r, err := http.NewRequest("GET", "http://smarp.smh.re/evercookie_png.php", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	handler(w, r, noop)
	expectCode(t, w, http.StatusNoContent)
}

func TestNext(t *testing.T) {
	r, err := http.NewRequest("GET", "http://smarp.smh.re/somethingelse", nil)
	isCalled := false
	f := func() {
		isCalled = true
	}
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	handler(w, r, f)
	if isCalled == false {
		t.Error("Next is not called")
	}
}
