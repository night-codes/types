package types

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

type (
	struct1 struct {
		A string
		B int
		C float64
	}
	structWithString struct {
		A string
		B int
		C float64
	}
)

func (s structWithString) String() string {
	return fmt.Sprintf("{A:%s, B:%d, C:%s}", s.A, s.B, strconv.FormatFloat(s.C, 'f', -1, 64))
}

func BenchmarkTestString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		String(21312.2134567)
	}
}

func BenchmarkTestFloat64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Float64("21312.2134567")
	}
}

func BenchmarkTestStr2Int(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Int("2134567")
	}
}

func BenchmarkTestFloat2Int(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Int(2134567.123)
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name string
		arg  interface{}
		want string
	}{
		{"[]byte", []byte{102, 111, 111, 44, 32, 98, 97, 114}, "foo, bar"},
		{"[]byte_emty", []byte{}, ""},
		{"string", "foo, bar", "foo, bar"},
		{"string empty", "", ""},
		{"int", int(123456), "123456"},
		{"-int", -123456, "-123456"},
		{"float64", 123.456, "123.456"},
		{"-float64", -123.456, "-123.456"},
		{"float32", float32(123.456), "123.456"},
		{"-float32", float32(-123.456), "-123.456"},
		{"bool_true", true, "true"},
		{"bool_false", false, "false"},
		{"[]string", []string{"one", "2", "tree"}, "one2tree"},
		{"struct", struct1{"test", 123, 123.45}, "{A:test B:123 C:123.45}"},
		{"struct_with_string()", structWithString{"test", 123, 123.45}, "{A:test, B:123, C:123.45}"},
		{"map[string]interface{}", map[string]interface{}{"test": true}, "map[test:true]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.arg); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	tests := []struct {
		name string
		arg  interface{}
		want int
	}{
		{"true", true, 0},
		{"123.12", 123.12, 123},
		{"string_123", "123", 123},
		{"-string_123", "-123", -123},
		{"string_without_num", "abcdef", 0},
		{"string_0", "0", 0},
		{"string_-0", "-0", 0},
		{"string_float", "0.123", 0},
		{"-string_123.45", "-123.45", 0},
		{"-string_with_space", " 123 ", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.arg); got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool(t *testing.T) {
	tests := []struct {
		name string
		arg  interface{}
		want bool
	}{
		{"string_true", "true", true},
		{"string_1", "1", true},
		{"0", 0, false}, // only 0 == false
		{"1", 1, true},
		{"-1", -1, true},
		{"100", 100, true},
		{"100.0", 100.0, true},
		{"string_0", "0", false},
		{"string_1", "11", false},
		{"string_with_space", " true", false},
		{"string empty", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bool(tt.arg); got != tt.want {
				t.Errorf("Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8(t *testing.T) {
	tests := []struct {
		name string
		arg  interface{}
		want int8
	}{
		{"123.12", 123.12, 123},
		{"-123.12", -123.12, -123},
		{"string_123", "123", 123},
		{"-string_123", "-123", -123},
		{"string_without_num", "abcdef", 0},
		{"string_0", "0", 0},
		{"string_-0", "-0", 0},
		{"string_float", "0.123", 0},
		{"-string_123.45", "-123.45", 0},
		{"-string_with_space", " 123 ", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8(tt.arg); got != tt.want {
				t.Errorf("Int8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16(t *testing.T) {
	tests := []struct {
		name string
		arg  interface{}
		want int16
	}{
		{"123.12", 123.12, 123},
		{"-123.12", -123.12, -123},
		{"string_123", "123", 123},
		{"-string_123", "-123", -123},
		{"string_without_num", "abcdef", 0},
		{"string_0", "0", 0},
		{"string_-0", "-0", 0},
		{"string_float", "0.123", 0},
		{"-string_123.45", "-123.45", 0},
		{"-string_with_space", " 123 ", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16(tt.arg); got != tt.want {
				t.Errorf("Int16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	tests := []struct {
		name string
		arg  interface{}
		want int32
	}{
		{"123.12", 123.12, 123},
		{"-123.12", -123.12, -123},
		{"string_123", "123", 123},
		{"-string_123", "-123", -123},
		{"string_without_num", "abcdef", 0},
		{"string_0", "0", 0},
		{"string_-0", "-0", 0},
		{"string_float", "0.123", 0},
		{"-string_123.45", "-123.45", 0},
		{"-string_with_space", " 123 ", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32(tt.arg); got != tt.want {
				t.Errorf("Int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRune(t *testing.T) {
	TestInt32(t)
}

func TestInt64(t *testing.T) {
	tests := []struct {
		name string
		arg  interface{}
		want int64
	}{
		{"123.12", 123.12, 123},
		{"-123.12", -123.12, -123},
		{"string_123", "123", 123},
		{"-string_123", "-123", -123},
		{"string_without_num", "abcdef", 0},
		{"string_0", "0", 0},
		{"string_-0", "-0", 0},
		{"string_float", "0.123", 0},
		{"-string_123.45", "-123.45", 0},
		{"-string_with_space", " 123 ", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64(tt.arg); got != tt.want {
				t.Errorf("Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByte(t *testing.T) {
	TestInt8(t)
}

func TestUint(t *testing.T) {
	tests := []struct {
		name string
		arg  interface{}
		want uint
	}{
		{"123.12", 123.12, 123},
		{"string_123", "12356", 12356},
		{"-string_123", "-123", 0},
		{"string_without_num", "abcdef", 0},
		{"string_0", "0", 0},
		{"string_-0", "-0", 0},
		{"string_float", "0.123", 0},
		{"-string_123.45", "-123.45", 0},
		{"-string_with_space", " 123 ", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint(tt.arg); got != tt.want {
				t.Errorf("Uint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8(t *testing.T) {
	tests := []struct {
		name string
		arg  interface{}
		want uint8
	}{
		{"123.12", 123.12, 123},
		{"-123.12", -103.12, 153}, // overflow
		{"string_123", "123", 123},
		{"string_12356", "12356", 255}, // overflow
		{"-string_123", "-123", 0},
		{"string_without_num", "abcdef", 0},
		{"string_0", "0", 0},
		{"string_-0", "-0", 0},
		{"string_float", "0.123", 0},
		{"-string_123.45", "-123.45", 0},
		{"-string_with_space", " 123 ", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint8(tt.arg); got != tt.want {
				t.Errorf("Uint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16(t *testing.T) {
	tests := []struct {
		name string
		arg  interface{}
		want uint16
	}{
		{"123.12", 123.12, 123},
		{"-123.12", -103.12, 65433}, // overflow
		{"string_123", "123", 123},
		{"string_123567", "123567", 65535}, // overflow
		{"-string_123", "-123", 0},
		{"string_without_num", "abcdef", 0},
		{"string_0", "0", 0},
		{"string_-0", "-0", 0},
		{"string_float", "0.123", 0},
		{"-string_123.45", "-123.45", 0},
		{"-string_with_space", " 123 ", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16(tt.arg); got != tt.want {
				t.Errorf("Uint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32(t *testing.T) {
	tests := []struct {
		name string
		arg  interface{}
		want uint32
	}{
		{"123.12", 123.12, 123},
		{"-123.12", -103.12, 4294967193}, // overflow
		{"string_123", "123", 123},
		{"string_12345678901", "12345678901", 4294967295}, // overflow
		{"-string_123", "-123", 0},
		{"string_without_num", "abcdef", 0},
		{"string_0", "0", 0},
		{"string_-0", "-0", 0},
		{"string_float", "0.123", 0},
		{"-string_123.45", "-123.45", 0},
		{"-string_with_space", " 123 ", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32(tt.arg); got != tt.want {
				t.Errorf("Uint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	tests := []struct {
		name string
		arg  interface{}
		want uint64
	}{
		{"123.12", 123.12, 123},
		{"-123.12", -103.12, 18446744073709551513}, // overflow
		{"string_123", "123", 123},
		{"string_12345678901", "12345678901", 12345678901},
		{"-string_123", "-123", 0},
		{"string_without_num", "abcdef", 0},
		{"string_0", "0", 0},
		{"string_-0", "-0", 0},
		{"string_float", "0.123", 0},
		{"-string_123.45", "-123.45", 0},
		{"-string_with_space", " 123 ", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64(tt.arg); got != tt.want {
				t.Errorf("Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	tests := []struct {
		name string
		arg  interface{}
		want float32
	}{
		{"123", 123, 123.0},
		{"true", true, 0},
		{"string_123", "123", 123.0},
		{"-string_123", "-123", -123.0},
		{"string_12345678901", "12345678901", 12345678901.0},
		{"string_without_num", "abcdef", 0},
		{"string_0", "0", 0},
		{"string_-0", "-0", 0},
		{"string_float", "0.123", 0.123},
		{"-string_123.45", "-123.45", -123.45},
		{"-string_123.45", "-123,45", -123.45},
		{"-string_123.45", "-123.123456789", -123.12346},
		{"-string_with_space", " 123 ", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32(tt.arg); got != tt.want {
				t.Errorf("Float32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	tests := []struct {
		name string
		arg  interface{}
		want float64
	}{
		{"string_123", "123", 123.0},
		{"-string_123", "-123", -123.0},
		{"string_without_num", "abcdef", 0},
		{"string_0", "0", 0},
		{"string_-0", "-0", 0},
		{"string_float", "0.123", 0.123},
		{"-string_123.45", "-123.45", -123.45},
		{"-string_123.45", "-123,45", -123.45},
		{"-string_123.45", "123.1234567891234567", 123.1234567891234567},
		{"-string_123.45", "-123.12345678912345", -123.12345678912345},
		{"-string_with_space", " 123 ", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64(tt.arg); got != tt.want {
				t.Errorf("Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args struct {
		iface []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantRet map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := Map(tt.args.iface...); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("Map() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func TestSlice(t *testing.T) {
	type args struct {
		iface []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantRet []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := Slice(tt.args.iface...); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("Slice() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func TestMGet(t *testing.T) {
	type args struct {
		obj  interface{}
		path []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MGet(tt.args.obj, tt.args.path...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPtr(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPtr(tt.args.s); got != tt.want {
				t.Errorf("IsPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNumber(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumber(tt.args.s); got != tt.want {
				t.Errorf("IsNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTypeName(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TypeName(tt.args.v); got != tt.want {
				t.Errorf("TypeName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKind(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Kind(tt.args.v); got != tt.want {
				t.Errorf("Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpected(t *testing.T) {
	type args struct {
		v     interface{}
		types []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Expected(tt.args.v, tt.args.types); got != tt.want {
				t.Errorf("Expected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetField(t *testing.T) {
	type args struct {
		structPtr interface{}
		name      string
		value     interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetField(tt.args.structPtr, tt.args.name, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("SetField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChangeStruct(t *testing.T) {
	type args struct {
		structPtr  interface{}
		changesMap map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ChangeStruct(tt.args.structPtr, tt.args.changesMap); (err != nil) != tt.wantErr {
				t.Errorf("ChangeStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChangeStructMust(t *testing.T) {
	type args struct {
		structPtr  interface{}
		changesMap map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ChangeStructMust(tt.args.structPtr, tt.args.changesMap); (err != nil) != tt.wantErr {
				t.Errorf("ChangeStructMust() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBytes2String(t *testing.T) {
	type args struct {
		bs []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bytes2String(tt.args.bs); got != tt.want {
				t.Errorf("Bytes2String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString2Bytes(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String2Bytes(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("String2Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
