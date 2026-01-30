package datatypes

import (
	"bytes"
	"encoding/binary"
)

func enc_Uint8(value uint8, buff *bytes.Buffer) error {
	err := binary.Write(buff, binary.LittleEndian, Uint8)
	if err != nil {
		return err
	}

	err = binary.Write(buff, binary.LittleEndian, value)
	if err != nil {
		return err
	}
	return nil
}
func dec_Uint8(buff *bytes.Buffer) (uint8, error) {
	var result uint8
	err := binary.Read(buff, binary.LittleEndian, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}
func enc_Uint16(value uint16, buff *bytes.Buffer) error {
	err := binary.Write(buff, binary.LittleEndian, Uint16)
	if err != nil {
		return err
	}
	err = binary.Write(buff, binary.LittleEndian, value)
	if err != nil {
		return err
	}
	return nil
}
func dec_Uint16(buff *bytes.Buffer) (uint16, error) {
	var result uint16
	err := binary.Read(buff, binary.LittleEndian, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func enc_Uint32(value uint32, buff *bytes.Buffer) error {
	err := binary.Write(buff, binary.LittleEndian, Uint32)
	if err != nil {
		return err
	}
	err = binary.Write(buff, binary.LittleEndian, value)
	if err != nil {
		return err
	}
	return nil
}
func dec_Uint32(buff *bytes.Buffer) (uint32, error) {
	var result uint32
	err := binary.Read(buff, binary.LittleEndian, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}
func enc_Uint64(value uint64, buff *bytes.Buffer) error {
	err := binary.Write(buff, binary.LittleEndian, Uint64)
	if err != nil {
		return err
	}
	err = binary.Write(buff, binary.LittleEndian, value)
	if err != nil {
		return err
	}
	return nil
}
func dec_Uint64(buff *bytes.Buffer) (uint64, error) {
	var result uint64
	err := binary.Read(buff, binary.LittleEndian, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}
