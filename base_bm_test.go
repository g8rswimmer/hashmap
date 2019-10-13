package hashmap

import "testing"

var benckmarkIntEntries = []entry{
	{key: 1, obj: "a"},
	{key: 11, obj: "aa"},
	{key: 111, obj: "aaa"},
	{key: 1111, obj: "aaaa"},
	{key: 11111, obj: "aaaaa"},
	{key: 111111, obj: "aaaaaa"},
	{key: 1111111, obj: "aaaaaaa"},
	{key: 111111111, obj: "aaaaaaaa"},
	{key: 1111111111, obj: "aaaaaaaaaa"},
	{key: 2, obj: "ab"},
	{key: 22, obj: "abb"},
	{key: 222, obj: "abbb"},
	{key: 2222, obj: "abbbb"},
	{key: 22222, obj: "abbbbb"},
	{key: 2222222, obj: "abbbbbb"},
	{key: 222222222, obj: "abbbbbbbbb"},
}

var benckmarkStringEntries = []entry{
	{key: "1", obj: "a"},
	{key: "11", obj: "aa"},
	{key: "111", obj: "aaa"},
	{key: "1111", obj: "aaaa"},
	{key: "11111", obj: "aaaaa"},
	{key: "111111", obj: "aaaaaa"},
	{key: "1111111", obj: "aaaaaaa"},
	{key: "111111111", obj: "aaaaaaaa"},
	{key: "1111111111", obj: "aaaaaaaaaa"},
	{key: "2", obj: "ab"},
	{key: "22", obj: "abb"},
	{key: "222", obj: "abbb"},
	{key: "2222", obj: "abbbb"},
	{key: "22222", obj: "abbbbb"},
	{key: "2222222", obj: "abbbbbb"},
	{key: "222222222", obj: "abbbbbbbbb"},
}

func BenchmarkBaseIntPut(b *testing.B) {
	for n := 0; n < b.N; n++ {
		k := 1
		hm := map[int]string{}
		for i := 0; i < 1000; i++ {
			hm[k] = "yep"
			k *= 3
		}
	}
}

func BenchmarkBaseIntGet(b *testing.B) {
	hm := map[int]string{}
	for _, e := range benckmarkIntEntries {
		hm[e.key.(int)] = e.obj.(string)
	}
	for n := 0; n < b.N; n++ {
		for _, e := range benckmarkIntEntries {
			if v, ok := hm[e.key.(int)]; ok {
				if v != e.obj {
					b.Errorf("BenchmarkGet = %v, want %v", v, e.obj)
				}
			} else {
				b.Errorf("not present: %d", e.key)
			}
		}
	}

}

func BenchmarkBaseStringPut(b *testing.B) {
	for n := 0; n < b.N; n++ {
		k := 1
		hm := map[string]string{}
		for i := 0; i < 1000; i++ {
			hm[string(k)] = "yep"
			k *= 3
		}
	}
}

func BenchmarkBaseStringGet(b *testing.B) {
	hm := map[string]string{}
	for _, e := range benckmarkStringEntries {
		hm[e.key.(string)] = e.obj.(string)
	}
	for n := 0; n < b.N; n++ {
		for _, e := range benckmarkStringEntries {
			if v, ok := hm[e.key.(string)]; ok {
				if v != e.obj {
					b.Errorf("BenchmarkGet = %v, want %v", v, e.obj)
				}
			} else {
				b.Errorf("not present: %d", e.key)
			}
		}
	}

}
