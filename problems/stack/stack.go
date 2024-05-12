package stack

import "errors"

type Stack struct {
	Arr []int
}

func NewStack() Stack {
	return Stack{
		Arr: []int{},
	}
}

func (s *Stack) Push(val int) {
	s.Arr = append(s.Arr, val)
}

func (s *Stack) Peep() (int, error) {
	if len(s.Arr) == 0 {
		return 0, errors.New("no values in stack")
	}

	val := s.Arr[len(s.Arr)-1]
	return val, nil
}

func (s *Stack) Pop() (int, error) {
	if len(s.Arr) == 0 {
		return 0, errors.New("no values in stack")
	}

	val := s.Arr[len(s.Arr)-1]
	s.Arr = s.Arr[:len(s.Arr)-1]

	return val, nil
}
