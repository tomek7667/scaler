package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/tomek7667/go-http-helpers/h"
	"github.com/tomek7667/scaler/internal/domain"
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

	s.r.With(h.WithRateLimit(h.NewRateLimiter(3, time.Second*15))).Get("/api/scales/{scaleId}", func(w http.ResponseWriter, r *http.Request) {
		scaleID := chi.URLParam(r, "scaleId")
		password := r.URL.Query().Get("password")
		scales := s.dber.GetScales()
		var scale *domain.Scale
		for _, s := range scales {
			if s.ID == scaleID {
				scale = &s
				break
			}
		}
		if scale == nil {
			h.ResNotFound(w, "scale")
			return
		}
		if scale.ScalePassword != password {
			h.ResUnauthorized(w)
			return
		}
		h.ResSuccess(w, scale.Entries)
	})

	s.r.With(h.WithRateLimit(h.NewRateLimiter(3, time.Minute*10))).Post("/api/scales", func(w http.ResponseWriter, r *http.Request) {
		type createScaleDto struct {
			Name     string `json:"name"`
			Password string `json:"password"`
		}
		dto, err := h.GetDto[createScaleDto](r)
		if err != nil {
			h.ResErr(w, err)
			return
		}
		scale := domain.Scale{
			Name:          dto.Name,
			ScalePassword: dto.Password,
			Entries:       []domain.Entry{},
		}

		s.dber.SaveScale(scale)
		h.ResSuccess(w, scale)
	})
}
