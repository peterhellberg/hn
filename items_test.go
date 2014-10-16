package hn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItem_8863(t *testing.T) {
	ts, c := testServerAndClientByFixture("8863")
	defer ts.Close()

	item, err := c.Item(8863)

	assert.Nil(t, err)
	assert.Equal(t, "dhouston", item.By)
	assert.Equal(t, 111, item.Score)
	assert.Equal(t, "http://www.getdropbox.com/u/2/screencast.html", item.URL)
}

func TestItem_8952(t *testing.T) {
	ts, c := testServerAndClientByFixture("8952")
	defer ts.Close()

	item, err := c.Item(8952)

	assert.Nil(t, err)
	assert.Equal(t, "nickb", item.By)
	assert.Equal(t, 1175727286, item.Timestamp)
}

func TestItemTime(t *testing.T) {
	item := Item{Timestamp: 1175727286}

	assert.Equal(t, "2007-04-0", item.Time().Format("2006-01-02"))
}
