package hn

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient()

	if got, want := c.BaseURL.String(), "https://hacker-news.firebaseio.com/v0/"; got != want {
		t.Fatalf(`c.BaseURL.String() = %q, want %q`, got, want)
	}

	if got, want := c.UserAgent, "hn.go/0.0.1"; got != want {
		t.Fatalf(`c.UserAgent = %q, want %q`, got, want)
	}
}

func TestNewRequest(t *testing.T) {
	r, err := NewClient(nil).NewRequest(fmt.Sprintf("foo?bar=%v", 123))

	if err != nil {
		t.Fatalf(`err != nil, got %v`, err)
	}

	if got, want := r.URL.String(), "https://hacker-news.firebaseio.com/v0/foo?bar=123"; got != want {
		t.Fatalf(`r.URL.String() = %q, want %q`, got, want)
	}
}

func testServerAndClient(body []byte) (*httptest.Server, *Client) {
	ts := testServer(body)

	c := DefaultClient
	c.BaseURL, _ = url.Parse(ts.URL)

	return ts, c
}

func testServerAndClientByFixture(fn string) (*httptest.Server, *Client) {
	body, _ := ioutil.ReadFile("_fixtures/" + fn + ".json")

	return testServerAndClient(body)
}

func testServer(body []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write(body)
		}))
}
