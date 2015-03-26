package evercookie

import (
	"net/http"
)

/*
func setCache(etag string, w web.ResponseWriter) {
	const cacheControlKey = "Cache-Control"
	h := w.Header()
	h.Set("Etag", etag)
	if h.Get(cacheControlKey) == "" {
		// only set cache control if there is none before
		// cache it ~ 1 month, as we revision almost all of our assets
		// 3000000 = 3e6s = 1.1408 month
		h.Set(cacheControlKey, "public, max-age=3000000")
	}
	return
}*/
type Config struct {
	EtagCookieName  string
	CacheCookieName string
	AuthPath        string
	PngPath         string
	EtagPath        string
	CachePath       string
}

func DefaultConfig() *Config {
	return &Config{
		EtagCookieName:  "evercookie_etag",
		CacheCookieName: "evercookie_cache",
		AuthPath:        "/evercookie_auth.php",
		PngPath:         "/evercookie_png.php",
		EtagPath:        "/evercookie_etag.php",
		CachePath:       "/evercookie_cache.php",
	}
}

/**
 * Port to Golang by TruongSinh <i@truongsinh.pro>
 * Original work by samy kamkar, https://github.com/samyk/evercookie/
 *
 */
func Evercookie(config *Config) func(http.ResponseWriter, *http.Request, func()) {
	return func(w http.ResponseWriter, r *http.Request, next func()) {
		switch r.URL.Path {
		case config.AuthPath:
			auth(w, r)
			return
		case config.EtagPath:
			etag(w, r, config.EtagCookieName)
			return
		case config.CachePath:
			cache(w, r, config.CacheCookieName)
			return
		case config.PngPath:
			/**
			 * PNG cache is technically no difference than text cache, but adds much more overhead to encode and decode.
			 * Return 201 No content
			 */
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next()
	}
}
