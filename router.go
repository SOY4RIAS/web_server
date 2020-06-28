package web_server

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func (r *Router) FindHandler(method string, path string) (http.HandlerFunc, bool, bool) {
	var pathExists, methodExists bool
	var handler http.HandlerFunc

	_, pathExists = r.rules[path]

	handler, methodExists = r.rules[path][method]

	return handler, pathExists, methodExists
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {

	var pathExists, methodExists bool
	var handler http.HandlerFunc

	handler, pathExists, methodExists = r.FindHandler(request.Method, request.URL.Path)

	if !pathExists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodExists {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, request)
}
