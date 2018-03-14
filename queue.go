package dumbqueue

//import "fmt"
import "errors"

type SlowGetQueue struct {
	a *Stack
	b *Stack
}

func NewSlowGetQueue() *SlowGetQueue {
	q := &SlowGetQueue{NewStack(), NewStack()}
	return q
}

func (q *SlowGetQueue) Put(ii int) {
	q.a.Push(ii)
}

func rev(src, dest *Stack) {
	for src.Size() > 0 {
		val, _ := src.Pop()
		dest.Push(val)
	}
}

func (q *SlowGetQueue) Get() (ret int, err error) {
	if q.a.Size() == 0 {
		return 0, errors.New("Queue is empty")
	}
	rev(q.a, q.b)

	var val interface{}
	val, err = q.b.Pop()
	if err != nil {
		return 0, errors.New("Something smells in Denmark...")
	}
	ret = val.(int)

	rev(q.b, q.a)
	err = nil
	return
}
