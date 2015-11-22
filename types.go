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

// String convert interface{} type to string if posible
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

// Int convert to int
func Int(s interface{}) (r int) {
	tmpr, _ := strconv.ParseInt(String(s), 0, 64)
	r = int(tmpr)
	return
}

// Int8 convert to int8
func Int8(s interface{}) (r int8) {
	tmpr, _ := strconv.ParseInt(String(s), 0, 64)
	r = int8(tmpr)
	return
}

// Int16 convert to int16
func Int16(s interface{}) (r int16) {
	tmpr, _ := strconv.ParseInt(String(s), 0, 64)
	r = int16(tmpr)
	return
}

// Int32 convert to int32
func Int32(s interface{}) (r int32) {
	tmpr, _ := strconv.ParseInt(String(s), 0, 64)
	r = int32(tmpr)
	return
}

// Rune convert to rune (int32)
func Rune(s interface{}) (r rune) {
	r = Int32(s)
	return
}

// Int64 convert to int64
func Int64(s interface{}) (r int64) {
	r, _ = strconv.ParseInt(String(s), 0, 64)
	return
}

// Byte convert to int8
func Byte(s interface{}) (r byte) {
	tmpr, _ := strconv.ParseUint(String(s), 0, 64)
	r = byte(tmpr)
	return
}

// Uint8 convert to uint8
func Uint8(s interface{}) (r uint8) {
	r = Byte(s)
	return
}

// Uint16 convert to uint16
func Uint16(s interface{}) (r uint16) {
	tmpr, _ := strconv.ParseUint(String(s), 0, 64)
	r = uint16(tmpr)
	return
}

// Uint32 convert to uint32
func Uint32(s interface{}) (r uint32) {
	tmpr, _ := strconv.ParseUint(String(s), 0, 64)
	r = uint32(tmpr)
	return
}

// Uint64 convert to uint64
func Uint64(s interface{}) (r uint64) {
	r, _ = strconv.ParseUint(String(s), 0, 64)
	return
}

// Float32 convert to float32
func Float32(s interface{}) (r float32) {
	tmpr, _ := strconv.ParseFloat(String(s), 64)
	r = float32(tmpr)
	return
}

// Float64 convert to float64
func Float64(s interface{}) (r float64) {
	r, _ = strconv.ParseFloat(String(s), 64)
	return
}

// IsNumber check type conformity
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

// TypeName return name of type
func TypeName(s interface{}) string {
	return fmt.Sprintf("%v", reflect.TypeOf(s))
}
