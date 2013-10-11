package events

import (

)
type EventHandler interface{
	HandleEvent(gEvent GVEvent)
	GetThreadContext() int64
	GetName() string
	TearDown()
	NeedsRead() bool
	NeedsConnect() bool
	NeedsWrite() bool
	NeedsAccept() bool
}