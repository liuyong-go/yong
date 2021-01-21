package main

import (
	"reflect"
	"testing"

	"github.com/liuyong-go/yong/core/gcache"
	"github.com/liuyong-go/yong/core/georm"
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
func TestDb(t *testing.T) {
	engine, _ := georm.NewEngine("mysql", "liuyong:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	defer engine.Close()
	s := engine.NewSession()
	result, err := s.Raw("insert into book values (?,?)", "test1", "10").Exec()
	if err != nil {
		georm.LogDBErr(err)
	}
	count, _ := result.RowsAffected()
	t.Log("count", count)

}
