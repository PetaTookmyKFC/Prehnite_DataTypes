package datatypes

import (
	"bytes"
	"encoding/binary"
)

func enc_Bool(value bool, buff *bytes.Buffer) error {
	// fmt.Println("Encoding : ", value)

	err := binary.Write(buff, binary.LittleEndian, Bool)
	if err != nil {
		return err
	}
	err = binary.Write(buff, binary.LittleEndian, value)
	if err != nil {
		return err
	}

	return nil
}

func dec_Bool(buff *bytes.Buffer) (bool, error) {
	var result bool
	err := binary.Read(buff, binary.LittleEndian, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}
