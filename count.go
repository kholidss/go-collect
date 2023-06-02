package collect

import (
	"reflect"
)

// The Count function counts the number of elements in the input value. It supports different data types,
// including string, integer types (int, int8, int16, int32, int64), struct, map, slice, and array.
func Count[T any](v T) (int, error) {
	switch val := any(v).(type) {
	case string:
		return countString(val), nil
	case int:
		return countInt(val), nil
	case int8:
		return countInt(int(val)), nil
	case int16:
		return countInt(int(val)), nil
	case int32:
		return countInt(int(val)), nil
	case int64:
		return countInt(int(val)), nil
	default:
		t := reflect.TypeOf(val)
		tv := reflect.ValueOf(val)
		if t.Kind() == reflect.Struct {
			return t.NumField(), nil
		}
		if t.Kind() == reflect.Map {
			return tv.Len(), nil
		}
		if t.Kind() == reflect.Slice {
			return tv.Len(), nil
		}
		if t.Kind() == reflect.Array {
			return tv.Len(), nil
		}
		return 0, errWrongDataType{
			process:    "count",
			actualType: typeof(val),
		}
	}
}

// The CountBy function counts the occurrences of each element in the input slice and returns a map with the elements as keys and their counts as values.
// The input slice can contain elements of any type.
func CountBy[V any](inp []V) map[interface{}]int {
	counts := make(map[interface{}]int)
	for _, item := range inp {
		counts[item]++
	}
	return counts
}
