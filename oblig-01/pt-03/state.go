package pt03

type ActorType int

const (
	ActorPerson ActorType = iota
	ActorFox    ActorType = iota
	ActorHen    ActorType = iota
	ActorGrain  ActorType = iota
)

type State bool

func NewState() State {
	return true
}

func ViewState(state State) string {
	return ""
}
