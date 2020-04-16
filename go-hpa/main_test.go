package main

import "testing"

func TestGreeting(t *testing.T) {
	got := greeting("Code.education Rocks!")
	want := "<b>Code.education Rocks!</b>"

	if got != want {
		t.Errorf("greeting('Code.education Rocks!') got: %v want: %v", got, want)
	}
}

func TestCalcular(t *testing.T) {
	got := calcular()
	want := 2.499957871049515e+11
	if got != want {
		t.Errorf("Problemas ao calcular")
	}
}
