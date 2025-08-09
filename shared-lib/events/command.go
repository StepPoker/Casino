package shared

type Command interface {
	GetName() string
	GetVersion() int
}

type GameCommand interface {
	Command
	GetGameID() string
}

type TableCommand interface {
	Command
	GetTableID() string
}

type HandCommand interface {
	Command
	GetHandID() string
}
