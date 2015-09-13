package hn

import "testing"

func TestTopStories(t *testing.T) {
	ts, c := testServerAndClientByFixture("topstories")
	defer ts.Close()

	top, err := c.TopStories()

	if err != nil {
		t.Fatalf(`err != nil, got %v`, err)
	}

	if got, want := len(top), 100; got != want {
		t.Fatalf(`len(top) = %d, want %d`, got, want)
	}
}

func TestMaxItem(t *testing.T) {
	ts, c := testServerAndClientByFixture("maxitem")
	defer ts.Close()

	item, err := c.MaxItem()

	if err != nil {
		t.Fatalf(`err != nil, got %v`, err)
	}

	if got, want := item, 8424452; got != want {
		t.Fatalf(`item = %d, want %d`, got, want)
	}
}

func TestUpdates(t *testing.T) {
	ts, c := testServerAndClientByFixture("updates")
	defer ts.Close()

	updates, err := c.Updates()

	if err != nil {
		t.Fatalf(`err != nil, got %v`, err)
	}

	if got, want := updates.Profiles[5], "benologist"; got != want {
		t.Fatalf(`updates.Profiles[5] = %q, want %q`, got, want)
	}

	if got, want := updates.Items[7], 8423650; got != want {
		t.Fatalf(`updates.Items[7] = %d, want %d`, got, want)
	}
}
