package events

import (
	"fmt"
)

type GVEvent struct {
	src, dst EventHandler
}

func (gEvent GVEvent) GetSrc() EventHandler{
	return gEvent.src
}
func (gEvent GVEvent) GetDst() EventHandler{
	return gEvent.dst
}
func (gEvent GVEvent) SetSrc (srcHandler EventHandler) GVEvent {
	gEvent.src = srcHandler
	return gEvent
}
func (gEvent GVEvent)SetDst(dstHandler EventHandler) GVEvent {
	gEvent.dst = dstHandler
	return gEvent
}
func Lite(){
	var a, b EventHandler
	k := GVEvent{src:a, dst:b}
	fmt.Println(k)
}