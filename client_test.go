package hn

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c := NewClient(nil)

	assert.Equal(t, "https://hacker-news.firebaseio.com/v0/", c.BaseURL.String())
	assert.Equal(t, "hn.go/0.0.1", c.UserAgent)
}

func TestNewRequest(t *testing.T) {
	r, err := NewClient(nil).NewRequest(fmt.Sprintf("foo?bar=%v", 123))

	assert.Nil(t, err)
	assert.Equal(t, "https://hacker-news.firebaseio.com/v0/foo?bar=123", r.URL.String())
}

func testServerAndClientByFixture(fn string) (*httptest.Server, *Client) {
	body, _ := ioutil.ReadFile("_fixtures/" + fn + ".json")

	ts := testServer(body)

	c := NewClient(nil)
	c.BaseURL, _ = url.Parse(ts.URL)

	return ts, c
}

func testServer(body []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write(body)
		}))
}
