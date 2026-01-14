package http

import (
	_ "embed"
	"net/http"
)

//go:embed index.html
var indexHTML string

func (s *Server) AddIndexRoute() {
	s.r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(indexHTML))
	})
}
