package pt03

import (
	"testing"
)

func TestState_PutInBoatShouldPutPersonInBoat(t *testing.T) {
	expected := `[ 🦊 🐔 🌾   ---\_\___🧑_/_________/---         ]`

	state := NewState()

	state = PutInBoat(state, ActorPerson)

	actual := ViewState(state)

	if actual != expected {
		t.Errorf("\nexpected: %s\nactual:   %s", expected, actual)
	}
}

func TestState_PutInBoatShouldPutHenInBoat(t *testing.T) {
	expected := `[ 🦊   🌾 🧑 ---\_\_🐔___/_________/---         ]`

	state := NewState()

	state = PutInBoat(state, ActorHen)

	actual := ViewState(state)

	if actual != expected {
		t.Errorf("\nexpected: %s\nactual:   %s", expected, actual)
	}
}

func TestState_PutOnLandShouldPutPersonOnLand(t *testing.T) {
	expected := `[ 🦊 🐔 🌾 🧑 ---\_\_____/_________/---         ]`

	state := NewState()

	state = PutInBoat(state, ActorPerson)
	state = PutOnLand(state, ActorPerson)

	actual := ViewState(state)

	if actual != expected {
		t.Errorf("\nexpected: %s\nactual:   %s", expected, actual)
	}
}

func TestState_PutOnLandShouldPutHenOnLand(t *testing.T) {
	expected := `[ 🦊 🐔 🌾 🧑 ---\_\_____/_________/---         ]`

	state := NewState()

	state = PutInBoat(state, ActorHen)
	state = PutOnLand(state, ActorHen)

	actual := ViewState(state)

	if actual != expected {
		t.Errorf("\nexpected: %s\nactual:   %s", expected, actual)
	}
}

func TestState_CrossRiverShouldMoveBoatToOtherBank(t *testing.T) {
	expected := `[ 🦊 🐔 🌾   ---\_________\___🧑_/_/---         ]`

	state := NewState()

	state = PutInBoat(state, ActorPerson)
	state = CrossRiver(state)

	actual := ViewState(state)

	if actual != expected {
		t.Errorf("\nexpected: %s\nactual:   %s", expected, actual)
	}
}
