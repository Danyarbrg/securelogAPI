package models

import (
	"time"
)

type LogEntry struct {
	ID			int
	Timestamp	time.Time
	EventType	string
	UserID		string
	Payload		string
}
