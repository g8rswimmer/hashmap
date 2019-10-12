package hashmap

import (
	"errors"
	"fmt"
	"testing"
)

type strMockBenchmark struct{}

func (i *strMockBenchmark) Hash(obj interface{}) (uint64, error) {
	str, ok := obj.(string)
	if ok == false {
		return 0, errors.New("What")
	}
	h := 0
	for _, c := range str {
		h += int(c)
	}
	return uint64(h), nil
}

func (i *strMockBenchmark) Equals(obj1 interface{}, obj2 interface{}) (bool, error) {
	e1, ok := obj1.(string)
	if ok == false {
		return false, errors.New("What")
	}
	e2, ok := obj2.(string)
	if ok == false {
		return false, errors.New("What")
	}
	return e1 == e2, nil
}

func strBenchmarkHashMapPut(size int, b *testing.B) {
	m := &strMockBenchmark{}
	hm := NewHashMap(uint64(size), m, m)
	for n := 0; n < b.N; n++ {
		for i := 0; i < 10000; i++ {
			if err := hm.Put(fmt.Sprint(i*3), "yeah"); err != nil {
				b.Error(err)
			}
		}
	}
}

func BenchmarkStringHashMapPut10(b *testing.B) {
	strBenchmarkHashMapPut(10, b)
}

func BenchmarkStringHashMapPut100(b *testing.B) {
	strBenchmarkHashMapPut(100, b)
}

func BenchmarkStringHashMapPut1000(b *testing.B) {
	strBenchmarkHashMapPut(1000, b)
}

func BenchmarkStrHashMapGet(b *testing.B) {
	testStrHashMapGets := testHashMapGets{
		hm: &HashMap{
			table: []entries{
				entries{
					entry{key: "(", obj: "a"},
					entry{key: "<", obj: "b"},
					entry{key: "_", obj: "c"},
				},
				nil,
				entries{
					entry{key: "4", obj: "aa"},
					entry{key: "H", obj: "ab"},
				},
				entries{
					entry{key: "q", obj: "baa"},
					entry{key: "l", obj: "bab"},
					entry{key: "5", obj: "bac"},
					entry{key: "N", obj: "bad"},
				},
				entries{
					entry{key: "^", obj: "cbaa"},
				},
			},
			size:  5,
			hash:  &strMockBenchmark{},
			equal: &strMockBenchmark{},
		},
		keys:   []interface{}{"^", "(", "5", "H", "q", "N", "<", "4", "_"},
		values: []interface{}{"cbaa", "a", "bac", "ab", "baa", "bad", "b", "aa", "c"},
	}

	for n := 0; n < b.N; n++ {
		for idx, k := range testStrHashMapGets.keys {
			if v, err := testStrHashMapGets.hm.Get(k); err == nil {
				if v != testStrHashMapGets.values[idx] {
					b.Errorf("BenchmarkIntHashMapGet = %v, want %v", v, testStrHashMapGets.values[idx])
				}
			} else {
				b.Error(err)
			}
		}
	}
}
