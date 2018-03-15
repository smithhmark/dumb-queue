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

func testErrorOnEmpty(qf QueueFactory, t *testing.T) {
	q:=qf()
	_, err := q.Get()
	if err == nil {
		t.Fatalf("Get on empty queue should error")
	}
}

func TestSGQErrorOnEmpty(t *testing.T) {
	testErrorOnEmpty(NewSlowGetQueue, t)
}
func TestSPQErrorOnEmpty(t *testing.T) {
	testErrorOnEmpty(NewSlowPutQueue, t)
}
func TestMQErrorOnEmpty(t *testing.T) {
	testErrorOnEmpty(NewModeQueue, t)
}

func testPutOneGetOneItem(qf QueueFactory, t *testing.T) {
	q := qf()
	testitem := 0
	if q.Size() != 0 {
		t.Fatalf("Queue should be empty")
	}
	q.Put(testitem)
	if q.Size() != 1 {
		t.Fatalf("Queue should have one item")
	}
	result, err := q.Get()
	if err != nil {
		t.Fatalf("Should be no error here")
	}
	if result != testitem {
		t.Fatalf("Didn't get out what we put in")
	}
	if q.Size() != 0 {
		t.Fatalf("Queue didn't shrink as part of Get")
	}
}

func TestSGQPutGet(t *testing.T) {
	testPutOneGetOneItem(NewSlowGetQueue, t)
}
func TestSPQPutGet(t *testing.T) {
	testPutOneGetOneItem(NewSlowPutQueue, t)
}
func TestMQPutGet(t *testing.T) {
	testPutOneGetOneItem(NewModeQueue, t)
}

func testMQPut(q *ModeQueue, item interface {}, pcnt, gcnt int, t *testing.T){
	q.Put(item)
	expectedSize := pcnt - gcnt

	if q.a.Size() != expectedSize {
		t.Fatalf("ModeQueue stack a should not be empty after(%v puts and %v gets)", pcnt, gcnt)
	}
	if q.b.Size() != 0 {
		t.Fatalf("ModeQueue stack b should start empty after(%v puts and %v gets)", pcnt, gcnt)
	}
	if q.putMode != true {
		t.Fatalf("ModeQueue should be in put mode after(%v puts and %v gets)", pcnt, gcnt)
	}
}

func testMQGet(q *ModeQueue, item interface {}, pcnt, gcnt int, t *testing.T){
	expectedSize := pcnt - gcnt
	rcvditem, err := q.Get()

	if expectedSize >= 0 {
		if err != nil { 
			t.Fatalf("there should be no error getting from non-empty ModeQueue(%v puts and %v gets)", pcnt, gcnt)
		}
		testMQState(q, 0, expectedSize, false, "after Get", t)
		if rcvditem != item {
			t.Fatalf("ModeQueue lost its item after(%v puts and %v gets)", pcnt, gcnt)
		}

	} else {
		if err == nil {
			t.Fatalf("there should be an error getting from empty ModeQueue(%v puts and %v gets)", pcnt, gcnt)
		}
	}
}

func testMQState(q *ModeQueue, asz, bsz int, putmode bool, msg string, t *testing.T) {
	if q.a.Size() != asz {
		t.Fatalf("ModeQueue stack A.Size() should be %v %v", asz, msg)
	}
	if q.b.Size() != bsz {
		t.Fatalf("ModeQueue stack B.Size() should be %v %v", bsz, msg)
	}
	if q.putMode != putmode {
		t.Fatalf("ModeQueue putmode should be %v %v", putmode, msg)
	}
}

func TestModeQueue(t *testing.T) {
	q := NewModeQueue().(*ModeQueue)

	testMQState(q, 0, 0, true, "at initialization", t)
	testMQPut(q, "item1", 1, 0, t)
	testMQPut(q, "item2", 2, 0, t)
	testMQGet(q, "item1", 2, 1, t)
	testMQGet(q, "item2", 2, 2, t)

	if q.Size() != 0 {
		t.Fatalf("ModeQueue should have 0 items")
	}
	testMQState(q, 0, 0, false, "after getting all items", t)

	testMQPut(q, "item3", 3,2,t)
}

func benchQPuts(q Queue, size int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		for ii := 0; ii < size; ii++ {
			q.Put(ii)
		}
	}
}

func benchQAlt(qf QueueFactory, size int, b *testing.B) {
	var test interface{}
	q := qf()
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

func benchQSwing(qf QueueFactory, cnt int, amp int, b *testing.B) {
	var test interface{}
	var data []int
	for v := 0; v < amp; v++ {
		data = append(data, v)
	}
	q := qf()
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

func BenchmarkSGQPut1(b *testing.B) { benchQPuts(NewSlowGetQueue(), 1000, b) }
func BenchmarkSGQPut2(b *testing.B) { benchQPuts(NewSlowGetQueue(), 10000, b) }
func BenchmarkSGQPut3(b *testing.B) { benchQPuts(NewSlowGetQueue(), 100000, b) }

func BenchmarkSGQAlt0(b *testing.B) { benchQAlt(NewSlowGetQueue, 1000, b) }
func BenchmarkSGQAlt1(b *testing.B) { benchQAlt(NewSlowGetQueue, 10000, b) }
func BenchmarkSGQAlt2(b *testing.B) { benchQAlt(NewSlowGetQueue, 100000, b) }
func BenchmarkSGQAlt3(b *testing.B) { benchQAlt(NewSlowGetQueue, 1000000, b) }
func BenchmarkSGQAlt4(b *testing.B) { benchQAlt(NewSlowGetQueue, 10000000, b) }
func BenchmarkSPQAlt0(b *testing.B) { benchQAlt(NewSlowPutQueue, 1000, b) }
func BenchmarkSPQAlt1(b *testing.B) { benchQAlt(NewSlowPutQueue, 10000, b) }
func BenchmarkSPQAlt2(b *testing.B) { benchQAlt(NewSlowPutQueue, 100000, b) }
func BenchmarkSPQAlt3(b *testing.B) { benchQAlt(NewSlowPutQueue, 1000000, b) }
func BenchmarkSPQAlt4(b *testing.B) { benchQAlt(NewSlowPutQueue, 10000000, b) }

func BenchmarkSGQ2Alt0(b *testing.B) { benchQAlt2(NewSlowGetQueue(), 1000, b) }
func BenchmarkSGQ2Alt1(b *testing.B) { benchQAlt2(NewSlowGetQueue(), 10000, b) }
func BenchmarkSGQ2Alt2(b *testing.B) { benchQAlt2(NewSlowGetQueue(), 100000, b) }
func BenchmarkSGQ2Alt3(b *testing.B) { benchQAlt2(NewSlowGetQueue(), 1000000, b) }
func BenchmarkSGQ2Alt4(b *testing.B) { benchQAlt2(NewSlowGetQueue(), 10000000, b) }

func BenchmarkSGQSw0(b *testing.B) {
	benchQSwing(NewSlowGetQueue, 10000000, 1, b)
}
func BenchmarkSGQSw1(b *testing.B) {
	benchQSwing(NewSlowGetQueue, 5000000, 2, b)
}
func BenchmarkSGQSw2(b *testing.B) {
	benchQSwing(NewSlowGetQueue, 2000000, 5, b)
}
func BenchmarkSGQSw3(b *testing.B) {
	benchQSwing(NewSlowGetQueue, 1000000, 10, b)
}
func BenchmarkSGQSw4(b *testing.B) {
	benchQSwing(NewSlowGetQueue, 500000, 20, b)
}

func BenchmarkSPQSw0(b *testing.B) {
	benchQSwing(NewSlowPutQueue , 10000000, 1, b)
}
func BenchmarkSPQSw1(b *testing.B) {
	benchQSwing(NewSlowPutQueue, 5000000, 2, b)
}
func BenchmarkSPQSw2(b *testing.B) {
	benchQSwing(NewSlowPutQueue, 2000000, 5, b)
}
func BenchmarkSPQSw3(b *testing.B) {
	benchQSwing(NewSlowPutQueue, 1000000, 10, b)
}
func BenchmarkSPQSw4(b *testing.B) {
	benchQSwing(NewSlowPutQueue, 500000, 20, b)
}

func BenchmarkMQSw0(b *testing.B) {
	benchQSwing(NewModeQueue , 10000000, 1, b)
}
func BenchmarkMQSw1(b *testing.B) {
	benchQSwing(NewModeQueue, 5000000, 2, b)
}
func BenchmarkMQSw2(b *testing.B) {
	benchQSwing(NewModeQueue, 2000000, 5, b)
}
func BenchmarkMQSw3(b *testing.B) {
	benchQSwing(NewModeQueue, 1000000, 10, b)
}
func BenchmarkMQSw4(b *testing.B) {
	benchQSwing(NewModeQueue, 500000, 20, b)
}

