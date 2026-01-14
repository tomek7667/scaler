package json

import (
	"slices"

	"github.com/google/uuid"
	"github.com/tomek7667/scaler/internal/domain"
)

func (c *Client) SaveScale(scale domain.Scale) {
	c.m.Lock()
	if scale.ID == "" {
		scale.ID = uuid.NewString()
	}

	idx := slices.IndexFunc(c.db.Scales, func(s domain.Scale) bool {
		return s.ID == scale.ID
	})
	if idx == -1 {
		c.db.Scales = append(c.db.Scales, scale)
	} else {
		c.db.Scales[idx] = scale
	}
	go c.autosave()
	c.m.Unlock()
}
