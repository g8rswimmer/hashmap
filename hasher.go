package hashmap

import (
	"fmt"
	"math"
)

// Hasher is an interface that can be used to create a 64 bit hash
type Hasher interface {
	Hash(interface{}) (uint64, error)
}

// DefaultHasher can be used to create hashes for data types,
// integers, strings, floats and booleans.
var DefaultHasher = defaultHasher{}

type defaultHasher struct{}

func (d defaultHasher) Hash(obj interface{}) (uint64, error) {
	switch obj.(type) {
	case int, int16, int32, int64, int8, uint, uint16, uint32, uint64, uint8:
		return d.intHash(obj)
	case string:
		str, _ := obj.(string)
		return d.stringHash(str)
	case float32, float64:
		return d.floatHash(obj)
	case bool:
		if b, _ := obj.(bool); b {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("hasher int: type not supported %T", obj)
	}
}

func (d defaultHasher) intHash(obj interface{}) (uint64, error) {

	switch obj.(type) {
	case int:
		num, _ := obj.(int)
		return uint64(num), nil
	case int16:
		num, _ := obj.(int16)
		return uint64(num), nil
	case int32:
		num, _ := obj.(int32)
		return uint64(num), nil
	case int64:
		num, _ := obj.(int64)
		return uint64(num), nil
	case int8:
		num, _ := obj.(int8)
		return uint64(num), nil
	case uint:
		num, _ := obj.(uint)
		return uint64(num), nil
	case uint16:
		num, _ := obj.(uint16)
		return uint64(num), nil
	case uint32:
		num, _ := obj.(uint32)
		return uint64(num), nil
	case uint64:
		num, _ := obj.(uint64)
		return uint64(num), nil
	case uint8:
		num, _ := obj.(uint8)
		return uint64(num), nil
	default:
		return 0, fmt.Errorf("hasher int: type not supported %T", obj)
	}
}
func (d defaultHasher) stringHash(obj string) (uint64, error) {
	var shift uint
	var h uint64
	for _, c := range obj {
		h += (uint64(c) << shift)
		shift++
		if shift >= 64 {
			shift = 0
		}
	}
	return h, nil
}

func (d defaultHasher) floatHash(obj interface{}) (uint64, error) {
	switch obj.(type) {
	case float32:
		num, _ := obj.(float32)
		return uint64(math.Float32bits(num)), nil
	case float64:
		num, _ := obj.(float64)
		return math.Float64bits(num), nil
	default:
		return 0, fmt.Errorf("hasher int: type not supported %T", obj)
	}
}
