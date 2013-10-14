package slicer

import (
	"govisor/events"
	"time"
)

type ReconnectEvent struct{
	timerEvent events.GVTimerEvent	
}

func NewReconnectEvent(secondsToNextReconnect int, src events.EventHandler) *ReconnectEvent{
	seconds := time.Now().Unix()
	expiryTime := 1000 * ( int64 (secondsToNextReconnect) + seconds)
	p:=events.NewGVTimerEvent(expiryTime, src, src)
	return &ReconnectEvent{timerEvent:*p}
}
