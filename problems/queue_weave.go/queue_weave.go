package queueweave

import "example.com/queue"

func Weave(q1, q2 queue.Queue) queue.Queue {
	wq := queue.NewQueue()

	var shortest int
	for {
		val1, err := q1.Dequeue()
		if err != nil {
			shortest = 1
			break
		}

		wq.Enqueue(val1)

		val2, err := q2.Dequeue()
		if err != nil {
			shortest = 2
			break
		}

		wq.Enqueue(val2)
	}

	if shortest == 1 {
		for {
			val2, err := q2.Dequeue()
			if err != nil {
				break
			}

			wq.Enqueue(val2)
		}
	} else {
		for {
			val1, err := q1.Dequeue()
			if err != nil {
				break
			}

			wq.Enqueue(val1)
		}
	}

	return wq

}
