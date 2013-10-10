package events

import (
	"fmt"
)
type EventHandler interface{
	HandleEvent()
	GetThreadContext() int64
	GetName() string
	TearDown()
	NeedsRead() bool
	NeedsConnect() bool
	NeedsWrite() bool
	NeedsAccept() bool
}
type GVEvent struct {
	src, dst EventHandler
}

func (gEvent GVEvent) GetSrc() EventHandler{
	return gEvent.src
}
func (gEvent GVEvent) GetDst() EventHandler{
	return gEvent.dst
}
func Lite(){
	var a, b EventHandler
	k := GVEvent{src:a, dst:b}
	fmt.Println(k)
}