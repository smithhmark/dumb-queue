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

func (q *SlowGetQueue) Get() (ret int, err error) {
	if q.a.Size() == 0 {
		return 0, errors.New("Queue is empty")
	}
	for q.a.Size() > 0 {
		val, _ := q.a.Pop()
		q.b.Push(val)
	}

	var val interface{}
	val, err = q.b.Pop()
	if err != nil {
		return 0, errors.New("Something smells in Denmark...")
	}
	ret = val.(int)

	for q.b.Size() > 0 {
		val, _ := q.b.Pop()
		q.a.Push(val)
	}
	err = nil
	return
}
