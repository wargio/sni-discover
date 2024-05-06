package main

import (
	"sync"
)

type Queue struct {
	elements []interface{}
	mutex    sync.Mutex
}

func (q *Queue) Size() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return len(q.elements)
}

func (q *Queue) Push(x interface{}) {
	q.mutex.Lock()
	q.elements = append(q.elements, x)
	q.mutex.Unlock()
}

func (q *Queue) Pop() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if len(q.elements) < 1 {
		return nil
	}

	var first interface{}
	a := q.elements
	l := len(a)
	first, q.elements = a[0], a[1:l]
	return first
}
