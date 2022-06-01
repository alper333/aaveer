package rate

import (
	"net/http"

	"github.com/conflux-chain/conflux-infura/util"
)

func HttpHandler(registry *Registry, next http.Handler) http.Handler {
	if registry == nil {
		return next
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, ok := util.GetIPAddress(r.Context())
		if !ok {
			next.ServeHTTP(w, r)
			return
		}

		limiter, ok := registry.Get("rpc.httpRequest")
		if !ok {
			next.ServeHTTP(w, r)
		} else if limiter.Allow(ip, 1) {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
		}
	})
}
