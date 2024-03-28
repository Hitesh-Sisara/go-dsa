package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []interface{}
}

func NewStack() *Stack {
	return &Stack{items: []interface{}{}}
}

func (s *Stack) push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) pop() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, errors.New("Stack is empty ")
	}

	item := s.items[len(s.items)-1]

	s.items = s.items[:len(s.items)-1]

	return item, nil
}

type Queue struct {
	items []interface{}
}

func NewQueue() *Queue {
	return &Queue{items: []interface{}{}}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (interface{}, error) {
	if len(q.items) == 0 {
		return nil, errors.New("Queue is empty")
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func main() {
	st := NewStack()
	qu := NewQueue()

	st.push("apple")
	st.push("banana")

	st.pop()

	fmt.Println(st.items...)

	qu.Enqueue("car")
	qu.Enqueue("bus")
	qu.Dequeue()

	fmt.Println(qu.items...)
}
