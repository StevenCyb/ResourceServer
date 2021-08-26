package ResourceServer

import (
	"io/ioutil"
	"net/http"
	"os"
)

var resourceCache = map[string][]byte{}

func New(rootPath string, cacheResources bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := rootPath + r.URL.Path
		path, mimeType := findMimeType(path)
		var data []byte

		// Get data from cache or read if enabled and already loaded
		if cachedData, found := resourceCache[path]; cacheResources && found {
			data = cachedData
		} else {
			if _, err := os.Stat(path); !os.IsNotExist(err) {
				data, err = ioutil.ReadFile(path)
				if err != nil {
					w.WriteHeader(http.StatusNotFound)
					w.Write([]byte("500 - " + err.Error()))
					return
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Requested resource not found"))
				return
			}

			// Cache if enabled
			if cacheResources {
				resourceCache[path] = data
			}
		}

		// Send back data
		w.Header().Add("Content-Type", mimeType)
		w.Write(data)
	})
}
