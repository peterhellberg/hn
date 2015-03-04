package hn

import (
	. "testing"

	"github.com/stretchr/testify/assert"
)

func TestTopStories(t *T) {
	ts, c := testServerAndClientByFixture("topstories")
	defer ts.Close()

	top, err := c.TopStories()

	assert.Nil(t, err)
	assert.Equal(t, 100, len(top))
}

func TestMaxItem(t *T) {
	ts, c := testServerAndClientByFixture("maxitem")
	defer ts.Close()

	item, err := c.MaxItem()

	assert.Nil(t, err)
	assert.Equal(t, 8424452, item)
}

func TestUpdates(t *T) {
	ts, c := testServerAndClientByFixture("updates")
	defer ts.Close()

	updates, err := c.Updates()

	assert.Nil(t, err)
	assert.Equal(t, "benologist", updates.Profiles[5])
	assert.Equal(t, 8423650, updates.Items[7])
}
