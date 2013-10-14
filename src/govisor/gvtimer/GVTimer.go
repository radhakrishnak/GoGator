package gvtimer

import (

)
//MAX_TIMEOUT long = 5000

type GVTimer struct{
	max_timeout int64
	min_timeout int64
}
func NewGVTimer() *GVTimer {
	return &GVTimer{max_timeout:5000, min_timeout:1}
}	

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	gvtimerevent := x.(*GVTimerEvent)
	gvtimerevent.index = n
	*pq = append(*pq, gvtimerevent)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	gvtimerevent := old[n-1]
	gvtimerevent.index = -1 // for safety
	*pq = old[0 : n-1]
	return gvtimerevent
}