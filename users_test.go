package hn

import (
	. "testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *T) {
	ts, c := testServerAndClientByFixture("peterhellberg")
	defer ts.Close()

	user, err := c.User("peterhellberg")

	assert.Nil(t, err)
	assert.Equal(t, 1300226645, user.Created)
}
