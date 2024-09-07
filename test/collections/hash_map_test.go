package test

import (
	collections "github.com/mobaijaavaer/golang-tools/collections/impl"
	"reflect"
	"testing"
)

// TestNewHashMap 测试NewHashMap函数是否正确地创建了一个新的HashMap实例。
func TestNewHashMap(t *testing.T) {
	hashMap := collections.NewHashMap[string, int]()
	if hashMap == nil || reflect.TypeOf(hashMap).String() != "*impl.HashMap[string,int]" {
		t.Errorf("collections.NewHashMap should return a non-nil HashMap instance.")
	}
}

// TestHashMapPut 测试Put方法是否能够正确地插入键值对。
func TestHashMapPut(t *testing.T) {
	hashMap := collections.NewHashMap[string, int]()
	hashMap.Put("key1", 1)
	if !hashMap.ContainsKey("key1") {
		t.Errorf("Key 'key1' should be present after Put.")
	}
}

// TestHashMapGet 测试Get方法是否能正确获取键对应的值。
func TestHashMapGet(t *testing.T) {
	hashMap := collections.NewHashMap[string, int]()
	hashMap.Put("key2", 2)
	value := hashMap.Get("key2")
	if value != 2 {
		t.Errorf("Expected value 2, got %v", value)
	}
}

// TestHashMapGetOrDefault 测试GetOrDefault方法是否能正确获取键对应的值或默认值。
func TestHashMapGetOrDefault(t *testing.T) {
	hashMap := collections.NewHashMap[string, int]()
	defaultValue := 0
	value := hashMap.GetOrDefault("nonExistentKey", defaultValue)
	if value != defaultValue {
		t.Errorf("Expected default value %v, got %v", defaultValue, value)
	}
}

// TestHashMapContainsKey 测试ContainsKey方法是否能正确判断键是否存在。
func TestHashMapContainsKey(t *testing.T) {
	hashMap := collections.NewHashMap[string, int]()
	hashMap.Put("key3", 3)
	exists := hashMap.ContainsKey("key3")
	if !exists {
		t.Errorf("Key 'key3' should exist.")
	}
}

// TestHashMapRemove 测试Remove方法是否能正确移除键值对。
func TestHashMapRemove(t *testing.T) {
	hashMap := collections.NewHashMap[string, int]()
	hashMap.Put("key4", 4)
	hashMap.Remove("key4")
	if hashMap.ContainsKey("key4") {
		t.Errorf("Key 'key4' should not exist after removal.")
	}
}

// TestHashMapReplace 测试Replace方法是否能正确替换键对应的值。
func TestHashMapReplace(t *testing.T) {
	hashMap := collections.NewHashMap[string, int]()
	hashMap.Put("key5", 5)
	hashMap.Replace("key5", 6)
	value := hashMap.Get("key5")
	if value != 6 {
		t.Errorf("Expected value 6 after replacement, got %v", value)
	}
}

// TestHashMapSize 测试Size方法是否能正确返回键值对的数量。
func TestHashMapSize(t *testing.T) {
	hashMap := collections.NewHashMap[string, int]()
	hashMap.Put("key6", 6)
	hashMap.Put("key7", 7)
	size := hashMap.Size()
	if size != 2 {
		t.Errorf("Expected size 2, got %d", size)
	}
}

// TestHashMapIsEmpty 测试IsEmpty方法是否能正确判断HashMap是否为空。
func TestHashMapIsEmpty(t *testing.T) {
	hashMap := collections.NewHashMap[string, int]()
	isEmpty := hashMap.IsEmpty()
	if !isEmpty {
		t.Errorf("Expected hashMap to be empty.")
	}
	hashMap.Put("key8", 8)
	isEmpty = hashMap.IsEmpty()
	if isEmpty {
		t.Errorf("Expected hashMap to be non-empty after insertion.")
	}
}

// TestHashMapClear 测试Clear方法是否能正确清空HashMap。
func TestHashMapClear(t *testing.T) {
	hashMap := collections.NewHashMap[string, int]()
	hashMap.Put("key9", 9)
	hashMap.Clear()
	if !hashMap.IsEmpty() {
		t.Errorf("Expected hashMap to be empty after Clear.")
	}
}

// TestHashMapGetKeys 测试GetKeys方法是否能正确获取所有的键。
func TestHashMapGetKeys(t *testing.T) {
	hashMap := collections.NewHashMap[string, int]()
	hashMap.Put("key10", 10)
	keys := hashMap.GetKeys()
	if keys.Size() != 1 {
		t.Errorf("Expected 1 key, got %d keys", keys.Size())
	}
}

// TestHashMapGetValues 测试GetValues方法是否能正确获取所有的值。
func TestHashMapGetValues(t *testing.T) {
	hashMap := collections.NewHashMap[string, int]()
	hashMap.Put("key11", 11)
	values := hashMap.GetValues()
	if values.Size() != 1 {
		t.Errorf("Expected 1 value, got %d values", values.Size())
	}

}
