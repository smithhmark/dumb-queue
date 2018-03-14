package dumbqueue

import "errors"

type Stack struct {
	head   *node
	length int
}

type node struct {
	value interface{}
	next  *node
}

func NewStack() *Stack {
	s := &Stack{nil, 0}
	return s
}

func newNode(item interface{}, next *node) *node {
	n := &node{item, next}
	return n
}

func (s *Stack) Size() int {
	return s.length
}

func (s *Stack) Push(item interface{}) {
	newHead := newNode(item, s.head)
	s.head, s.length = newHead, s.length+1
}

func (s *Stack) Pop() (interface{}, error) {
	if s.length == 0 {
		return nil, errors.New("Pop on empty Stack")
	}
	np := s.head
	s.head = np.next
	s.length--
	np.next = nil
	return np.value, nil
}
