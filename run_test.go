package main

import (
	"reflect"
	"testing"

	"github.com/liuyong-go/yong/core/gcache"
)

//go test -v -run TestGetter run_test.go
func TestGetter(t *testing.T) {
	var f gcache.Getter = gcache.GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})
	expect := []byte("key1")
	if v, _ := f.Get("key"); !reflect.DeepEqual(v, expect) {
		t.Error("call back fail")
	}
	t.Log("test")
}

//go test -v -run TestConsistent run_test.go
func TestConsistent(t *testing.T) {
	var hash = gcache.NewHashMap(2, gcache.Hash2)
	var keys = []string{"1", "8", "13", "17", "21"}
	hash.Add(keys...)
	value := hash.Get("3")
	t.Log("test", value)
}
