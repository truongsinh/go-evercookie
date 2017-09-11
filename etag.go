package evercookie

import (
	"html"
	"io"
	"net/http"
)

/**
 * Port to NodeJS by TruongSinh <i@truongsinh.pro>
 * Defined by samy kamkar, https://github.com/samyk/evercookie/blob/master/evercookie_etag.php
 *
 * This is the server-side ETag software which tags a user by
 * using the Etag HTTP header, as well as If-None-Match to check
 * if the user has been tagged before.
 */
func etag(w http.ResponseWriter, r *http.Request, etagCookieName string) {
	var ecValue string
	c, err := r.Cookie(etagCookieName)
	if err != nil || c.Value == "" {
		ecValue = r.Header.Get("If-None-Match")
		w.Header().Set("Etag", html.EscapeString(ecValue))
		w.WriteHeader(304)
		return
	}
	ecValue = c.Value
	if ecValue != "" {
		w.Header().Set("Etag", html.EscapeString(ecValue))
		io.WriteString(w, html.EscapeString(ecValue))
		return
	}
	w.WriteHeader(304)
	return
}
