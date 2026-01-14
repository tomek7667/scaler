package json

import (
	"fmt"
	"os"
	"sync"

	"github.com/tomek7667/scaler/internal/domain"
)

type Client struct {
	Path string
	db   Db
	m    sync.Mutex
}

type Db struct {
	Scales []domain.Scale `json:"scales"`
}

func New() (*Client, error) {
	c := &Client{
		Path: "./scales.db.json",
		m:    sync.Mutex{},
	}
	if !c.dbExists() {
		err := c.writeDb()
		if err != nil {
			return nil, fmt.Errorf("failed to write default db: %w", err)
		}
	}
	if err := c.readdb(); err != nil {
		return nil, fmt.Errorf("failed to load the database: %w", err)
	}
	return c, nil
}

func (c *Client) dbExists() bool {
	_, err := os.Stat(c.Path)
	return os.IsExist(err) || err == nil
}

func (c *Client) writeDb() error {
	err := os.WriteFile(c.Path, []byte(`{"scales":[]}`), 0o644)
	if err != nil {
		return fmt.Errorf("failed to create default db: %w", err)
	}
	return nil
}
