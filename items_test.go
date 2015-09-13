package hn

import "testing"

func TestItem_8863(t *testing.T) {
	ts, c := testServerAndClientByFixture("8863")
	defer ts.Close()

	item, err := c.Item(8863)

	if err != nil {
		t.Fatalf(`err != nil, got %v`, err)
	}

	if got, want := item.By, "dhouston"; got != want {
		t.Fatalf(`item.By = %q, want %q`, got, want)
	}

	if got, want := item.Score, 111; got != want {
		t.Fatalf(`item.Score = %d, want %d`, got, want)
	}

	if got, want := item.URL, "http://www.getdropbox.com/u/2/screencast.html"; got != want {
		t.Fatalf(`item.URL = %q, want %q`, got, want)
	}
}

func TestItem_8952(t *testing.T) {
	ts, c := testServerAndClientByFixture("8952")
	defer ts.Close()

	item, err := c.Item(8952)

	if err != nil {
		t.Fatalf(`err != nil, got %v`, err)
	}

	if got, want := item.By, "nickb"; got != want {
		t.Fatalf(`item.By = %q, want %q`, got, want)
	}

	if got, want := item.Timestamp, 1175727286; got != want {
		t.Fatalf(`item.Timestamp = %d, want %d`, got, want)
	}
}

func TestItemTime(t *testing.T) {
	item := Item{Timestamp: 1175727286}

	if got, want := item.Time().UTC().Format("2006-01-02"), "2007-04-04"; got != want {
		t.Fatalf(`item.Time().UTC().Format("2006-01-02") = %q, want %q`, got, want)
	}
}
