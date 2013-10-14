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


