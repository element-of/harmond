package main

import "testing"

func TestNickSort(t *testing.T) {
	if l := NickSort("HUNTER2"); l != "hunter2" {
		t.Fatalf("NickSort(\"HUNTER2\") should be \"hunter2\", got: %s", l)
	}
}

func TestServerNextUID(t *testing.T) {
	s := newServer(Config{ServerID: "420"})
	if u := s.NextUID(); u != "420100001" {
		t.Fatalf("UID incrementation failed, expected 420100001, got: %s", u)
	}
}
