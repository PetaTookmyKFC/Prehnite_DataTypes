package datatypes

import (
	"bytes"
	"encoding/binary"
	"errors"
)

func enc_String(value string, buff *bytes.Buffer) error {

	// fmt.Println("Encoding : ", value)

	// Write the binary for the type string
	err := binary.Write(buff, binary.LittleEndian, String)
	if err != nil {
		return err
	}
	// Write the length of the string to be saved
	err = binary.Write(buff, binary.LittleEndian, uint32(len(value)))
	if err != nil {
		return err
	}

	// Write the string to the buffer
	_, err = buff.WriteString(value)
	if err != nil {
		return err
	}

	return nil
}

func dec_String(value *bytes.Buffer) (string, error) {

	var NumberRead uint32
	err := binary.Read(value, binary.LittleEndian, &NumberRead)
	if err != nil {
		return "", err
	}

	if NumberRead <= 0 {
		return "", errors.New("string doesn't have a set length")
	}

	// result := make([]byte, NumberRead)
	result := make([]byte, NumberRead)
	err = binary.Read(value, binary.LittleEndian, &result)
	if err != nil {
		return "", err
	}
	res := string(result)

	return res, nil
}
