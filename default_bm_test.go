package hashmap

import (
	"fmt"
	"testing"
)

func defaultintBenchmarkHashMapPut(size int, b *testing.B) {
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

func BenchmarkDefaultIntHashMapPut10(b *testing.B) {
	defaultintBenchmarkHashMapPut(10, b)
}

func BenchmarkDefaultIntHashMapPut100(b *testing.B) {
	defaultintBenchmarkHashMapPut(100, b)
}

func BenchmarkDefaultIntHashMapPut1000(b *testing.B) {
	defaultintBenchmarkHashMapPut(1000, b)
}

func defaultStrBenchmarkHashMapPut(size int, b *testing.B) {
	hm := NewHashMap(uint64(size), DefaultHasher, DefaultEqualer)
	for n := 0; n < b.N; n++ {
		for i := 0; i < 10000; i++ {
			if err := hm.Put(fmt.Sprint(i*3), "yeah"); err != nil {
				b.Error(err)
			}
		}
	}
}

func BenchmarkDefaultStringHashMapPut10(b *testing.B) {
	defaultStrBenchmarkHashMapPut(10, b)
}

func BenchmarkDefaultStringHashMapPut100(b *testing.B) {
	defaultStrBenchmarkHashMapPut(100, b)
}

func BenchmarkDefaultStringHashMapPut1000(b *testing.B) {
	defaultStrBenchmarkHashMapPut(1000, b)
}

func BenchmarkDefaultIntHashMapGet(b *testing.B) {
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
			hash:  DefaultHasher,
			equal: DefaultEqualer,
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

func BenchmarkDefaultStrHashMapGet(b *testing.B) {
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
			hash:  DefaultHasher,
			equal: DefaultEqualer,
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
