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

BenchmarkIntHashMapPut10-8               	    5000	    350583 ns/op	   80000 B/op	   10000 allocs/op
BenchmarkIntHashMapPut100-8              	    5000	    347208 ns/op	   80000 B/op	   10000 allocs/op
BenchmarkIntHashMapPut1000-8             	    5000	    348855 ns/op	   80005 B/op	   10000 allocs/op

BenchmarkIntHashMapGet-8                 	 5000000	       279 ns/op	       0 B/op	       0 allocs/op
```
### Custom String Hasher and Equaler
```
goos: darwin
goarch: amd64
pkg: github.com/g8rswimmer/hashmap

BenchmarkStringHashMapPut10-8            	      30	  46066780 ns/op	  354573 B/op	   30003 allocs/op
BenchmarkStringHashMapPut100-8           	      50	  35299887 ns/op	  340529 B/op	   30005 allocs/op
BenchmarkStringHashMapPut1000-8          	      50	  35104093 ns/op	  340964 B/op	   30005 allocs/op

BenchmarkStrHashMapGet-8                 	 5000000	       362 ns/op	       0 B/op	       0 allocs/op
```
### Defualt Int Hasher and Equaler
```
goos: darwin
goarch: amd64
pkg: github.com/g8rswimmer/hashmap
BenchmarkDefaultIntHashMapPut10-8        	    2000	   1136376 ns/op	   80001 B/op	   10000 allocs/op
BenchmarkDefaultIntHashMapPut100-8       	    2000	   1128874 ns/op	   80001 B/op	   10000 allocs/op
BenchmarkDefaultIntHashMapPut1000-8      	    2000	   1162005 ns/op	   80012 B/op	   10000 allocs/op

BenchmarkDefaultIntHashMapGet-8          	 1000000	      1529 ns/op	       0 B/op	       0 allocs/op
```
### Default String Hasger and Equaler
```
goos: darwin
goarch: amd64
pkg: github.com/g8rswimmer/hashmap
BenchmarkDefaultStringHashMapPut10-8     	       3	 379296362 ns/op	  607112 B/op	   30039 allocs/op
BenchmarkDefaultStringHashMapPut100-8    	      50	  37249455 ns/op	  336513 B/op	   30015 allocs/op
BenchmarkDefaultStringHashMapPut1000-8   	     100	  14145849 ns/op	  329786 B/op	   30023 allocs/op

BenchmarkDefaultStrHashMapGet-8          	 1000000	      1540 ns/op	       0 B/op	       0 allocs/op
```