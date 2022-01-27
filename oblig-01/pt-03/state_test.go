package pt03

import (
	"testing"
)

func TestState_ViewShouldReturnCorrectViewForInitialState(t *testing.T) {
	expected := `[ 🦊 🐔 🌾 🧑 ---\_\_____/_________/---         ]`

	state := NewState()
	actual := ViewState(state)

	if actual != expected {
		t.Errorf("\nexpected: %s\nactual:   %s", expected, actual)
	}
}
