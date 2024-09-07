package test

import (
	collections "github.com/mobaijaavaer/golang-tools/collections/impl"
	"testing"
)

// 测试 NewArrayList 函数
func TestNewArrayList(t *testing.T) {
	list := collections.NewArrayList[int]()
	if list == nil || len(list.GetData()) > 0 {
		t.Errorf("NewArrayList() did not create an empty ArrayList")
	}
}

// 测试 NewArrayLists 函数
func TestNewArrayLists(t *testing.T) {
	source := []int{1, 2, 3}
	list := collections.NewArrayLists(source)
	if list == nil || list.GetData() == nil || len(list.GetData()) != len(source) {
		t.Errorf("NewArrayLists() did not create an ArrayList with the source GetData()")
	}
}

// 测试 Add 方法
func TestAdd(t *testing.T) {
	list := collections.NewArrayList[int]()
	list.Add(1)
	if list.GetData() == nil || len(list.GetData()) != 1 || list.GetData()[0] != 1 {
		t.Errorf("Add() did not add an element to the ArrayList")
	}
}

// 测试 Get 方法
func TestGet(t *testing.T) {
	list := collections.NewArrayList[int]()
	list.Add(1)
	if list.Get(0) != 1 {
		t.Errorf("Get() did not return the correct element from the ArrayList")
	}
}

// 测试 AddIf 方法
func TestAddIf(t *testing.T) {
	list := collections.NewArrayList[int]()
	predicate := func(i int) bool { return i > 0 }
	list.AddIf(1, predicate)
	list.AddIf(-1, predicate)
	if len(list.GetData()) != 1 || list.GetData()[0] != 1 {
		t.Errorf("AddIf() did not add an element to the ArrayList based on predicate")
	}
}

// 测试 AddAll 方法
func TestAddAll(t *testing.T) {
	list := collections.NewArrayList[int]()
	items := []int{1, 2, 3}
	list.AddAll(items)
	if len(list.GetData()) != len(items) {
		t.Errorf("AddAll() did not add all elements to the ArrayList")
	}
}

// 测试 Clear 方法
func TestClear(t *testing.T) {
	list := collections.NewArrayList[int]()
	list.Add(1)
	list.Clear()
	if list.GetData() == nil || len(list.GetData()) != 0 {
		t.Errorf("Clear() did not clear the ArrayList")
	}
}

// 测试 Contains 方法
func TestContains(t *testing.T) {
	list := collections.NewArrayList[int]()
	list.Add(1)
	if !list.Contains(1) {
		t.Errorf("Contains() did not find an element in the ArrayList")
	}
}

// 测试 ContainsAll 方法
func TestContainsAll(t *testing.T) {
	list := collections.NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	items := []int{1, 2}
	if !list.ContainsAll(items) {
		t.Errorf("ContainsAll() did not find all elements in the ArrayList")
	}
}

// 测试 IsEmpty 方法
func TestIsEmpty(t *testing.T) {
	list := collections.NewArrayList[int]()
	if !list.IsEmpty() {
		t.Errorf("IsEmpty() did not return true for an empty ArrayList")
	}
}

// 测试 Remove 方法
func TestRemove(t *testing.T) {
	list := collections.NewArrayList[int]()
	list.Add(1)
	list.Remove(1)
	if len(list.GetData()) != 0 {
		t.Errorf("Remove() did not remove an element from the ArrayList")
	}
}

// 测试 Size 方法
func TestSize(t *testing.T) {
	list := collections.NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	if list.Size() != 2 {
		t.Errorf("Size() did not return the correct size of the ArrayList")
	}
}

// 测试 ToString 方法
func TestToString(t *testing.T) {
	list := collections.NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	expected := "[1 2]"
	if list.ToString() != expected {
		t.Errorf("ToString() did not return the correct string representation of the ArrayList")
	}
}

// 测试 Distinct 方法
func TestDistinct(t *testing.T) {
	list := collections.NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(1)
	list.Distinct()
	if len(list.GetData()) != 2 {
		t.Errorf("Distinct() did not remove duplicate elements from the ArrayList")
	}
}

// 测试 Sort 方法
func TestSort(t *testing.T) {
	list := collections.NewArrayList[int]()
	list.Add(3)
	list.Add(1)
	list.Add(2)
	list.Sort(func(i int) int { return i })
	if list.GetData()[0] != 1 || list.GetData()[1] != 2 || list.GetData()[2] != 3 {
		t.Errorf("Sort() did not sort the ArrayList correctly")
	}
}

// 测试 NewEmpty 方法
func TestNewEmpty(t *testing.T) {
	list := collections.NewArrayList[int]()
	newList := list.NewEmpty()
	if newList == nil || len(newList.GetData()) > 0 {
		t.Errorf("NewEmpty() did not create a new empty ArrayList")
	}
}
