package hn

import "fmt"

// ItemsService communicates with the news
// related endpoints in the Hacker News API
type ItemsService interface {
	Get(id int) (*Item, error)
}

// itemsService implements ItemsService.
type itemsService struct {
	client *Client
}

// ItemID represents a item id
type ItemID int

// Item represents a item
type Item struct {
	ID      int    `json:"id"`
	Parent  int    `json:"parent"`
	Kids    []int  `json:"kids"`
	Parts   []int  `json:"parts"`
	Score   int    `json:"score"`
	Time    int    `json:"time"`
	By      string `json:"by"`
	Type    string `json:"type"`
	Title   string `json:"title"`
	Text    string `json:"text"`
	URL     string `json:"url"`
	Dead    bool   `json:"dead"`
	Deleted bool   `json:"deleted"`
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
