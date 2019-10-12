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

func strBenchmarkHashMapRand(size int, b *testing.B) {
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

func BenchmarkStringHashMap10(b *testing.B) {
	strBenchmarkHashMapRand(10, b)
}

func BenchmarkStringHashMap100(b *testing.B) {
	strBenchmarkHashMapRand(100, b)
}

func BenchmarkStringHashMap1000(b *testing.B) {
	strBenchmarkHashMapRand(1000, b)
}
