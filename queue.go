package dumbqueue

import "fmt"
import "errors"

type SlowGetQueue struct {
        a []int
        b []int
}

func (q *SlowGetQueue) Put(ii int) {
        q.a = append(q.a, ii);
}

func (q *SlowGetQueue) Get() (ret int, err error) {
        if len(q.a) == 0 {
                return 0, errors.New("Queue is empty")
        }
        err = nil
        fmt.Println("begin Get")
        for _, ii := range q.a {
                q.b = append(q.b, ii)
        }
        fmt.Println("clear a")
        q.a = q.a[:0]
        // q.a = nil

        fmt.Println("get return val")
        ret = q.b[len(q.b)-1];

        fmt.Println("shrinking b")
        q.b = q.b[:len(q.b)-1];

        fmt.Println("shuffling into a")
        for _, ii := range q.b {
                q.a = append(q.a, ii);
        }
        fmt.Println("clearing b")
        q.b = q.b[:0]
        // q.b = nil
        return
}


