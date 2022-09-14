package LinkedList

import (
	"errors"
	"fmt"
)

type Node[T comparable] struct {
	data T
	next *Node[T]
	prev *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
}

func NewList[T comparable]() *LinkedList[T] {
	list := &LinkedList[T]{head: nil}
	return list
}

func (l *LinkedList[T]) Add(item T) {
	if l.head == nil {
		l.head = &Node[T]{data: item, next: nil, prev: nil}
	} else {
		l.Insert(item, l.Count())
	}
}

func (l *LinkedList[T]) Insert(item T, index int) error {
	if index > l.Count() {
		return errors.New("out of range")
	} else {
		if l.head == nil {
			l.head = &Node[T]{data: item, next: nil, prev: nil}

		} else if index == l.Count() {
			tmp := l.head
			for tmp.next != nil {
				tmp = tmp.next
			}
			tmp.next = &Node[T]{data: item, next: nil, prev: tmp}

		} else if index == 0 {

			tmp := l.head

			tmp.prev = &Node[T]{data: item, prev: nil, next: tmp}

			l.head = tmp.prev

		} else {
			tmp := l.head

			for i := 0; i < index; i++ {
				tmp = tmp.next
			}
			prev := tmp.prev

			tmp.prev = &Node[T]{data: item, prev: prev, next: tmp}

			if prev != nil {
				prev.next = tmp.prev
			}

		}
	}
	return nil
}

func (l *LinkedList[T]) Remove(item T) error {
	if l.IsEmpty() {
		return errors.New("list is empty")
	}
	index, err := l.IndexOf(item)

	if err == nil {
		l.RemoveAt(index)

	}

	return nil

}

func (l *LinkedList[T]) RemoveAt(index int) error {
	if l.IsEmpty() {
		return errors.New("list is empty")
	} else if index >= l.Count() {
		return errors.New("out of range")
	}

	tmp := l.head

	for i := 0; i < index; i++ {
		tmp = tmp.next
	}

	prev := tmp.prev

	if prev != nil {
		prev.next = tmp.next
	} else {
		l.head = tmp.next
	}
	if tmp.next != nil {
		tmp.next.prev = prev
	}
	return nil
}

func (l *LinkedList[T]) Find(item T) bool {
	if l.IsEmpty() {
		return false
	}
	tmp := l.head

	for tmp != nil {
		if tmp.data == item {
			return true
		}
		tmp = tmp.next
	}
	return false
}

func (l *LinkedList[T]) ElementAt(index int) (T, error) {
	if l.IsEmpty() {
		var result T
		return result, errors.New("list is empty")
	} else if index >= l.Count() {
		var result T
		return result, errors.New("out of range")
	}

	tmp := l.head

	for i := 0; i < index; i++ {
		tmp = tmp.next
	}

	return tmp.data, nil
}

func (l *LinkedList[T]) IndexOf(item T) (int, error) {
	if l.IsEmpty() {
		var result int
		return result, errors.New("list is empty")
	}
	tmp := l.head

	counter := 0

	for tmp != nil {
		if tmp.data == item {
			return counter, nil
		}
		counter++
		tmp = tmp.next
	}
	var result int

	return result, errors.New("item not found")

}

func (l *LinkedList[T]) Clear() {
	l.head = nil
}

func (l *LinkedList[T]) First() (T, error) {
	if l.IsEmpty() {
		var result T

		return result, errors.New("list is empty")
	}

	return l.head.data, nil
}
func (l *LinkedList[T]) Last() (T, error) {
	if l.IsEmpty() {
		var result T

		return result, errors.New("list is empty")
	}

	data, _ := l.ElementAt(l.Count() - 1)

	return data, nil
}
func (l *LinkedList[T]) Count() int {
	counter := 0

	tmp := l.head

	for tmp != nil {
		counter++
		tmp = tmp.next
	}

	return counter
}
func (l *LinkedList[T]) IsEmpty() bool {
	return l.head == nil
}
func (l *LinkedList[T]) Print() {
	tmp := l.head

	for tmp != nil {
		fmt.Print(tmp.data, " | ")
		tmp = tmp.next
	}
	fmt.Println()
}
