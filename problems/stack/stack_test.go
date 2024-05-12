package stack_test

import (
	"testing"

	"example.com/stack"
	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {

	s := stack.NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	result, err := s.Peek()
	require.NoError(t, err)
	require.Equal(t, 3, result)

	result, err = s.Pop()
	require.NoError(t, err)
	require.Equal(t, 3, result)

	result, err = s.Peek()
	require.NoError(t, err)
	require.Equal(t, 2, result)

	_, err = s.Pop()
	require.NoError(t, err)
	_, err = s.Pop()
	require.NoError(t, err)

	_, err = s.Peek()
	require.Error(t, err)

	_, err = s.Pop()
	require.Error(t, err)
}
