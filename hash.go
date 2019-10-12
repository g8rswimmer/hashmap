package hashmap

import "errors"

type entry struct {
	key interface{}
	obj interface{}
}

type entries []entry

type HashMap struct {
	hash  Hasher
	equal Equaler
	table []entries
	size  uint64
}

func NewHashMap(size uint64, hash Hasher, equal Equaler) *HashMap {
	return &HashMap{
		hash:  hash,
		equal: equal,
		table: make([]entries, size),
		size:  size,
	}
}

func (h *HashMap) Get(k interface{}) (interface{}, error) {
	hash, err := h.hash.Hash(k)
	if err != nil {
		return nil, err
	}
	idx := int(hash % h.size)
	list := h.table[idx]
	switch len(list) {
	case 0:
		return nil, errors.New("hash map: unable to find entry")
	default:
		for _, en := range list {
			eq, err := h.equal.Equals(en.key, k)
			switch {
			case err != nil:
				return nil, err
			case eq:
				return en.obj, nil
			}
		}
	}
	return nil, errors.New("hash map: key was not found in list")
}

func (h *HashMap) Put(k, v interface{}) error {
	hash, err := h.hash.Hash(k)
	if err != nil {
		return err
	}
	e := entry{
		key: k,
		obj: v,
	}
	idx := int(hash % h.size)
	list := h.table[idx]
	switch len(list) {
	case 0:
		list = append(list, e)
	default:
		found := false
		for i, en := range list {
			eq, err := h.equal.Equals(en.key, k)
			switch {
			case err != nil:
				return err
			case eq:
				list[i] = e
				found = true
			}
			if found {
				break
			}
		}
		if found == false {
			list = append(list, e)
		}
	}
	h.table[idx] = list

	return nil
}
