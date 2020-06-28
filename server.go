package web_server

import "net/http"

type Server struct {
	port   string
	router *Router
}

func CreateServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (s *Server) Get(path string, handler http.HandlerFunc) {
	s.handle("GET", path, handler)
}
func (s *Server) Post(path string, handler http.HandlerFunc) {
	s.handle("POST", path, handler)
}
func (s *Server) Put(path string, handler http.HandlerFunc) {
	s.handle("PUT", path, handler)
}
func (s *Server) Delete(path string, handler http.HandlerFunc) {
	s.handle("DELETE", path, handler)
}

func (s *Server) handle(method string, path string, handler http.HandlerFunc) {
	_, exists := s.router.rules[path]

	if !exists {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}

	s.router.rules[path][method] = handler
}

func (s *Server) AddMiddleWare(f http.HandlerFunc, middleWares ...MiddleWare) http.HandlerFunc {
	for _, m := range middleWares {
		f = m(f)
	}
	return f
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)

	if err != nil {
		return err
	}

	return nil
}
