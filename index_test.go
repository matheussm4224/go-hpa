package main

import(
	"testing"
)

func TestMessage(t *testing.T) {
	message := Load()

	if message != "Code.education Rocks" {
		t.Errorf("Error ao carregar");
	}
}