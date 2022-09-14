package LinkedList

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Exit(exitVal)
}
func TestAdd(t *testing.T) {
	list := NewList[int]()
	list.Add(10)
	list.Add(15)
	list.Add(17)
	list.Add(19)

	expectedValue := 19
	actualValue, _ := list.Last()

	if expectedValue != actualValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestInsert(t *testing.T) {
	list := NewList[int]()
	list.Add(10)
	list.Add(15)
	list.Add(17)
	list.Insert(99, 1)

	expectedValue := 99
	actualValue, _ := list.ElementAt(1)

	if expectedValue != actualValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRemove(t *testing.T) {
	list := NewList[int]()
	list.Add(10)
	list.Add(15)
	list.Add(17)
	list.Remove(10)
	expectedValue := 15
	actualValue, _ := list.First()

	if expectedValue != actualValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRemoveAt(t *testing.T) {
	list := NewList[int]()
	list.Add(10)
	list.Add(15)
	list.Add(17)
	list.RemoveAt(0)
	expectedValue := 15
	actualValue, _ := list.First()

	if expectedValue != actualValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestFind(t *testing.T) {
	list := NewList[int]()
	list.Add(10)
	list.Add(15)
	list.Add(17)
	expectedValue := true
	actualValue := list.Find(17)

	if expectedValue != actualValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestElementAt(t *testing.T) {
	list := NewList[int]()
	list.Add(18)
	list.Add(24)
	list.Add(15)
	list.Add(96)
	var table = []struct {
		index         int
		expectedValue int
	}{
		{0, 18},
		{1, 24},
		{2, 15},
		{3, 96},
	}

	for _, value := range table {
		testName := fmt.Sprintf("Test input:%d", value.index)

		t.Run(testName, func(t *testing.T) {
			actualValue, _ := list.ElementAt(value.index)

			if actualValue != value.expectedValue {
				t.Errorf("Got %v expected %v", actualValue, value.expectedValue)
			}
		})
	}
}

func TestIndexOf(t *testing.T) {
	list := NewList[int]()
	list.Add(10)
	list.Add(15)
	list.Add(17)
	expectedValue := 2
	actualValue, _ := list.IndexOf(17)

	if expectedValue != actualValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestClear(t *testing.T) {
	list := NewList[int]()
	list.Add(10)
	list.Add(15)
	list.Add(17)
	list.Clear()
	expectedValue := true
	actualValue := list.IsEmpty()

	if expectedValue != actualValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestFirst(t *testing.T) {
	list := NewList[int]()
	list.Add(10)
	list.Add(15)
	list.Add(17)
	expectedValue := 10
	actualValue, _ := list.First()

	if expectedValue != actualValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestLast(t *testing.T) {
	list := NewList[int]()
	list.Add(10)
	list.Add(15)
	list.Add(17)
	expectedValue := 17
	actualValue, _ := list.Last()

	if expectedValue != actualValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestCount(t *testing.T) {
	list := NewList[int]()
	list.Add(10)
	list.Add(15)
	list.Add(17)
	expectedValue := 3
	actualValue := list.Count()

	if expectedValue != actualValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestIsEmpty(t *testing.T) {
	list := NewList[int]()
	list.Add(10)
	list.Add(15)
	list.Add(17)
	expectedValue := false
	actualValue := list.IsEmpty()

	if expectedValue != actualValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}
