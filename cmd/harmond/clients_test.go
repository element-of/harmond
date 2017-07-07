package main

import "testing"

func validateClientTestClientSet(t *testing.T, c *Client, wantNick, wantUser, wantID string) {
	if c.Nick != wantNick {
		t.Fatalf("expected c.Nick to be %s, got: %v", wantNick, c.Nick)
	}

	if c.User != wantUser {
		t.Fatalf("expected c.User to be %s, got: %v", wantUser, c.User)
	}

	if c.ID != wantID {
		t.Fatalf("expected c.ID to be %s, got: %v", wantID, c.ID)
	}
}

func TestClientSet(t *testing.T) {
	cs := newClientSet()
	c := &Client{
		Nick: "AzureDiamond",
		User: "hunter2",
		ID:   "420AAAAAA",
	}

	ok := cs.Add(c)
	if !ok {
		t.Fatalf("cs.Add failed")
	}

	c1, ok := cs.GetNick(c.Nick)
	if !ok {
		t.Fatalf("cs.GetNick(%q) failed", c.Nick)
	}
	validateClientTestClientSet(t, c1, c.Nick, c.User, c.ID)

	c2, ok := cs.GetID(c.ID)
	if !ok {
		t.Fatalf("cs.GetID(%q) failed", c.ID)
	}
	validateClientTestClientSet(t, c2, c.Nick, c.User, c.ID)

	c.Nick = "SuperHot"
	cs.ChangeNick(c, "AzureDiamond")

	c3, ok := cs.GetNick(c.Nick)
	if !ok {
		t.Fatalf("cs.GetNick(%q) failed", c.Nick)
	}
	validateClientTestClientSet(t, c3, c.Nick, c.User, c.ID)
}
