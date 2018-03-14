package dumbqueue

//import "fmt"
import "errors"

type Queue interface {
	Put(interface{})
	Get() (interface {}, error)
	Size() int
}

type QueueFactory func () Queue

type StackBasedQueue interface {
	Stack1() (*Stack)
	Stack2() (*Stack)
}

func rev(src, dest *Stack) {
	for src.Size() > 0 {
		val, _ := src.Pop()
		dest.Push(val)
	}
}

type SlowGetQueue struct {
	a *Stack
	b *Stack
}

func NewSlowGetQueue() Queue {
	q := &SlowGetQueue{NewStack(), NewStack()}
	return q
}

func (q SlowGetQueue) Stack1()  *Stack { return q.a }
func (q SlowGetQueue) Stack2()  *Stack { return q.b }

func (q SlowGetQueue) Put(item interface{}) {
	q.a.Push(item)
}
func (q SlowGetQueue) Size() int {
	return q.a.Size()
}

func (q SlowGetQueue) Get() (ret interface{}, err error) {
	if q.a.Size() == 0 {
		return 0, errors.New("Queue is empty")
	}
	rev(q.a, q.b)

	ret, err = q.b.Pop()
	if err != nil {
		return 0, errors.New("Something smells in Denmark...")
	}

	rev(q.b, q.a)
	err = nil
	return
}

type SlowPutQueue struct {
	a *Stack
	b *Stack
}

func (q SlowPutQueue) Stack1()  *Stack { return q.a }
func (q SlowPutQueue) Stack2()  *Stack { return q.b }

func NewSlowPutQueue() Queue {
	q := &SlowPutQueue{NewStack(), NewStack()}
	return q
}

func (q SlowPutQueue) Size() int {
	return q.a.Size()
}

func (q SlowPutQueue) Put(item interface{}) {
	rev(q.a, q.b)
	q.b.Push(item)
	rev(q.b, q.a)
}

func (q SlowPutQueue) Get() (ret interface{}, err error) {
	if q.a.Size() == 0 {
		return 0, errors.New("Queue is empty")
	}

	ret, err = q.a.Pop()
	if err != nil {
		return 0, errors.New("Something smells in Denmark...")
	}
	err = nil
	return
}
