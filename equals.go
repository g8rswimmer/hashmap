package hashmap

import "reflect"

// Equaler is the interface to compare to objects
type Equaler interface {
	Equals(interface{}, interface{}) (bool, error)
}

// DefaultEqualer is the default equalers
var DefaultEqualer = defaultEqualer{}

type defaultEqualer struct{}

func (d defaultEqualer) Equals(obj1, obj2 interface{}) (bool, error) {
	eq := reflect.DeepEqual(obj1, obj2)
	return eq, nil
}
