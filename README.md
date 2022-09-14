# Linked List
A linked list is a common data structure made of one or more than one nodes. Each node contains a value and a pointer to the previous/next node forming the chain-like structure. These nodes are stored randomly in the system's memory, which improves its space complexity compared to the array.

## Installation
```
go get github.com/muhammetcagatay/go-linked-list
```

## Using
```
package main

import (
	"github.com/muhammetcagatay/go-linked-list/LinkedList"
)

func main() {
	list := LinkedList.NewList[int]()

	list.Add(5)
	list.Add(7)
	list.Add(8)

	list.Print()
}
```

## Test
```
cd LinkedList
go test -v
```
