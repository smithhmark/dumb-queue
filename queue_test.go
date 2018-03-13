package dumbqueue

import (
        "testing"
        //"fmt"
)

func TestSPQConstruction(t *testing.T) {
        //fmt.Println("start")
        q  := new(SlowGetQueue)
        if len(q.a) != 0 || len(q.b) != 0 {
                t.Fatalf("new queue has stuff in it")
        }
}

func TestSPQPutOneItem(t *testing.T) {
        //fmt.Println("start")
        q  := new(SlowGetQueue)
        q.Put(0)
        if len(q.a) != 1 && len(q.b) != 0{
                t.Fatalf("Put lost item")
        }
}

func TestSPQPutOneGetOneItem(t *testing.T) {
        //fmt.Println("start")
        q  := new(SlowGetQueue)
        item := 0
        q.Put(item)
        if len(q.a) != 1 && len(q.b) != 0{
                t.Fatalf("Put lost item")
        }
        ii := q.Get()
        if ii != item {
                t.Fatalf("Get did not return item")
        }
        if len(q.a) != 0 && len(q.b) != 0{
                t.Fatalf("Get didn't clear out inventory")
        }
}
