package hn

// LiveService communicates with the news
// related endpoints in the Hacker News API
type LiveService interface {
	TopStories() ([]int, error)
	MaxItem() (int, error)
	Updates() (*Updates, error)
}

// liveService implements LiveService.
type liveService struct {
	client *Client
}

// Updates contains the latest updated items and profiles
type Updates struct {
	Items    []int    `json:"items"`
	Profiles []string `json:"profiles"`
}

// TopStories is a convenience method proxying Live.TopStories
func (c *Client) TopStories() ([]int, error) {
	return c.Live.TopStories()
}

// TopStories retrieves the current top stories
func (s *liveService) TopStories() ([]int, error) {
	req, err := s.client.NewRequest(s.topStoriesPath())
	if err != nil {
		return nil, err
	}

	var value []int
	_, err = s.client.Do(req, &value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (s *liveService) topStoriesPath() string {
	return "topstories.json"
}

// MaxItem is a convenience method proxying Live.MaxItem
func (c *Client) MaxItem() (int, error) {
	return c.Live.MaxItem()
}

// MaxItem retrieves the current largest item id
func (s *liveService) MaxItem() (int, error) {
	req, err := s.client.NewRequest(s.maxItemPath())
	if err != nil {
		return 0, err
	}

	var value int
	_, err = s.client.Do(req, &value)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func (s *liveService) maxItemPath() string {
	return "maxitem.json"
}

// Updates is a convenience method proxying Live.Updates
func (c *Client) Updates() (*Updates, error) {
	return c.Live.Updates()
}

// Updates retrieves the current largest item id
func (s *liveService) Updates() (*Updates, error) {
	req, err := s.client.NewRequest(s.updatesPath())
	if err != nil {
		return nil, err
	}

	var value Updates
	_, err = s.client.Do(req, &value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func (s *liveService) updatesPath() string {
	return "updates.json"
}
