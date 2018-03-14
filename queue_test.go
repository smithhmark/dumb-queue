package dumbqueue

import (
	"testing"
	//"fmt"
)

func TestSPQConstruction(t *testing.T) {
	//fmt.Println("start")
	q := NewSlowGetQueue()
	if q.a.Size() != 0 || q.b.Size() != 0 {
		t.Fatalf("new queue has stuff in it")
	}
}

func TestSPQPutOneItem(t *testing.T) {
	//fmt.Println("start")
	q := NewSlowGetQueue()
	q.Put(0)
	if q.a.Size() != 1 && q.b.Size() != 0 {
		t.Fatalf("Put lost item")
	}
}

func TestSPQErrorOnEmpty(t *testing.T) {
	q := NewSlowGetQueue()
	_, err := q.Get()
	if err == nil {
		t.Fatalf("Get on empty queue should error")
	}
}

func TestSPQPutOneGetOneItem(t *testing.T) {
	//fmt.Println("start")
	q := NewSlowGetQueue()
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

func benchQPuts(q *SlowGetQueue, size int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		for ii := 0; ii < size; ii++ {
			q.Put(ii)
		}
	}
}

func benchQAlt(q *SlowGetQueue, size int, b *testing.B) {
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

func benchQAlt2(q *SlowGetQueue, size int, b *testing.B) {
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

func benchQSwing(q *SlowGetQueue, cnt int, amp int, b *testing.B) {
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

func BenchmarkSPQPut0(b *testing.B) { benchQPuts(NewSlowGetQueue(), 100, b) }
func BenchmarkSPQPut1(b *testing.B) { benchQPuts(NewSlowGetQueue(), 1000, b) }
func BenchmarkSPQPut2(b *testing.B) { benchQPuts(NewSlowGetQueue(), 10000, b) }
func BenchmarkSPQPut3(b *testing.B) { benchQPuts(NewSlowGetQueue(), 100000, b) }

func BenchmarkSPQAlt0(b *testing.B) { benchQAlt(NewSlowGetQueue(), 1000, b) }
func BenchmarkSPQAlt1(b *testing.B) { benchQAlt(NewSlowGetQueue(), 10000, b) }
func BenchmarkSPQAlt2(b *testing.B) { benchQAlt(NewSlowGetQueue(), 100000, b) }
func BenchmarkSPQAlt3(b *testing.B) { benchQAlt(NewSlowGetQueue(), 1000000, b) }
func BenchmarkSPQAlt4(b *testing.B) { benchQAlt(NewSlowGetQueue(), 10000000, b) }

func BenchmarkSPQAlt20(b *testing.B) { benchQAlt2(NewSlowGetQueue(), 1000, b) }
func BenchmarkSPQAlt21(b *testing.B) { benchQAlt2(NewSlowGetQueue(), 10000, b) }
func BenchmarkSPQAlt22(b *testing.B) { benchQAlt2(NewSlowGetQueue(), 100000, b) }
func BenchmarkSPQAlt23(b *testing.B) { benchQAlt2(NewSlowGetQueue(), 1000000, b) }
func BenchmarkSPQAlt24(b *testing.B) { benchQAlt2(NewSlowGetQueue(), 10000000, b) }

func BenchmarkSPQSw0(b *testing.B) {
	benchQSwing(NewSlowGetQueue(), 10000000, 1, b)
}
func BenchmarkSPQSw1(b *testing.B) {
	benchQSwing(NewSlowGetQueue(), 5000000, 2, b)
}
func BenchmarkSPQSw2(b *testing.B) {
	benchQSwing(NewSlowGetQueue(), 2000000, 5, b)
}
func BenchmarkSPQSw3(b *testing.B) {
	benchQSwing(NewSlowGetQueue(), 1000000, 10, b)
}
func BenchmarkSPQSw4(b *testing.B) {
	benchQSwing(NewSlowGetQueue(), 500000, 20, b)
}
