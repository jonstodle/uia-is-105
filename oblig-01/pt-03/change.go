package pt03

func PutInBoat(state State, actorType ActorType) State {
	for _, actor := range state {
		if actor.Type == actorType {
			if actor.Location == LocationWestLand {
				actor.Location = LocationWestBoat
			} else if actor.Location == LocationEastLand {
				actor.Location = LocationEastBoat
			}
			break
		}
	}

	return state
}

func PutOnLand(state State, actorType ActorType) State {
	for _, actor := range state {
		if actor.Type == actorType {
			if actor.Location == LocationWestBoat {
				actor.Location = LocationWestLand
			} else if actor.Location == LocationEastBoat {
				actor.Location = LocationEastLand
			}
			break
		}
	}

	return state
}

func CrossRiver(state State) State {
	if state.anyInLocation(LocationWestBoat) {
		for _, actor := range state {
			if actor.Location == LocationWestBoat {
				actor.Location = LocationEastBoat
			}
		}
	} else if state.anyInLocation(LocationEastBoat) {
		for _, actor := range state {
			if actor.Location == LocationEastBoat {
				actor.Location = LocationWestBoat
			}
		}
	}

	return state
}
