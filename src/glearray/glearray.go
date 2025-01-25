package glearray_P

import gleam_P "example.com/todo/gleam"

type Array_t[T gleam_P.Type[T]] struct {
	values []T
}

func (a Array_t[T]) Hash() uint32 {
	h := gleam_P.NewOrderedCollectionHasher()
	for _, v := range a.values {
		h.WriteHash(v.Hash())
	}
	return h.Sum()
}

func (a Array_t[T]) Equal(o Array_t[T]) bool {
	if len(a.values) != len(o.values) {
		return false
	}
	for i, v := range a.values {
		if !v.Equal(o.values[i]) {
			return false
		}
	}
	return true
}

func New[T gleam_P.Type[T]]() Array_t[T] {
	return Array_t[T]{}
}

func FromList[T gleam_P.Type[T]](list gleam_P.List_t[T]) Array_t[T] {
	res := New[T]()
	for {
		switch listC := list.(type) {
		case gleam_P.Empty_c[T]:
			return res
		case gleam_P.Nonempty_c[T]:
			{
				res.values = append(res.values, listC.P_0)
				list = listC.P_1
			}
		}
	}
}

func ToList[T gleam_P.Type[T]](array Array_t[T]) gleam_P.List_t[T] {
	return gleam_P.ToList(array.values...)
}

func Length[T gleam_P.Type[T]](array Array_t[T]) gleam_P.Int_t {
	return gleam_P.Int_t(len(array.values))
}

func doGet[T gleam_P.Type[T]](array Array_t[T], index gleam_P.Int_t) T {
	return array.values[index]
}

func doSet[T gleam_P.Type[T]](array Array_t[T], index gleam_P.Int_t, value T) Array_t[T] {
	res := make([]T, len(array.values))
	copy(res, array.values)
	res[index] = value
	return Array_t[T]{res}
}

func CopyPush[T gleam_P.Type[T]](array Array_t[T], value T) Array_t[T] {
	res := make([]T, len(array.values)+1)
	copy(res, array.values)
	res = append(res, value)
	return Array_t[T]{res}
}

func doInsert[T gleam_P.Type[T]](array Array_t[T], index gleam_P.Int_t, value T) Array_t[T] {
	res := make([]T, len(array.values)+1)
	copy(res, array.values[:index])
	res[index] = value
	copy(res[index+1:], array.values[index:])
	return Array_t[T]{res}
}
