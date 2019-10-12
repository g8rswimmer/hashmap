package hashmap

import "errors"

type mockHasher struct{}

func (m *mockHasher) Hash(obj interface{}) (uint64, error) {
	h, ok := obj.(uint64)
	if ok == false {
		return 0, errors.New("What")
	}
	return h, nil
}

type mockEqualer struct{}

func (m *mockEqualer) Equals(obj1 interface{}, obj2 interface{}) (bool, error) {
	e1, ok := obj1.(uint64)
	if ok == false {
		return false, errors.New("What")
	}
	e2, ok := obj2.(uint64)
	if ok == false {
		return false, errors.New("What")
	}
	return e1 == e2, nil
}
