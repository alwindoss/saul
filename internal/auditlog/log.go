package auditlog

type EventType int

const (
	AccessControlEvent EventType = iota
	SystemEvent
)

type Log struct {
	EventType EventType
}
