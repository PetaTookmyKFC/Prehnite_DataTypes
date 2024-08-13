package datatypes

import (
	"bytes"
	"encoding/binary"
	"errors"
)

func enc_Array(value []interface{}, buff *bytes.Buffer) error {

	var ArrayType DType

	// Set type to array
	err := binary.Write(buff, binary.LittleEndian, Array)
	if err != nil {
		return err
	}

	// Get the length of the array
	length := len(value)
	if length <= 0 {
		return errors.New("array is empty, can't store an empty array xx")
	}

	// Save the length of the array
	err = binary.Write(buff, binary.LittleEndian, uint32(length))
	if err != nil {
		return err
	}
	// Get the type of the first variable in the array
	if value := value[0]; value != nil {
		ArrayType = GetType(value)
	}

	if ArrayType == Invalid {
		return errors.New("array type is invalid")
	}

	// Write the type of the array
	err = binary.Write(buff, binary.LittleEndian, ArrayType)
	if err != nil {
		return err
	}

	// Loop through all of the items in the array
	for _, item := range value {

		if GetType(item) != ArrayType {
			err = errors.New("array can only have one type of item")
			return err
		}

		err = _Encode(item, buff)
		if err != nil {
			// err = prehnitelogs.Warn(err.Error() + fmt.Sprintf(" : %v", item))
			return err
		}
	}

	return nil
}

func dec_Array(buff *bytes.Buffer) ([]interface{}, error) {

	// Get the length of the array
	var length uint32
	err := binary.Read(buff, binary.LittleEndian, &length)
	if err != nil {
		return nil, err
	}

	// Get the type of the array
	var ArrayType DType
	err = binary.Read(buff, binary.LittleEndian, &ArrayType)
	if err != nil {
		return nil, err
	}

	if ArrayType == Invalid {
		return nil, errors.New("array type is invalid")
	}

	// Loop through all of the items in the array
	var result []interface{}
	for i := 0; i < int(length); i++ {
		value, DT, err := _Decode(buff)
		if err != nil {
			return nil, err
		}

		if DT != ArrayType {
			err = errors.New("array can only have one type of item")
			return nil, err
		}
		result = append(result, value)
	}

	return result, nil
}
