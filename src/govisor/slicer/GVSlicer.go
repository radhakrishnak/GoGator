package slicer

import (

)
const MessagesPerRead int = 50
const max_allowed_buffer_ids = 256

type GVSlicer struct{
	sliceName string
	// some variables yet to be constructed
	hostname string
	reconnectSeconds int
	maxReconnectSeconds int //should be initialised
	port int // tcp port of controller
	isConnected bool
	connectCount int //init to 0
	missSendLength int
	allowAllPorts bool
	isShutdown bool
	lldpOptIn bool // default to true
	floodPerms bool
	allowedPorts map[int8]bool
	reconnectEventScheduled bool //default to false
	fmlimit int
}

func NewGVSlicer(sName string) *GVSlicer{
	sl := GVSlicer{sliceName:sName}
	sl.missSendLength = 128
	sl.fmlimit = -1
	sl.maxReconnectSeconds = 15
	sl.lldpOptIn = true
	return &sl
}

func (gSlicer *GVSlicer)SetMissSendLength(val int){
	gSlicer.missSendLength = val
} 

func (gSlicer GVSlicer)GetMissSendLength() int{
	return gSlicer.missSendLength
}

func (gSlicer GVSlicer)updatePortList(){
	//need to implement after starting Config structs.
}







