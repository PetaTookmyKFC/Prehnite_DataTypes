package datatypes

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func enc_Float64(value float64, buff *bytes.Buffer) error {
	fmt.Println("Encoding : ", value)

	err := binary.Write(buff, binary.LittleEndian, Float64)
	if err != nil {
		return err
	}

	err = binary.Write(buff, binary.LittleEndian, value)
	if err != nil {
		return err
	}
	return nil
}
func dec_Float64(buff *bytes.Buffer) (float64, error) {
	var result float64
	err := binary.Read(buff, binary.LittleEndian, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func enc_Float32(value float32, buff *bytes.Buffer) error {
	fmt.Println("Encoding : ", value)

	err := binary.Write(buff, binary.LittleEndian, Float32)
	if err != nil {
		return err
	}

	err = binary.Write(buff, binary.LittleEndian, value)
	if err != nil {
		return err
	}
	return nil
}

func dec_Float32(buff *bytes.Buffer) (float32, error) {
	var result float32
	err := binary.Read(buff, binary.LittleEndian, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}
