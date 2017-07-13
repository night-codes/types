package types

import (
	"errors"
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
func String(s interface{}) string {
	if s == nil {
		return ""
	}
	switch v := s.(type) {
	case string:
		return v
	case *string:
		return *v
	case []byte:
		return string(v)
	case *[]byte:
		return string(*v)
	case []string:
		return strings.Join(v, "")
	case *[]string:
		return strings.Join(*v, "")
	}

	return fmt.Sprintf("%v", s)
}

// Int converting
func Int(s interface{}) int {
	switch v := s.(type) {
	case int:
		return int(v)
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	}

	tmpr, _ := strconv.Atoi(String(s))
	return tmpr
}

// Bool converting
func Bool(s interface{}) bool {
	switch v := s.(type) {
	case bool:
		return v
	case int:
		return v != 0
	case int8:
		return v != 0
	case int16:
		return v != 0
	case int32:
		return v != 0
	case int64:
		return v != 0
	case uint:
		return v != 0
	case uint8:
		return v != 0
	case uint16:
		return v != 0
	case uint32:
		return v != 0
	case uint64:
		return v != 0
	case float32:
		return v != 0
	case float64:
		return v != 0
	}

	r, _ := strconv.ParseBool(String(s))
	return r
}

// Int8 converting
func Int8(s interface{}) int8 {
	switch v := s.(type) {
	case int:
		return int8(v)
	case int8:
		return int8(v)
	case int16:
		return int8(v)
	case int32:
		return int8(v)
	case int64:
		return int8(v)
	case uint:
		return int8(v)
	case uint8:
		return int8(v)
	case uint16:
		return int8(v)
	case uint32:
		return int8(v)
	case uint64:
		return int8(v)
	case float32:
		return int8(v)
	case float64:
		return int8(v)
	}

	tmpr, _ := strconv.ParseInt(String(s), 10, 8)
	return int8(tmpr)
}

// Int16 converting
func Int16(s interface{}) int16 {
	switch v := s.(type) {
	case int:
		return int16(v)
	case int8:
		return int16(v)
	case int16:
		return int16(v)
	case int32:
		return int16(v)
	case int64:
		return int16(v)
	case uint:
		return int16(v)
	case uint8:
		return int16(v)
	case uint16:
		return int16(v)
	case uint32:
		return int16(v)
	case uint64:
		return int16(v)
	case float32:
		return int16(v)
	case float64:
		return int16(v)
	}

	tmpr, _ := strconv.ParseInt(String(s), 10, 16)
	return int16(tmpr)

}

// Int32 converting
func Int32(s interface{}) int32 {
	switch v := s.(type) {
	case int:
		return int32(v)
	case int8:
		return int32(v)
	case int16:
		return int32(v)
	case int32:
		return int32(v)
	case int64:
		return int32(v)
	case uint:
		return int32(v)
	case uint8:
		return int32(v)
	case uint16:
		return int32(v)
	case uint32:
		return int32(v)
	case uint64:
		return int32(v)
	case float32:
		return int32(v)
	case float64:
		return int32(v)
	}

	tmpr, _ := strconv.ParseInt(String(s), 10, 32)
	return int32(tmpr)
}

// Rune converting (int32)
func Rune(s interface{}) (r rune) {
	r = Int32(s)
	return
}

// Int64 converting
func Int64(s interface{}) int64 {
	switch v := s.(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return int64(v)
	case uint:
		return int64(v)
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		return int64(v)
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	}

	r, _ := strconv.ParseInt(String(s), 10, 64)
	return r
}

// Byte converting
func Byte(s interface{}) byte {
	return byte(Uint8(s))
}

// Uint converting
func Uint(s interface{}) uint {
	switch v := s.(type) {
	case int:
		return uint(v)
	case int8:
		return uint(v)
	case int16:
		return uint(v)
	case int32:
		return uint(v)
	case int64:
		return uint(v)
	case uint:
		return uint(v)
	case uint8:
		return uint(v)
	case uint16:
		return uint(v)
	case uint32:
		return uint(v)
	case uint64:
		return uint(v)
	case float32:
		return uint(v)
	case float64:
		return uint(v)
	}

	tmpr, _ := strconv.ParseUint(String(s), 10, 64)
	return uint(tmpr)
}

// Uint8 converting
func Uint8(s interface{}) uint8 {
	switch v := s.(type) {
	case int:
		return uint8(v)
	case int8:
		return uint8(v)
	case int16:
		return uint8(v)
	case int32:
		return uint8(v)
	case int64:
		return uint8(v)
	case uint:
		return uint8(v)
	case uint8:
		return uint8(v)
	case uint16:
		return uint8(v)
	case uint32:
		return uint8(v)
	case uint64:
		return uint8(v)
	case float32:
		return uint8(v)
	case float64:
		return uint8(v)
	}

	tmpr, _ := strconv.ParseUint(String(s), 10, 8)
	return uint8(tmpr)
}

// Uint16 converting
func Uint16(s interface{}) uint16 {
	switch v := s.(type) {
	case int:
		return uint16(v)
	case int8:
		return uint16(v)
	case int16:
		return uint16(v)
	case int32:
		return uint16(v)
	case int64:
		return uint16(v)
	case uint:
		return uint16(v)
	case uint8:
		return uint16(v)
	case uint16:
		return uint16(v)
	case uint32:
		return uint16(v)
	case uint64:
		return uint16(v)
	case float32:
		return uint16(v)
	case float64:
		return uint16(v)
	}

	tmpr, _ := strconv.ParseUint(String(s), 10, 16)
	return uint16(tmpr)
}

// Uint32 converting
func Uint32(s interface{}) uint32 {
	switch v := s.(type) {
	case int:
		return uint32(v)
	case int8:
		return uint32(v)
	case int16:
		return uint32(v)
	case int32:
		return uint32(v)
	case int64:
		return uint32(v)
	case uint:
		return uint32(v)
	case uint8:
		return uint32(v)
	case uint16:
		return uint32(v)
	case uint32:
		return uint32(v)
	case uint64:
		return uint32(v)
	case float32:
		return uint32(v)
	case float64:
		return uint32(v)
	}

	tmpr, _ := strconv.ParseUint(String(s), 10, 32)
	return uint32(tmpr)
}

// Uint64 converting
func Uint64(s interface{}) (r uint64) {
	switch v := s.(type) {
	case int:
		return uint64(v)
	case int8:
		return uint64(v)
	case int16:
		return uint64(v)
	case int32:
		return uint64(v)
	case int64:
		return uint64(v)
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return uint64(v)
	case float32:
		return uint64(v)
	case float64:
		return uint64(v)
	}

	r, _ = strconv.ParseUint(String(s), 10, 64)
	return
}

// Float32 converting
func Float32(s interface{}) float32 {
	switch v := s.(type) {
	case int:
		return float32(v)
	case int8:
		return float32(v)
	case int16:
		return float32(v)
	case int32:
		return float32(v)
	case int64:
		return float32(v)
	case uint:
		return float32(v)
	case uint8:
		return float32(v)
	case uint16:
		return float32(v)
	case uint32:
		return float32(v)
	case uint64:
		return float32(v)
	case float32:
		return float32(v)
	case float64:
		return float32(v)
	}

	tmpr, _ := strconv.ParseFloat(String(s), 32)
	return float32(tmpr)
}

// Float64 converting
func Float64(s interface{}) float64 {
	switch v := s.(type) {
	case int:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint:
		return float64(v)
	case uint8:
		return float64(v)
	case uint16:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case float32:
		return float64(v)
	case float64:
		return float64(v)
	}

	r, _ := strconv.ParseFloat(String(s), 64)
	return r
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
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	}

	return false
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

// SetField applies value to struct field
func SetField(structPtr interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(structPtr).Elem()
	fieldValue := structValue.FieldByName(name)

	if !fieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in structPtr", name)
	}

	if !fieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	fieldType := fieldValue.Type()
	val := reflect.ValueOf(value)
	if fieldType != val.Type() && fieldType.Kind() != reflect.Interface {
		return errors.New("Provided value type didn't match structPtr field type")
	}

	fieldValue.Set(val)
	return nil
}

// ChangeStruct applies map of changes to struct
func ChangeStruct(structPtr interface{}, changesMap map[string]interface{}) {
	for k, v := range changesMap {
		SetField(structPtr, k, v)
	}
}
