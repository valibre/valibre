package main

import "testing"

func TestTest(t *testing.T) {
	if 1 != 1 {
		t.Errorf("Universe has failed: 1 != 1")
        }
}