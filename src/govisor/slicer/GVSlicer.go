package slicer

import (

)
const MessagesPerRead int = 50

type GVSlicer struct{
	sliceName string
	// some variables yet to be constructed
	hostname string
	reconnectSeconds int
	maxReconnectSeconds int //should be initialised
	port int
	isConnected bool
	connectCount int //init to 0
	missSendLength int8
	allowAllPorts bool
	isShutdown bool
	lldpOptIn bool // default to true
	floodPerms bool
	allowedPorts map[int8]bool
	reconnectEventScheduled bool //default to false
	
}

func (gSlicer GVSlicer)updatePortList(){
	//need to implement after starting Config structs.
}







