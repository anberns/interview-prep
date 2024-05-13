package queue

import "errors"

type Queue struct {
	arr []int
}

func NewQueue() Queue {
	return Queue{
		arr: []int{},
	}
}

func (q *Queue) Peek() (int, error) {
	if len(q.arr) == 0 {
		return 0, errors.New("queue empty")
	}

	return q.arr[0], nil
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.arr) == 0 {
		return 0, errors.New("queue empty")
	}

	val := q.arr[0]
	q.arr = q.arr[1:]

	return val, nil
}

func (q *Queue) Enqueue(val int) {
	q.arr = append(q.arr, val)
}
