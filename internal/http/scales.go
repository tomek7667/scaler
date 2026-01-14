package http

import (
	"net/http"

	"github.com/tomek7667/go-http-helpers/h"
)

func (s *Server) AddScalesRoute() {
	s.r.Get("/api/scales", func(w http.ResponseWriter, r *http.Request) {
		_scales := s.dber.GetScales()

		type scale struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Entries int    `json:"entries"`
		}

		scales := []scale{}
		for _, s := range _scales {
			scales = append(scales, scale{
				ID:      s.ID,
				Name:    s.Name,
				Entries: len(s.Entries),
			})
		}
		h.ResSuccess(w, scales)
	})
}
