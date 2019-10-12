package hashmap

import (
	"errors"
	"testing"
)

type intMockBenchmark struct{}

func (i *intMockBenchmark) Hash(obj interface{}) (uint64, error) {
	h, ok := obj.(int)
	if ok == false {
		return 0, errors.New("What")
	}
	return uint64(h), nil
}

func (i *intMockBenchmark) Equals(obj1 interface{}, obj2 interface{}) (bool, error) {
	e1, ok := obj1.(int)
	if ok == false {
		return false, errors.New("What")
	}
	e2, ok := obj2.(int)
	if ok == false {
		return false, errors.New("What")
	}
	return e1 == e2, nil
}

func intBenchmarkHashMapPut(size int, b *testing.B) {
	m := &intMockBenchmark{}
	for n := 0; n < b.N; n++ {
		hm := NewHashMap(uint64(size), m, m)
		for i := 0; i < 10000; i++ {
			k := 3
			if err := hm.Put(k, "yeah"); err != nil {
				b.Error(err)
			}
			k *= 31
		}
	}
}

func BenchmarkIntHashMapPut10(b *testing.B) {
	intBenchmarkHashMapPut(10, b)
}

func BenchmarkIntHashMapPut100(b *testing.B) {
	intBenchmarkHashMapPut(100, b)
}

func BenchmarkIntHashMapPut1000(b *testing.B) {
	intBenchmarkHashMapPut(1000, b)
}

func BenchmarkIntHashMapGet(b *testing.B) {
	testIntHashMapGets := testHashMapGets{
		hm: &HashMap{
			table: []entries{
				entries{
					entry{key: 0, obj: "a"},
					entry{key: 10, obj: "b"},
					entry{key: 30, obj: "c"},
				},
				nil,
				entries{
					entry{key: 12, obj: "aa"},
					entry{key: 2, obj: "ab"},
				},
				entries{
					entry{key: 33, obj: "baa"},
					entry{key: 93, obj: "bab"},
					entry{key: 43, obj: "bac"},
					entry{key: 73, obj: "bad"},
				},
				entries{
					entry{key: 84, obj: "cbaa"},
				},
			},
			size:  5,
			hash:  &intMockBenchmark{},
			equal: &intMockBenchmark{},
		},
		keys:   []interface{}{84, 0, 43, 2, 33, 73, 10, 12, 30},
		values: []interface{}{"cbaa", "a", "bac", "ab", "baa", "bad", "b", "aa", "c"},
	}

	for n := 0; n < b.N; n++ {
		for idx, k := range testIntHashMapGets.keys {
			if v, err := testIntHashMapGets.hm.Get(k); err == nil {
				if v != testIntHashMapGets.values[idx] {
					b.Errorf("BenchmarkIntHashMapGet = %v, want %v", v, testIntHashMapGets.values[idx])
				}
			} else {
				b.Error(err)
			}
		}
	}
}
