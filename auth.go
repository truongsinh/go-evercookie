package evercookie

import (
	"bytes"
	"encoding/base64"
	"html"
	"io"
	"net/http"
	"strings"
)

func auth(w http.ResponseWriter, r *http.Request) {
	goto Start
Unauthorized:
	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Set("WWW-Authenticate", "Basic")
	return
Start:
	auth := strings.Split(r.Header.Get("Authorization"), " ")
	if len(auth) < 2 || auth[1] == "" {
		goto Unauthorized
	}
	dec, err := base64.StdEncoding.DecodeString(auth[1])
	if err != nil {
		goto Unauthorized
	}
	b := bytes.Split(dec, []byte(":"))
	if len(b) == 0 {
		goto Unauthorized
	}
	io.WriteString(w, html.EscapeString(string(b[0])))
}
