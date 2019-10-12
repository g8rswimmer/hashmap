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

func intBenchmarkHashMapRand(size int, b *testing.B) {
	m := &intMockBenchmark{}
	hm := NewHashMap(uint64(size), m, m)
	for n := 0; n < b.N; n++ {
		for i := 0; i < 10000; i++ {
			k := 3
			if err := hm.Put(k, "yeah"); err != nil {
				b.Error(err)
			}
			k *= 31
		}
	}
}

func BenchmarkIntHashMap10(b *testing.B) {
	intBenchmarkHashMapRand(10, b)
}

func BenchmarkIntHashMap100(b *testing.B) {
	intBenchmarkHashMapRand(100, b)
}

func BenchmarkIntHashMap1000(b *testing.B) {
	intBenchmarkHashMapRand(1000, b)
}
