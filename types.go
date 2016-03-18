package types

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type (
	// Obj object type
	Obj map[string]interface{}

	// Arr array type
	Arr []interface{}

	// Mixed mixed type
	Mixed interface{}
)

// String convert interface{} type to string if it posible
func String(s interface{}) (r string) {
	if s == nil {
		r = ""
		return
	}
	switch v := s.(type) {
	case string:
		r = v
	case []byte:
		r = string(v)
	case []string:
		r = strings.Join(v, "")
	default:
		r = fmt.Sprintf("%v", v)
	}
	return
}

// Int converting
func Int(s interface{}) (r int) {
	tmpr, _ := strconv.ParseInt(String(s), 0, 64)
	r = int(tmpr)
	return
}

// Bool converting
func Bool(s interface{}) (r bool) {
	r, _ = strconv.ParseBool(String(s))
	return
}

// Int8 converting
func Int8(s interface{}) (r int8) {
	tmpr, _ := strconv.ParseInt(String(s), 0, 64)
	r = int8(tmpr)
	return
}

// Int16 converting
func Int16(s interface{}) (r int16) {
	tmpr, _ := strconv.ParseInt(String(s), 0, 64)
	r = int16(tmpr)
	return
}

// Int32 converting
func Int32(s interface{}) (r int32) {
	tmpr, _ := strconv.ParseInt(String(s), 0, 64)
	r = int32(tmpr)
	return
}

// Rune converting to rune (int32)
func Rune(s interface{}) (r rune) {
	r = Int32(s)
	return
}

// Int64 converting
func Int64(s interface{}) (r int64) {
	r, _ = strconv.ParseInt(String(s), 0, 64)
	return
}

// Byte converting
func Byte(s interface{}) (r byte) {
	tmpr, _ := strconv.ParseUint(String(s), 0, 64)
	r = byte(tmpr)
	return
}

// Uint converting
func Uint(s interface{}) (r uint) {
	tmpr, _ := strconv.ParseUint(String(s), 0, 64)
	r = uint(tmpr)
	return
}

// Uint8 converting
func Uint8(s interface{}) (r uint8) {
	r = Byte(s)
	return
}

// Uint16 converting
func Uint16(s interface{}) (r uint16) {
	tmpr, _ := strconv.ParseUint(String(s), 0, 64)
	r = uint16(tmpr)
	return
}

// Uint32 converting
func Uint32(s interface{}) (r uint32) {
	tmpr, _ := strconv.ParseUint(String(s), 0, 64)
	r = uint32(tmpr)
	return
}

// Uint64 converting
func Uint64(s interface{}) (r uint64) {
	r, _ = strconv.ParseUint(String(s), 0, 64)
	return
}

// Float32 converting
func Float32(s interface{}) (r float32) {
	tmpr, _ := strconv.ParseFloat(String(s), 64)
	r = float32(tmpr)
	return
}

// Float64 converting
func Float64(s interface{}) (r float64) {
	r, _ = strconv.ParseFloat(String(s), 64)
	return
}

// Map convert to map[string]interface{}
func Map(iface ...interface{}) (ret map[string]interface{}) {
	ret = map[string]interface{}{}
	if len(iface) == 0 || iface[0] == nil {
		return
	}
	t := iface[0]
	val := reflect.ValueOf(t)
	if IsPtr(t) {
		val = val.Elem()
	}
	if val.Kind().String() != "map" {
		return
	}
	val = val.Convert(reflect.TypeOf(map[string]interface{}{}))
	i := val.Interface()
	ret = i.(map[string]interface{})
	return
}

// Slice convert to []interface{}
func Slice(iface ...interface{}) (ret []interface{}) {
	ret = []interface{}{}
	if len(iface) == 0 || iface[0] == nil {
		return
	}
	t := iface[0]
	val := reflect.ValueOf(t)
	if IsPtr(t) {
		val = val.Elem()
	}
	if val.Kind() != reflect.Slice || val.Len() == 0 {
		return
	}

	ret = make([]interface{}, val.Len())
	for i := 0; i < val.Len(); i++ {
		ret[i] = val.Index(i).Interface()
	}
	return
}

// MGet getting data from deep interface{} map by path
func MGet(obj interface{}, path ...interface{}) interface{} {
	if len(path) != 0 {
		if m, ok := Map(obj)[String(path[0])]; ok {
			if len(path) == 1 {
				return m
			}
			return MGet(m, path[1:]...)
		}
		return nil
	}
	return obj
}

// IsPtr detect pointer type
func IsPtr(s interface{}) bool {
	return reflect.TypeOf(s).String()[0] == '*'
}

// IsNumber check number type conformity
func IsNumber(s interface{}) bool {
	switch s.(type) {
	case int:
		return true
	case int8:
		return true
	case int16:
		return true
	case int32:
		return true
	case int64:
		return true
	case uint:
		return true
	case uint8:
		return true
	case uint16:
		return true
	case uint32:
		return true
	case uint64:
		return true
	case float32:
		return true
	case float64:
		return true
	default:
		return false
	}
}

// TypeName return name of variable type
func TypeName(v interface{}) string {
	return fmt.Sprintf("%v", reflect.TypeOf(v))
}

// Kind return variable kind
func Kind(v interface{}) string {
	val := reflect.ValueOf(v)
	if IsPtr(v) {
		val = val.Elem()
	}
	return val.Kind().String()
}

// Expected return conformity variable type to expected types
func Expected(v interface{}, types []string) bool {
	tp := TypeName(v)
	for _, value := range types {
		if tp == value {
			return true
		}
	}
	return false
}
