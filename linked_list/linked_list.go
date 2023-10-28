package linkedlist

import "errors"

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

type Linkedlist[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

func New[T any]() *Linkedlist[T] {
	n := &Node[T]{}
	l := &Linkedlist[T]{
		Head: n,
		Tail: n,
	}

	return l
}

func (l *Linkedlist[T]) checkRemoveAbility() error {
	if l.Size == 0 {
		return errors.New("cannot remove element from empty slice")
	}

	return nil
}

func (l *Linkedlist[T]) nodeAt(pos int) *Node[T] {
	if l.Size == 0 {
		return nil
	}

	node := l.Head
	pos = (l.Size+pos)%l.Size + 1 // This trick allows us to use negative pos

	if pos < 0 {
		return nil
	}

	for pos != 0 && node != nil {
		node = node.Next
		pos--
	}

	return node
}

func (l *Linkedlist[T]) Append(val T) {
	node := &Node[T]{
		Val:  val,
		Next: nil,
	}

	l.Tail.Next = node
	l.Tail = node
	l.Size++
}

func (l *Linkedlist[T]) Prepend(val T) {
	node := &Node[T]{
		Val:  val,
		Next: l.Head.Next,
	}

	l.Head.Next = node
	l.Size++
}

func (l *Linkedlist[T]) Insert(pos int, val T) {
	prevNode := l.nodeAt(pos - 1)
	node := &Node[T]{
		Val:  val,
		Next: prevNode.Next,
	}

	prevNode.Next = node
}

func (l *Linkedlist[T]) Pop() (T, error) {
	if err := l.checkRemoveAbility(); err != nil {
		return *new(T), err
	}

	prevNode := l.nodeAt(l.Size - 1)
	node := l.Tail

	l.Tail = prevNode
	prevNode.Next = nil
	l.Size--

	return node.Val, nil
}

func (l *Linkedlist[T]) PopLeft() (T, error) {
	if err := l.checkRemoveAbility(); err != nil {
		return *new(T), err
	}

	node := l.Head.Next
	l.Head.Next = node.Next
	l.Size--

	return node.Val, nil
}

func (l *Linkedlist[T]) At(pos int) T {
	return l.nodeAt(pos).Val
}

func (l *Linkedlist[T]) RemoveAt(pos int) error {
	if err := l.checkRemoveAbility(); err != nil {
		return err
	}

	prevNode := l.nodeAt(pos - 1)
	prevNode.Next = prevNode.Next.Next
	l.Size--

	return nil
}

func (l *Linkedlist[T]) UpdateAt(pos int, val T) {
	l.nodeAt(pos).Val = val
}

func (l *Linkedlist[T]) ToSlice() []T {
	slice := make([]T, 0)
	node := l.Head.Next

	for node != nil {
		slice = append(slice, node.Val)
		node = node.Next
	}

	return slice
}

func FromSlice[T any](elems []T) *Linkedlist[T] {
	l := New[T]()

	for _, e := range elems {
		l.Append(e)
	}

	return l
}
