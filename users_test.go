package hn

import "testing"

func TestUser(t *testing.T) {
	ts, c := testServerAndClientByFixture("peterhellberg")
	defer ts.Close()

	user, err := c.User("peterhellberg")

	if err != nil {
		t.Fatalf(`err != nil, got %v`, err)
	}

	if got, want := user.Created, 1300226645; got != want {
		t.Fatalf(`user.Created = %d, want %d`, got, want)
	}

	if got, want := user.CreatedTime().UTC().Format("2006-01-02"), "2011-03-15"; got != want {
		t.Fatalf(`user.CreatedTime().UTC().Format("2006-01-02") = %q, want %q`, got, want)
	}
}

func TestMissingUser(t *testing.T) {
	ts, c := testServerAndClient([]byte(`{}`))
	defer ts.Close()

	if _, err := c.User(""); err != errMissingID {
		t.Fatalf(`err != errMissingID, got %v`, err)
	}
}
