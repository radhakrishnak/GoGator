package utils

// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
)



// A PriorityQueue implements heap.Interface and holds GVTimerEvents.
type PriorityQueue []*GVTimeEvent

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
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

