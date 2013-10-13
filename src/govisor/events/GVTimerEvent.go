package events

import (

)
var ID int = 0
type GVTimerEvent struct{
	gEvent GVEvent
	expireTime int64
	id int
}

func NewGVTimerEvent(expTime int64, source EventHandler, dest EventHandler) *GVTimerEvent {
	ID = ID + 1
	return &GVTimerEvent{gEvent:*NewGVEvent(source,dest),expireTime:expTime, id:ID}
}

func (timer *GVTimerEvent) SetExpireTime(expireTime int64){
	timer.expireTime = expireTime
}

func (timer GVTimerEvent) GetExpireTime() int64{
	return timer.expireTime
}
func (timer *GVTimerEvent) Setid(id int){
	timer.id = id
}
func (timer GVTimerEvent) Getid() int{
	return timer.id
}

func (timer GVTimerEvent)CompareTo(timerEvent GVTimerEvent) int64{
	if timerEvent.expireTime != timer.expireTime{
		return (timer.expireTime - timerEvent.expireTime)	
	}else{
		return (int64)(timer.id - timerEvent.id)
	}
}