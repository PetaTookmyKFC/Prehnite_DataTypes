package datatypes

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

func enc_Map(value map[string]interface{}, buff *bytes.Buffer) error {
	// Write the binary for the type map
	err := binary.Write(buff, binary.LittleEndian, Map)
	if err != nil {
		return err
	}
	// Loop through all the items in the map
	for key, item := range value {

		// Encode the key
		err = enc_String(key, buff)
		if err != nil {
			return err
		}

		// Encode the value
		err = _Encode(item, buff)
		if err != nil {
			return err
		}
	}

	err = binary.Write(buff, binary.LittleEndian, EOR)
	if err != nil {
		return err
	}

	return nil
}

func dec_Map(buff *bytes.Buffer) (map[string]interface{}, error) {

	var result map[string]interface{} = make(map[string]interface{})

	for buff.Len() > 0 {
		// Decode Key
		key, r, err := _Decode(buff)
		if err != nil {
			return nil, err
		}

		if r == EOR {
			// This is the end of the map
			return result, nil
		}

		if r != String {
			return nil, errors.New("map key is not a string")
		}

		// Decode Value
		value, r, err := _Decode(buff)
		if err != nil {
			return nil, err
		}
		if r == Invalid {
			return nil, errors.New("map value is invalid")
		}
		result[key.(string)] = value
		fmt.Println(key, ":", value)
	}

	// prehnitelogs.Info(fmt.Sprint(result))
	return result, nil
}
