package expr

import (
	"fmt"
	"reflect"

	"github.com/datasweet/cast"
)

func toNumber(n Node, val interface{}) float64 {
	v, ok := cast.AsFloat64(val)
	if ok {
		return v
	}
	panic(fmt.Sprintf("cannot convert %v (type %T) to type float64", n, val))
}

func isNumber(val interface{}) bool {
	return val != nil && reflect.TypeOf(val).Kind() == reflect.Float64
}

func canBeNumber(val interface{}) bool {
	if val != nil {
		return isNumberType(reflect.TypeOf(val))
	}
	return false
}

func equal(left, right interface{}) bool {
	if isNumber(left) && canBeNumber(right) {
		right, _ := cast.AsFloat64(right)
		return left == right
	} else if canBeNumber(left) && isNumber(right) {
		left, _ := cast.AsFloat64(left)
		return left == right
	} else {
		return reflect.DeepEqual(left, right)
	}
}

func extract(val interface{}, i interface{}) (interface{}, bool) {
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Array, reflect.Slice, reflect.String:
		n, ok := cast.AsInt64(i)
		if !ok {
			break
		}
		value := v.Index(int(n))
		if value.IsValid() && value.CanInterface() {
			return value.Interface(), true
		}
	case reflect.Map:
		value := v.MapIndex(reflect.ValueOf(i))
		if value.IsValid() && value.CanInterface() {
			return value.Interface(), true
		}
	case reflect.Struct:
		value := v.FieldByName(reflect.ValueOf(i).String())
		if value.IsValid() && value.CanInterface() {
			return value.Interface(), true
		}
	case reflect.Ptr:
		value := v.Elem()
		if value.IsValid() && value.CanInterface() {
			return extract(value.Interface(), i)
		}
	}
	return nil, false
}

func getFunc(val interface{}, name string) (interface{}, bool) {
	v := reflect.ValueOf(val)
	d := v
	if v.Kind() == reflect.Ptr {
		d = v.Elem()
	}

	switch d.Kind() {
	case reflect.Map:
		value := d.MapIndex(reflect.ValueOf(name))
		if value.IsValid() && value.CanInterface() {
			return value.Interface(), true
		}
		// A map may have method too.
		if v.NumMethod() > 0 {
			method := v.MethodByName(name)
			if method.IsValid() && method.CanInterface() {
				return method.Interface(), true
			}
		}
	case reflect.Struct:
		method := v.MethodByName(name)
		if method.IsValid() && method.CanInterface() {
			return method.Interface(), true
		}

		// If struct has not method, maybe it has func field.
		// To access this field we need dereferenced value.
		value := d.FieldByName(name)
		if value.IsValid() && value.CanInterface() {
			return value.Interface(), true
		}
	}
	return nil, false
}

func isNil(val interface{}) bool {
	v := reflect.ValueOf(val)
	return !v.IsValid() || v.IsNil()
}

func isNumberType(t reflect.Type) bool {
	t = dereference(t)
	if t != nil {
		switch t.Kind() {
		case reflect.Float32, reflect.Float64:
			fallthrough
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fallthrough
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return true
		}
	}
	return false
}

func dereference(ntype reflect.Type) reflect.Type {
	if ntype == nil {
		return nil
	}
	if ntype.Kind() == reflect.Ptr {
		ntype = dereference(ntype.Elem())
	}
	return ntype
}
