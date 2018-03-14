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
