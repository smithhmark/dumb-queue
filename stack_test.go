package dumbqueue

import "testing"

func TestNewStack(t *testing.T) {
	s := NewStack()
	if s.head != nil {
		t.Error("head should be nil")
	}

	if s.length != 0 {
		t.Error("length of empty stack should be 0")
	}
}

func TestNewNode(t *testing.T) {
	item := "Hi"
	n := newNode(item, nil)
	if n.value != item {
		t.Error("where did the item to store go?")
	}
	if n.next != nil {
		t.Error("a singleton node shouldn't have a next")
	}
}

func TestPush(t *testing.T) {
	s := NewStack()
	item := "Howdy"
	s.Push(item)

	if s.length != 1 {
		t.Error("Size wrong after put")
	}
	if s.head == nil {
		t.Error("head should point to a node and not be nil")
	}
}

func TestPushTwice(t *testing.T) {
	s := NewStack()
	items := []string{"Howdy", "duty"}
	s.Push(items[0])
	s.Push(items[1])

	var np *node

	if s.head == nil {
		t.Error("second Push broke head")
	}
	np = s.head
	if np.next == nil {
		t.Error("head node not pointing to next")
	}
	if np.value != items[1] {
		t.Error("lost second item")
	}
	np = np.next
	if np.next != nil {
		t.Error("second and final node not pointing to nil")
	}
	if np.value != items[0] {
		t.Error("lost first item")
	}
}

func TestPopEmpty(t *testing.T) {
	s := NewStack()
	val, err := s.Pop()

	if val != nil {
		t.Error("should return nil from empty Stack")
	}
	if err == nil {
		t.Error("there should be a non-nil error")
	}
}

func TestPushOnePopTwo(t *testing.T) {
	s := NewStack()
	item := "salutations"
	s.Push(item)
	if s.Size() != 1 {
		t.Error("What happened to my item?")
	}
	val, err := s.Pop()
	if val != item {
		t.Error("failed to retrieve stored item")
	}
	if err != nil {
		t.Error("Error should be nil when Popping from non-empty Stack")
	}

	val, err = s.Pop()
	if val != nil {
		t.Error("should return nil from empty Stack")
	}
	if err == nil {
		t.Error("there should be a non-nil error")
	}
}
