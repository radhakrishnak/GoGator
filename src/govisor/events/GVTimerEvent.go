package events

import (

)
type ID int 
type GVTimerEvent struct{
	expireTime int64
	id int
}

func NewGVTimerEvent() *GVTimerEvent {
	return &GVTimerEvent{expireTime:50, id:1}
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