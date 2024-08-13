package datatypes

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

/*
	TODO :
		// * Set version
		* Set type as structure
		* Create a empty map[string]interface to hold all the structure data and keys
		* Loop through all the keys and values of the structure and append them to the map
		* Convert the map to byte array


		-> Func ( *PointerToEmptyStruct )
		<-> Decode ( *Pointer )


*/

func enc_Struct(value interface{}, buff *bytes.Buffer) error {
	// Write the type for struct
	err := binary.Write(buff, binary.LittleEndian, Struct)
	if err != nil {
		return err
	}

	// Convert the struct to map
	ma, err := StructToMap(value)
	if err != nil {
		return err
	}

	err = _Encode(ma, buff)
	if err != nil {
		return err
	}
	return nil
}

func dec_Struct(buff *bytes.Buffer, target interface{}) error {
	// Decode the map
	ma, ty, err := _Decode(buff)
	if err != nil {
		return err
	}

	if ty != Map {
		return fmt.Errorf("expected map, got %s", ty)
	}
	// Place the map into structure
	err = MapToStruct(ma.(map[string]interface{}), target)
	if err != nil {
		return err
	}

	return nil
}

func StructToMap(str interface{}) (map[string]interface{}, error) {

	// Make response map
	var result map[string]interface{} = make(map[string]interface{})

	// Save the type of the struct
	var Type = reflect.TypeOf(str)
	var value = reflect.ValueOf(str)
	// Loop through the fields
	for i := 0; i < Type.NumField(); i++ {
		field := Type.Field(i)
		result[field.Name] = value.Field(i).Interface()
	}

	return result, nil
}
func MapToStruct(m map[string]interface{}, target interface{}) error {
	v := reflect.ValueOf(target)

	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("target must be a pointer to a struct")
	}

	t := v.Elem().Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if value, ok := m[field.Name]; ok {

			if reflect.TypeOf(value) != field.Type {
				return fmt.Errorf("type mismatch for field %s", field.Name)
			}

			v.Elem().Field(i).Set(reflect.ValueOf(value))
		}
	}
	return nil
}
