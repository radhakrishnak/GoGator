package main

import (
	"govisor/events"
	"fmt"
	"govisor/gvtimer"
	"govisor/slicer"
)
func String() {
	fmt.Println("hi")

}

func main() {
	String()
	events.Lite()
	test()
}
func test(){
	var s events.ConfigUpdateEvent
	var sc, dt events.EventHandler
	events.New(s)
	c:= new(events.ConfigUpdateEvent)
	c.GetConfig()
	var sad = gvtimer.NewGVTimer()
	fmt.Println(sad)
	v:=new(events.GVTimerEvent)
	fmt.Println(v.Getid())
	var a = events.NewGVTimerEvent(100, sc, dt)
	fmt.Println(a.Getid(),a.GetExpireTime())
	var b = events.NewGVTimerEvent(101, sc, dt)
	fmt.Println(b.Getid(),b.GetExpireTime())
	var f = events.NewGVTimerEvent(102, sc, dt)
	fmt.Println(f.Getid(),f.GetExpireTime())
	st := slicer.GVSlicer{}
	fmt.Println(st)	
}