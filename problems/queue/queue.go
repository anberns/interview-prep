package queue

import "errors"

type Queue struct {
	Arr []int
}

func NewQueue() Queue {
	return Queue{
		Arr: []int{},
	}
}

func (q *Queue) Peek() (int, error) {
	if len(q.Arr) == 0 {
		return 0, errors.New("queue empty")
	}

	return q.Arr[0], nil
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.Arr) == 0 {
		return 0, errors.New("queue empty")
	}

	val := q.Arr[0]
	q.Arr = q.Arr[1:]

	return val, nil
}

func (q *Queue) Enqueue(val int) {
	q.Arr = append(q.Arr, val)
}
