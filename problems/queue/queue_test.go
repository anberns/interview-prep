package queue_test

import (
	"testing"

	"example.com/queue"
	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {

	q := queue.NewQueue()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	result, err := q.Peek()
	require.NoError(t, err)
	require.Equal(t, 1, result)

	result, err = q.Dequeue()
	require.NoError(t, err)
	require.Equal(t, 1, result)

	result, err = q.Peek()
	require.NoError(t, err)
	require.Equal(t, 2, result)

	q.Enqueue(5)

	result, err = q.Peek()
	require.NoError(t, err)
	require.Equal(t, 2, result)

	_, err = q.Dequeue()
	require.NoError(t, err)

	result, err = q.Peek()
	require.NoError(t, err)
	require.Equal(t, 3, result)

	_, err = q.Dequeue()
	require.NoError(t, err)

	result, err = q.Peek()
	require.NoError(t, err)
	require.Equal(t, 5, result)

	_, err = q.Dequeue()
	require.NoError(t, err)

	_, err = q.Peek()
	require.Error(t, err)

	_, err = q.Dequeue()
	require.Error(t, err)
}
