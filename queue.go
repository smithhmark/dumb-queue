package dumbqueue

//import "fmt"
import "errors"

type SlowGetQueue struct {
	a []int
	b []int
}

func (q *SlowGetQueue) Put(ii int) {
	q.a = append(q.a, ii)
}

func (q *SlowGetQueue) Get() (ret int, err error) {
	if len(q.a) == 0 {
		return 0, errors.New("Queue is empty")
	}
	err = nil
	for _, ii := range q.a {
		q.b = append(q.b, ii)
	}
	q.a = q.a[:0]
	// q.a = nil

	ret = q.b[len(q.b)-1]

	q.b = q.b[:len(q.b)-1]

	for _, ii := range q.b {
		q.a = append(q.a, ii)
	}
	q.b = q.b[:0]
	// q.b = nil
	return
}
