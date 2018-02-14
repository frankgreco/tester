package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/frankgreco/tester/pkg/apis"
)

// Dump is a middleware that dumps http request details to the reponse
func Dump(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data, err := json.Marshal(apis.RequestDetails{
		Method:  r.Method,
		Headers: reduceMapArray(r.Header),
		Path:    r.URL.EscapedPath(),
		Query:   reduceMapArray(r.URL.Query()),
	})

	if err != nil {
		data = []byte("ERROR")
	}

	w.Write(data)
}

func reduceMapArray(arr map[string][]string) map[string]string {
	toReturn := make(map[string]string, len(arr))
	for k, v := range arr {
		toReturn[k] = strings.Join(v, ",")
	}
	return toReturn
}
