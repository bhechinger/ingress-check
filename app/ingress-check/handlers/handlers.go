// Package handlers contains the full set of handler functions and routes
// supported by the web api.
package handlers

import (
	"net/http"

	"github.com/dimfeld/httptreemux"
)

// API constructs an http.Handler with all application routes defined.
func API() *httptreemux.ContextMux {
	mux := httptreemux.NewContextMux()
	mux.Handle(http.MethodGet, "/test", readiness)
	return mux
}
