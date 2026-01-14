package json

import "github.com/tomek7667/scaler/internal/domain"

func (c *Client) GetScales() []domain.Scale {
	return c.db.Scales
}
