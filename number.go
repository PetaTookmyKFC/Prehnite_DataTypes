package datatypes

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func enc_Int8(value int8, buff *bytes.Buffer) error {

	err := binary.Write(buff, binary.LittleEndian, Int8)
	if err != nil {
		return err
	}

	err = binary.Write(buff, binary.LittleEndian, value)
	if err != nil {
		return err
	}
	return nil
}
func dec_Int8(buff *bytes.Buffer) (int8, error) {
	var result int8
	err := binary.Read(buff, binary.LittleEndian, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func enc_Int16(value int16, buff *bytes.Buffer) error {
	err := binary.Write(buff, binary.LittleEndian, Int16)
	if err != nil {
		return err
	}

	err = binary.Write(buff, binary.LittleEndian, value)
	if err != nil {
		return err
	}
	return nil
}
func dec_Int16(buff *bytes.Buffer) (int16, error) {
	var result int16
	err := binary.Read(buff, binary.LittleEndian, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}
func enc_Int32(value int32, buff *bytes.Buffer) error {
	fmt.Println("Encoding : ", value)

	err := binary.Write(buff, binary.LittleEndian, Int32)
	if err != nil {
		return err
	}

	err = binary.Write(buff, binary.LittleEndian, value)
	if err != nil {
		return err
	}

	return nil
}

func dec_Int32(buff *bytes.Buffer) (int32, error) {
	var result int32
	err := binary.Read(buff, binary.LittleEndian, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func enc_Int64(value int64, buff *bytes.Buffer) error {
	fmt.Println("Encoding : ", value)

	err := binary.Write(buff, binary.LittleEndian, Int64)
	if err != nil {
		return err
	}

	err = binary.Write(buff, binary.LittleEndian, value)
	if err != nil {
		return err
	}
	return nil
}
func dec_Int64(buff *bytes.Buffer) (int64, error) {
	var result int64
	err := binary.Read(buff, binary.LittleEndian, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}
