package datatypes

import (
	"bytes"
	"encoding/binary"
	"errors"
	"reflect"
)

type DType uint8

const (
	Invalid DType = iota
	EOR           // End of recursion ( to signal the end of a repeating data structure ... used for maps)
	Bool
	ConvInt //  ( converts to int64 )
	Int8
	Int16
	Int32
	Int64
	Float32
	Float64
	String
	Array
	Map
	Struct
)

func (t DType) String() string {
	s := []string{"Invalid", "End of Recursion", "Bool", "ConvInt", "int8", "int16", "Int32", "Int64", "Float32", "Float64", "String", "Array", "Map", "Struct"}
	return s[t]
}
func GetType(value any) DType {

	// Check if the type is a struct...
	// This is because the type struct doesn't hande wildcards with .(type)
	if reflect.TypeOf(value).Kind() == reflect.Struct {
		return Struct
	}

	switch value.(type) {
	case bool:
		return Bool
	case int:
		return ConvInt
	case int8:
		return Int8
	case int16:
		return Int16
	case int32:
		return Int32
	case int64:
		return Int64
	case float64:
		return Float64
	case float32:
		return Float32
	case string:
		return String
	case []interface{}:
		return Array
	case map[string]interface{}:
		return Map
	case struct{}:
		return Struct
	}
	return Invalid
}

func Encode(pV any) ([]byte, error) {
	var err error

	buff := bytes.NewBuffer([]byte{})
	err = EncodeVersion(buff)
	if err != nil {
		return nil, err
	}
	t := GetType(pV)

	var value any
	if t == ConvInt {
		value = int64(pV.(int))
	} else {
		value = pV
	}

	err = _Encode(value, buff)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), err
}

func Decode(value []byte) (any, DType, error) {

	tVal := bytes.NewBuffer(value)

	// check the version matches
	valid, err := v_CheckCanDecode(tVal)
	if !valid {
		return nil, Invalid, err
	}
	return _Decode(tVal)
}

func DecodeStruct(value []byte, target interface{}) (DType, error) {
	tVal := bytes.NewBuffer(value)

	valid, err := v_CheckCanDecode(tVal)
	if !valid {
		return Invalid, err
	}

	// Check if the type is a struct
	var Dt DType
	err = binary.Read(tVal, binary.LittleEndian, &Dt)
	if err != nil {
		return Invalid, err
	}
	if Dt != Struct {
		return Invalid, errors.New("expected struct")
	}

	err = dec_Struct(tVal, target)
	if err != nil {
		return Invalid, err
	}
	return Struct, nil
}

func _Decode(buff *bytes.Buffer) (any, DType, error) {
	// A Copy but doesn't check for version and takes in a buffer instead of a byte array

	var Dt DType
	err := binary.Read(buff, binary.LittleEndian, &Dt)
	if err != nil {
		return nil, Invalid, err
	}

	var res any
	switch Dt {
	case Bool:
		res, err = dec_Bool(buff)
	case String:
		res, err = dec_String(buff)
	// Some of the different sizes of ints
	case Int8:
		res, err = dec_Int8(buff)
	case Int16:
		res, err = dec_Int16(buff)
	case Int32:
		res, err = dec_Int32(buff)
	case Int64:
		res, err = dec_Int64(buff)
	case Float32:
		res, err = dec_Float32(buff)
	case Float64:
		res, err = dec_Float64(buff)
	case Array:
		res, err = dec_Array(buff)
		// Unknown / unregistered
	case Map:
		res, err = dec_Map(buff)
	case EOR:
		res, err = nil, nil
	case Struct:
		err = errors.New("structs are not supported by this method. please use decodestruct")
	default:
		fallthrough
	case Invalid:
		{
			return nil, Invalid, err
		}
	}

	return res, Dt, err
}

func _Encode(data any, buff *bytes.Buffer) error {
	var err error

	t := GetType(data)

	var value any
	if t == ConvInt {
		value = int64(data.(int))
	} else {
		value = data
	}

	switch GetType(value) {
	case Bool:
		err = enc_Bool(value.(bool), buff)
	case Int8:
		err = enc_Int8(value.(int8), buff)
	case Int16:
		err = enc_Int16(value.(int16), buff)
	case Int32:
		err = enc_Int32(value.(int32), buff)
	case Int64, ConvInt:
		err = enc_Int64(value.(int64), buff)
	case Float32:
		err = enc_Float32(value.(float32), buff)
	case Float64:
		err = enc_Float64(value.(float64), buff)
	case String:
		err = enc_String(value.(string), buff)
	case Array:
		err = enc_Array(value.([]interface{}), buff)
	case Map:
		err = enc_Map(value.(map[string]interface{}), buff)
	case Struct:
		err = enc_Struct(value, buff)
	default:
		err = errors.New("unsupported type")
		return err
	}

	return err
}

func AreEqual(a, b any) bool {

	// Check if the types are the same

	if GetType(a) != GetType(b) {
		return false
	}

	t := GetType(a)

	switch t {
	case String, Bool, Array:
		{
			return reflect.DeepEqual(a, b)
		}

	case Int8, Int16, Int32, Int64, Float32, Float64, ConvInt:
		{
			// Try deep equal
			if reflect.DeepEqual(a, b) {
				return true
			}

			// Try casting
			var a1 interface{}
			var b1 interface{}

			a1 = a
			b1 = b

			if GetType(a1) == ConvInt {
				a1 = int64(a.(int))
			}
			if GetType(b1) == ConvInt {
				a1 = int64(b.(int))
			}

			return reflect.DeepEqual(a1, b1)

		}
	case Map:
		{
			return _comapreMaps(a.(map[string]any), b.(map[string]any))
		}
	case Struct:
		{
			if !reflect.DeepEqual(a, b) {
				return false
			}
			if a != b {
				return false
			}
			return true
		}
	}
	return false // Should Never get here
}

func _comapreMaps(a, b map[string]any) bool {
	// Check if the maps are the same length
	if len(a) != len(b) {
		return false
	}

	for key, value := range a {

		if !AreEqual(value, b[key]) {
			val, ok := b[key]
			if !ok {
				return false
			}
			if !reflect.DeepEqual(value, val) {
				if GetType(val) == Map {
					if !_comapreMaps(value.(map[string]any), val.(map[string]any)) {
						return false
					}
				} else {

					if GetType(val) == ConvInt {
						val = int64(val.(int))
					}
					if GetType(value) == ConvInt {
						value = int64(value.(int))
					}

					if !reflect.DeepEqual(value, val) {
						return false
					}
				}
			}
		}
	}
	return true
}
