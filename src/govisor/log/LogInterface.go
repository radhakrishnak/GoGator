package log

import (
	"govisor/events"
	"log"
)

type LogInterface interface{
	init() bool
	log(level log.Logger, time int64, source events.EventHandler, msg string)
}
