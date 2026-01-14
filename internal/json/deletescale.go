package json

import (
	"slices"

	"github.com/tomek7667/scaler/internal/domain"
)

func (c *Client) DeleteScale(id string) {
	c.m.Lock()
	idx := slices.IndexFunc(c.db.Scales, func(s domain.Scale) bool {
		return s.ID == id
	})
	if idx != -1 {
		c.db.Scales = append(c.db.Scales[:idx], c.db.Scales[idx+1:]...)
	}
	go c.autosave()
	c.m.Unlock()
}
