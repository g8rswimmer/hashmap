package hashmap

type entry struct {
	key interface{}
	obj interface{}
}

type entries []entry

type hashMap struct {
	hash  Hasher
	equal Equaler
	table []entries
	size  uint64
}

func (h *hashMap) Get(k interface{}) (interface{}, error) {
	return nil, nil
}

func (h *hashMap) Put(k, v interface{}) error {
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
