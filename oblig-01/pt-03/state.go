package pt03

import (
	"fmt"
	"strings"
)

type ActorType int

const (
	ActorPerson ActorType = iota
	ActorFox    ActorType = iota
	ActorHen    ActorType = iota
	ActorGrain  ActorType = iota
)

type ActorLocation int

const (
	LocationWestLand ActorLocation = iota
	LocationWestBoat ActorLocation = iota
	LocationEastBoat ActorLocation = iota
	LocationEastLand ActorLocation = iota
)

type Actor struct {
	Type     ActorType
	Location ActorLocation
}

func (a Actor) visual() string {
	switch a.Type {
	case ActorPerson:
		return "🧑"
	case ActorFox:
		return "🦊"
	case ActorHen:
		return "🐔"
	case ActorGrain:
		return "🌾"
	default:
		panic(fmt.Sprintf("Unsupported actory type: %v", a.Type))
	}
}

func (a Actor) view(location ActorLocation) string {
	if a.Location == location {
		return a.visual()
	} else {
		return " "
	}
}

type State []*Actor

func (s State) get(actorType ActorType) *Actor {
	for _, actor := range s {
		if actor.Type == actorType {
			return actor
		}
	}

	panic(fmt.Sprintf("Unsupported actor type: %v", actorType))
}

func (s State) anyInLocation(location ActorLocation) bool {
	for _, actor := range s {
		if actor.Location == location {
			return true
		}
	}

	return false
}

func (s State) viewBoat(location ActorLocation) string {
	passenger := strings.TrimSpace(fmt.Sprintf(
		"%s%s%s",
		s.get(ActorFox).view(location),
		s.get(ActorHen).view(location),
		s.get(ActorGrain).view(location),
	))
	if passenger == "" {
		passenger = "_"
	}

	person := strings.TrimSpace(s.get(ActorPerson).view(location))
	if person == "" {
		person = "_"
	}

	return fmt.Sprintf(`\_%s_%s_/`, passenger, person)
}

func NewState() State {
	return State{
		&Actor{
			Type:     ActorPerson,
			Location: LocationWestLand,
		},
		&Actor{
			Type:     ActorFox,
			Location: LocationWestLand,
		},
		&Actor{
			Type:     ActorHen,
			Location: LocationWestLand,
		},
		&Actor{
			Type:     ActorGrain,
			Location: LocationWestLand,
		},
	}
}

func ViewState(state State) string {
	westBank := "_______"
	eastBank := "_______"

	if !state.anyInLocation(LocationEastBoat) {
		westBank = state.viewBoat(LocationWestBoat)
	} else {
		eastBank = state.viewBoat(LocationEastBoat)
	}

	return fmt.Sprintf(
		`[ %s %s %s %s ---\_%s_%s_/--- %s %s %s %s ]`,
		state.get(ActorFox).view(LocationWestLand),
		state.get(ActorHen).view(LocationWestLand),
		state.get(ActorGrain).view(LocationWestLand),
		state.get(ActorPerson).view(LocationWestLand),
		westBank,
		eastBank,
		state.get(ActorPerson).view(LocationEastLand),
		state.get(ActorGrain).view(LocationEastLand),
		state.get(ActorHen).view(LocationEastLand),
		state.get(ActorFox).view(LocationEastLand),
	)
}
