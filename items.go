package hn

import (
	"fmt"
	"time"
)

// ItemsService communicates with the news
// related endpoints in the Hacker News API
type ItemsService interface {
	Get(id int) (*Item, error)
}

// itemsService implements ItemsService.
type itemsService struct {
	client *Client
}

// Item represents a item
type Item struct {
	ID        int    `json:"id"`
	Parent    int    `json:"parent"`
	Kids      []int  `json:"kids"`
	Parts     []int  `json:"parts"`
	Score     int    `json:"score"`
	Timestamp int    `json:"time"`
	By        string `json:"by"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	URL       string `json:"url"`
	Dead      bool   `json:"dead"`
	Deleted   bool   `json:"deleted"`
}

func (i *Item) Time() time.Time {
	return time.Unix(int64(i.Timestamp), 0)
}

// Item is a convenience method proxying Items.Get
func (c *Client) Item(id int) (*Item, error) {
	return c.Items.Get(id)
}

// Get retrieves an item with the given id
func (s *itemsService) Get(id int) (*Item, error) {
	req, err := s.client.NewRequest(s.getPath(id))
	if err != nil {
		return nil, err
	}

	var item Item
	_, err = s.client.Do(req, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *itemsService) getPath(id int) string {
	return fmt.Sprintf("item/%v.json", id)
}
