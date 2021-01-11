package gcache

import (
	"hash/crc32"
	"sort"
	"strconv"
)

//Hash maps bytes to uint32
type Hash func(data []byte) uint32

//Map constains all hashed keys
type Map struct {
	hash     Hash           //采取依赖注入的方式，允许用于替换成自定义的 Hash 函数，也方便测试时替换，默认为 crc32.ChecksumIEEE 算法。
	replicas int            //虚拟节点倍数 replicas
	keys     []int          //哈希环 keys 所有虚拟节点key
	hashMap  map[int]string //键是虚拟节点的哈希值，值是真实节点的名称
}

//NewHashMap 实例化
func NewHashMap(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

//Add adds some keys to the hash.
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

//Get gets the closest item in the hash to the provided key.
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	hash := int(m.hash([]byte(key)))
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	return m.hashMap[m.keys[idx%len(m.keys)]]
}

//Hash1 自定义hash
func Hash1(data []byte) uint32 {
	var num uint32
	if len(data) == 2 {
		num = uint32(data[0]) | uint32(data[1])<<8
	} else {
		num = uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16
	}
	return num
}

//Hash2 自定义hash
func Hash2(data []byte) uint32 {
	i, _ := strconv.Atoi(string(data))
	num := uint32(i)
	return num

}
