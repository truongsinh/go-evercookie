package evercookie

import (
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
func cache(w http.ResponseWriter, r *http.Request, cacheCookieName string) {
	c, err := r.Cookie(cacheCookieName)
	if err != nil || c.Value == "" {
		w.WriteHeader(304)
		return
	}
	h := w.Header()
	h.Set("Content-Type", "text/html")
	h.Set("Expires", "Tue, 31 Dec 2030 23:30:45 GMT")
	h.Set("Cache-Control", "private, max-age=630720000")

	io.WriteString(w, c.Value)
	return
}
