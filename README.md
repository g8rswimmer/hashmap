# hashmap
This is a `golang` library for a hash map.

## Getting Started 

### Installing
To start using `hashmap`, install `GO` and run `go get`
```
go get -u github.com/g8rswimmer/hashmap
```
This will retrive the package

## Usage 
The user is free to implement the `Hasher` and/or `Equaler` interfaces, or use the default ones provided.

```
    hm := hashmap.NewHashMap(uint64(100), DefaultHasher, DefaultEqualer)
    hm.Put("hello", "this is an example")
    value := hm.Get("hello")
    fmt.Printf("%v\n", value)
```

## Benchmarking

### Custom Int Hasher and Equaler
```
goos: darwin
goarch: amd64
pkg: github.com/g8rswimmer/hashmap

BenchmarkBaseIntPut-8      	                   20000	     99603 ns/op	  122106 B/op	      39 allocs/op
BenchmarkBaseIntGet-8      	                10000000	       192 ns/op	       0 B/op	       0 allocs/op

BenchmarkIntHashMapPut10-8               	    3000	    372615 ns/op	   80272 B/op	   10002 allocs/op
BenchmarkIntHashMapPut100-8              	    5000	    376295 ns/op	   82720 B/op	   10002 allocs/op
BenchmarkIntHashMapPut1000-8             	    5000	    385656 ns/op	  104608 B/op	   10002 allocs/op

BenchmarkIntHashMapGet-8                 	 5000000	       303 ns/op	       0 B/op	       0 allocs/op
```
### Custom String Hasher and Equaler
```
goos: darwin
goarch: amd64
pkg: github.com/g8rswimmer/hashmap

BenchmarkBaseStringPut-8   	                30000	     42470 ns/op	    5808 B/op	    1002 allocs/op
BenchmarkBaseStringGet-8   	                5000000	       260 ns/op	       0 B/op	       0 allocs/op

BenchmarkStringHashMapPut10-8            	      30	  49736044 ns/op	 1352660 B/op	   30120 allocs/op
BenchmarkStringHashMapPut100-8           	      30	  39182395 ns/op	 1339108 B/op	   30292 allocs/op
BenchmarkStringHashMapPut1000-8          	      30	  38746536 ns/op	 1361063 B/op	   30292 allocs/op

BenchmarkStrHashMapGet-8                 	 5000000	       383 ns/op	       0 B/op	       0 allocs/op
```
### Defualt Int Hasher and Equaler
```
goos: darwin
goarch: amd64
pkg: github.com/g8rswimmer/hashmap

BenchmarkBaseIntPut-8      	                20000	     99603 ns/op	  122106 B/op	      39 allocs/op
BenchmarkBaseIntGet-8      	                10000000	       192 ns/op	       0 B/op	       0 allocs/op

BenchmarkDefaultIntHashMapPut10-8        	    1000	   1154594 ns/op	   80272 B/op	   10002 allocs/op
BenchmarkDefaultIntHashMapPut100-8       	    1000	   1263153 ns/op	   82723 B/op	   10002 allocs/op
BenchmarkDefaultIntHashMapPut1000-8      	    2000	   1208452 ns/op	  104608 B/op	   10002 allocs/op

BenchmarkDefaultIntHashMapGet-8          	 1000000	      1610 ns/op	       0 B/op	       0 allocs/op
```
### Default String Hasher and Equaler
```
goos: darwin
goarch: amd64
pkg: github.com/g8rswimmer/hashmap

BenchmarkBaseStringPut-8   	                30000	     42470 ns/op	    5808 B/op	    1002 allocs/op
BenchmarkBaseStringGet-8   	                5000000	       260 ns/op	       0 B/op	       0 allocs/op

BenchmarkDefaultStringHashMapPut10-8     	       3	 401019598 ns/op	 1180477 B/op	   30116 allocs/op
BenchmarkDefaultStringHashMapPut100-8    	      30	  40128445 ns/op	 1139150 B/op	   30802 allocs/op
BenchmarkDefaultStringHashMapPut1000-8   	     100	  14235148 ns/op	 1284724 B/op	   32357 allocs/op

BenchmarkDefaultStrHashMapGet-8          	 1000000	      1618 ns/op	       0 B/op	       0 allocs/op
```