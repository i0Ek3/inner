// package inner offers you guys a method to 
// get the structure inner of any type value x.
package inner

import (
    "fmt"
    "reflect"
)

const (
    conf = 0
    bg   = 0
    text = 31
)

// Inner shows x's inner structure
func Inner(val string, x interface{}) {
    fmt.Printf(Cyan(fmt.Sprintf("[\"%s\" => %T]\n", val, x)))
    inner(val, reflect.ValueOf(x))
}

// inner is the unexported implementation of the Inner call
func inner(val string, x reflect.Value) {
    switch x.Kind() {
    case reflect.Invalid:
        fmt.Printf("%s = invalid\n", val)
    case reflect.Slice, reflect.Array:
        for i := 0; i < x.Len(); i++ {
            inner(fmt.Sprintf("%s[%d]", val, i), x.Index(i))
        }
    case reflect.Struct:
        for i := 0; i < x.NumField(); i++ {
            fieldPath := fmt.Sprintf("%s.%s", val, x.Type().Field(i).Name)
            inner(fieldPath, x.Field(i))
        }
    case reflect.Map:
        for _, key := range x.MapKeys() {
            inner(fmt.Sprintf("%s[%s]", val, formatX(key)), x.MapIndex(key))
        }
    case reflect.Ptr:
        if x.IsNil() {
            fmt.Printf("%s = nil\n", val)
        } else {
            inner(fmt.Sprintf("(*%s)", val), x.Elem())
        }
    case reflect.Interface:
        if x.IsNil() {
            fmt.Printf("%s = nil\n", val)
        } else {
            fmt.Printf("%s.type = %s\n", val, x.Elem().Type())
            inner(val + ".value", x.Elem())
        }
    default:
        fmt.Printf("%s = %s\n", val, formatX(x))
    }
}
