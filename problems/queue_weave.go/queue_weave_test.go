package queueweave_test

import (
	"testing"

	"example.com/queue"
	queueweave "example.com/queue_weave.go"
	"github.com/stretchr/testify/require"
)

func TestQueueWeave(t *testing.T) {

	q1 := queue.NewQueue()
	q2 := queue.NewQueue()

	q1.Enqueue(1)
	q1.Enqueue(2)
	q1.Enqueue(3)
	q2.Enqueue(4)
	q2.Enqueue(5)
	q2.Enqueue(6)

	result := queueweave.Weave(q1, q2)
	resultArr := make([]int, 0, 6)
	for {
		v, err := result.Dequeue()
		if err != nil {
			break
		}
		resultArr = append(resultArr, v)
	}
	require.Equal(t, []int{1, 4, 2, 5, 3, 6}, resultArr)
}
