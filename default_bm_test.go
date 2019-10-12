package hashmap

import (
	"fmt"
	"testing"
)

func defaultintBenchmarkHashMapRand(size int, b *testing.B) {
	hm := NewHashMap(uint64(size), DefaultHasher, DefaultEqualer)
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

func BenchmarkDefaultIntHashMap10(b *testing.B) {
	defaultintBenchmarkHashMapRand(10, b)
}

func BenchmarkDefaultIntHashMap100(b *testing.B) {
	defaultintBenchmarkHashMapRand(100, b)
}

func BenchmarkDefaultIntHashMap1000(b *testing.B) {
	defaultintBenchmarkHashMapRand(1000, b)
}

func defaultStrBenchmarkHashMapRand(size int, b *testing.B) {
	hm := NewHashMap(uint64(size), DefaultHasher, DefaultEqualer)
	for n := 0; n < b.N; n++ {
		for i := 0; i < 10000; i++ {
			if err := hm.Put(fmt.Sprint(i*3), "yeah"); err != nil {
				b.Error(err)
			}
		}
	}
}

func BenchmarkDefaultStringHashMap10(b *testing.B) {
	defaultStrBenchmarkHashMapRand(10, b)
}

func BenchmarkDefaultStringHashMap100(b *testing.B) {
	defaultStrBenchmarkHashMapRand(100, b)
}

func BenchmarkDefaultStringHashMap1000(b *testing.B) {
	defaultStrBenchmarkHashMapRand(1000, b)
}
