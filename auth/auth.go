package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("x-access-token")

		json.NewEncoder(w).Encode(r)
		header = strings.TrimSpace(header)

		if header == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Missing auth token")
			return
		} else {
			json.NewEncoder(w).Encode(fmt.Sprintf("Token found. Value %s", header))
		}
		next.ServeHTTP(w, r)
	})
}
