package expr

import (
	"reflect"
)

func contains(needle interface{}, array interface{}) interface{} {
	arr, ok := asArray(array)
	if !ok {
		return false
	}

	if an, ok := asArray(needle); ok {
		cnt := len(an)
		out := make([]interface{}, cnt)
		for i := 0; i < cnt; i++ {
			out[i] = containsImpl(an[i], arr)
		}
		return out
	}

	return containsImpl(needle, arr)
}

func notcontains(needle interface{}, array interface{}) interface{} {
	arr, ok := asArray(array)
	if !ok {
		return true
	}

	if an, ok := asArray(needle); ok {
		cnt := len(an)
		out := make([]interface{}, cnt)
		for i := 0; i < cnt; i++ {
			out[i] = !containsImpl(an[i], arr)
		}
		return out
	}

	return !containsImpl(needle, arr)
}

func containsImpl(needle interface{}, arr []interface{}) bool {
	for _, v := range arr {
		if reflect.DeepEqual(needle, v) {
			return true
		}
	}
	return false
}
