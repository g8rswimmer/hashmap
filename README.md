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

BenchmarkIntHashMapPut10-8               	    5000	    350583 ns/op	  
BenchmarkIntHashMapPut100-8              	    5000	    347208 ns/op	  
BenchmarkIntHashMapPut1000-8             	    5000	    348855 ns/op	  

BenchmarkIntHashMapGet-8                 	 5000000	       279 ns/op	  
```
### Custom String Hasher and Equaler
```
goos: darwin
goarch: amd64
pkg: github.com/g8rswimmer/hashmap

BenchmarkStringHashMapPut10-8            	      30	  46066780 ns/op	  
BenchmarkStringHashMapPut100-8           	      50	  35299887 ns/op	  
BenchmarkStringHashMapPut1000-8          	      50	  35104093 ns/op	  

BenchmarkStrHashMapGet-8                 	 5000000	       362 ns/op	  
```
### Defualt Int Hasher and Equaler
```
goos: darwin
goarch: amd64
pkg: github.com/g8rswimmer/hashmap
BenchmarkDefaultIntHashMapPut10-8        	    2000	   1136376 ns/op	   
BenchmarkDefaultIntHashMapPut100-8       	    2000	   1128874 ns/op	   
BenchmarkDefaultIntHashMapPut1000-8      	    2000	   1162005 ns/op	   

BenchmarkDefaultIntHashMapGet-8          	 1000000	      1529 ns/op	   
```
### Default String Hasher and Equaler
```
goos: darwin
goarch: amd64
pkg: github.com/g8rswimmer/hashmap
BenchmarkDefaultStringHashMapPut10-8                 100         345026505 ns/op
BenchmarkDefaultStringHashMapPut100-8               1000          37173400 ns/op
BenchmarkDefaultStringHashMapPut1000-8              2000          13320337 ns/op

BenchmarkDefaultIntHashMapGet-8                     20000000      1490 ns/op	       
```