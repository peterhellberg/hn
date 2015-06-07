package hn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	ts, c := testServerAndClientByFixture("peterhellberg")
	defer ts.Close()

	user, err := c.User("peterhellberg")

	assert.Nil(t, err)
	assert.Equal(t, 1300226645, user.Created)
	assert.Equal(t, "2011-03-15", user.CreatedTime().UTC().Format("2006-01-02"))
}

func TestMissingUser(t *testing.T) {
	ts, c := testServerAndClient([]byte(`{}`))
	defer ts.Close()

	_, err := c.User("")
	assert.Equal(t, errMissingID, err)
}
