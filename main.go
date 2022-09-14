package main

import (
	"fmt"

	"github.com/muhammetcagatay/go-linked-list/LinkedList"
)

func main() {
	list := LinkedList.NewList[int]()

	list.Add(5)
	list.Add(7)
	list.Add(8)

	list.Remove(5)

	list.Print()
	fmt.Println("Count:", list.Count())
}
