package framework

import "time"

type RuntimeEvent string

const (
	RuntimeEventPanicRecovered    RuntimeEvent = "panic_recovered"
	RuntimeEventShutdownStarted   RuntimeEvent = "shutdown_started"
	RuntimeEventShutdownCompleted RuntimeEvent = "shutdown_completed"
)

type RuntimeEventPayload struct {
	Event     RuntimeEvent
	Timestamp time.Time
	Data      any
}

type RuntimeEventHandler func(RuntimeEventPayload)
