package shared

import (
	"time"
)

type Event interface {
	GetID() int
	GetName() string
	GetTimestamp() time.Time
}

type GameEvent interface {
	Event
	GetGameName() string
}

type TableEvent interface {
	Event
	GetTableID() string
}

type HandEvent interface {
	Event
	GetHandID() string
}
