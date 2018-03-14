package dumbqueue

import (
	"testing"
	//"fmt"
)

func TestSGQConstruction(t *testing.T) {
	q := NewSlowGetQueue().(*SlowGetQueue)
	if q.a.Size() != 0 || q.b.Size() != 0 {
		t.Fatalf("new queue has stuff in it")
	}
}

func TestSGQPutOneItem(t *testing.T) {
	q := NewSlowGetQueue().(*SlowGetQueue)
	q.Put(0)
	if q.a.Size() != 1 && q.b.Size() != 0 {
		t.Fatalf("Put lost item")
	}
}

func TestSGQErrorOnEmpty(t *testing.T) {
	q := NewSlowGetQueue().(*SlowGetQueue)
	_, err := q.Get()
	if err == nil {
		t.Fatalf("Get on empty queue should error")
	}
}

func TestSGQPutOneGetOneItem(t *testing.T) {
	//fmt.Println("start")
	q := NewSlowGetQueue().(*SlowGetQueue)
	item := 0
	q.Put(item)
	if q.a.Size() != 1 && q.b.Size() != 0 {
		t.Fatalf("Put lost item")
	}
	ii, err := q.Get()
	if err != nil {
		t.Fatalf("Get should work if theres data in the queue")
	}
	if ii != item {
		t.Fatalf("Get did not return item")
	}
	if q.a.Size() != 0 && q.b.Size() != 0 {
		t.Fatalf("Get didn't clear out inventory")
	}
}

func benchQPuts(q Queue, size int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		for ii := 0; ii < size; ii++ {
			q.Put(ii)
		}
	}
}

func benchQAlt(q Queue, size int, b *testing.B) {
	var test interface{}
	for n := 0; n < b.N; n++ {
		for ii := 0; ii < size; ii++ {
			q.Put(ii)
			test, _ = q.Get()
			if test.(int) != ii {
				b.Error("args")
			}
		}
	}
}

func benchQAlt2(q Queue, size int, b *testing.B) {
	var test interface{}
	testsz := size >> 1
	for n := 0; n < b.N; n++ {
		for ii := 0; ii < testsz; ii++ {
			q.Put(ii)
			q.Put(ii * 2)
			test, _ = q.Get()
			if test.(int) != ii {
				b.Error("args")
			}
			test, _ = q.Get()
			if test.(int) != ii*2 {
				b.Error("args")
			}
		}
	}
}

func benchQSwing(q Queue, cnt int, amp int, b *testing.B) {
	var test interface{}
	var data []int
	for v := 0; v < amp; v++ {
		data = append(data, v)
	}
	for n := 0; n < b.N; n++ {
		for cycle := 0; cycle < cnt; cycle++ {
			for i := 0; i < amp; i++ {
				q.Put(data[i])
			}
			for i := 0; i < amp; i++ {
				test, _ = q.Get()
				if test.(int) != data[i] {
					b.Error("lost a value")
				}
			}
		}
	}
}

func BenchmarkSGQPut0(b *testing.B) {
	var q Queue
	q = NewSlowGetQueue()
	benchQPuts(q, 100, b)
}

func BenchmarkSGQPut1(b *testing.B) { benchQPuts(NewSlowGetQueue(), 1000, b) }
func BenchmarkSGQPut2(b *testing.B) { benchQPuts(NewSlowGetQueue(), 10000, b) }
func BenchmarkSGQPut3(b *testing.B) { benchQPuts(NewSlowGetQueue(), 100000, b) }

func BenchmarkSGQAlt0(b *testing.B) { benchQAlt(NewSlowGetQueue(), 1000, b) }
func BenchmarkSGQAlt1(b *testing.B) { benchQAlt(NewSlowGetQueue(), 10000, b) }
func BenchmarkSGQAlt2(b *testing.B) { benchQAlt(NewSlowGetQueue(), 100000, b) }
func BenchmarkSGQAlt3(b *testing.B) { benchQAlt(NewSlowGetQueue(), 1000000, b) }
func BenchmarkSGQAlt4(b *testing.B) { benchQAlt(NewSlowGetQueue(), 10000000, b) }

func BenchmarkSGQAlt20(b *testing.B) { benchQAlt2(NewSlowGetQueue(), 1000, b) }
func BenchmarkSGQAlt21(b *testing.B) { benchQAlt2(NewSlowGetQueue(), 10000, b) }
func BenchmarkSGQAlt22(b *testing.B) { benchQAlt2(NewSlowGetQueue(), 100000, b) }
func BenchmarkSGQAlt23(b *testing.B) { benchQAlt2(NewSlowGetQueue(), 1000000, b) }
func BenchmarkSGQAlt24(b *testing.B) { benchQAlt2(NewSlowGetQueue(), 10000000, b) }

func BenchmarkSGQSw0(b *testing.B) {
	benchQSwing(NewSlowGetQueue(), 10000000, 1, b)
}
func BenchmarkSGQSw1(b *testing.B) {
	benchQSwing(NewSlowGetQueue(), 5000000, 2, b)
}
func BenchmarkSGQSw2(b *testing.B) {
	benchQSwing(NewSlowGetQueue(), 2000000, 5, b)
}
func BenchmarkSGQSw3(b *testing.B) {
	benchQSwing(NewSlowGetQueue(), 1000000, 10, b)
}
func BenchmarkSGQSw4(b *testing.B) {
	benchQSwing(NewSlowGetQueue(), 500000, 20, b)
}

