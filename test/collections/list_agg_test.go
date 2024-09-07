package test

import (
	collections2 "github.com/mobaijaavaer/golang-tools/collections"
	collections "github.com/mobaijaavaer/golang-tools/collections/impl"
	"testing"
)

// 测试 Filter 函数
func TestFilter(t *testing.T) {
	list := collections.NewArrayList[int]()
	list.AddAll([]int{1, 2, 3, 4, 5})
	filtered := collections2.Filter[int](list, func(i int) bool {
		return i%2 == 0
	})
	expected := collections.NewArrayList[int]()
	expected.AddAll([]int{2, 4})
	if len(filtered.GetData()) != len(expected.GetData()) {
		t.Errorf("Filter returned a list of incorrect length: got %d, want %d", len(filtered.GetData()), len(expected.GetData()))
	}
	for i, v := range filtered.GetData() {
		if v != expected.GetData()[i] {
			t.Errorf("Filter returned incorrect value at index %d: got %d, want %d", i, v, expected.GetData()[i])
		}
	}
}

// 测试 Reduce 函数
func TestReduce(t *testing.T) {
	list := collections.NewArrayList[int]()
	list.AddAll([]int{1, 2, 3, 4, 5})
	sum := collections2.Reduce[int, int](list, 0, func(acc, i int) int {
		return acc + i
	})
	if sum != 15 {
		t.Errorf("Reduce returned incorrect sum: got %d, want %d", sum, 15)
	}
}

// 测试 ToMap 函数
func TestToMap(t *testing.T) {
	list := collections.NewArrayList[int]()
	list.AddAll([]int{1, 2, 3, 4, 5})
	mapped := collections2.ToMap[int, int](list, func(i int) (int, int) {
		return i, i * 2
	})
	for i, v := range mapped {
		if v != i*2 {
			t.Errorf("ToMap returned incorrect value for key %d: got %d, want %d", i, v, i*2)
		}
	}
}

// 测试 GroupingBy 函数
func TestGroupingBy(t *testing.T) {
	list := collections.NewArrayList[int]()
	list.AddAll([]int{1, 2, 3, 4, 5})
	grouped := collections2.GroupingBy[int, int](list, func(i int) int {
		return i / 2
	})
	// 验证分组结果
	expectedGroups := map[int][]int{
		0: {1},
		1: {2, 3},
		2: {4, 5},
	}
	for key, groupList := range grouped {
		expectedValues := expectedGroups[key]
		actualValues := groupList.GetData()

		if len(actualValues) != len(expectedValues) {
			t.Errorf("GroupingBy returned incorrect number of elements for key %d: got %d, want %d", key, len(actualValues), len(expectedValues))
		}

		for i, v := range actualValues {
			if v != expectedValues[i] {
				t.Errorf("GroupingBy returned incorrect value for key %d at index %d: got %d, want %d", key, i, v, expectedValues[i])
			}
		}
	}
}
