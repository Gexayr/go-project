package event

const (
	EventLinkVisited
)

type Event struct {
	Type string
	Data any
}

type EventBus struct {
	bus chan Event
}

func NewEventBus() *EventBus {
	return &EventBus{
		bus: make(chan Event),
	}
}

func (e *EventBus) Publish(evt Event) {
	e.bus <- evt
}

func (e *EventBus) Subscribe() <-chan Event {
	return e.bus
}
