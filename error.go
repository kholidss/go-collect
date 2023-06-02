package collect

import (
	"fmt"
)

type Errors map[string]error

// Error returns the error string of Errors.
// func (es Errors) Error() string {
// 	if len(es) == 0 {
// 		return ""
// 	}

// 	keys := []string{}
// 	for key := range es {
// 		keys = append(keys, key)
// 	}
// 	sort.Strings(keys)

// 	s := ""
// 	for i, key := range keys {
// 		if i > 0 {
// 			s += "; "
// 		}
// 		if errs, ok := es[key].(Errors); ok {
// 			s += fmt.Sprintf("%v: (%v)", key, errs)
// 		} else {
// 			s += fmt.Sprintf("%v: %v", key, es[key].Error())
// 		}
// 	}
// 	return s + "."
// }

type errWrongDataType struct {
	process    string
	actualType string
}

func (w errWrongDataType) Error() string {
	return fmt.Sprintf("Cannot %s, expected something compatible with: %s", w.process, w.actualType)
}
