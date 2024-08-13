package datatypes

import (
	"bytes"
	"encoding/binary"
	"errors"
)

const current = uint32(0)

const (
	Uint8 uint8 = iota
	Uint16
	Uint32
	Uint64
)

// Append to me when i change the handing of decoding and encoding the binary files.
// This is most important if i change the size of the data types as i need more than 255 types.
// This is unlickly, but it might happen or that the decoding of the variables is changes slightly that
// may corrupt some of the results from the binary files.
var Versions = map[uint32]Version{
	0: {Name: "PreDevelopment-0", DataTypeSize: Uint8},
}

type Version struct {
	Name         string
	DataTypeSize uint8
	// BackwardCompatible bool
}

// As this is mainly used for the dataTypes at the moment, getting and setting datatype will be implemented here

func EncodeVersion(buff *bytes.Buffer) error {
	err := binary.Write(buff, binary.LittleEndian, current)
	if err != nil {
		return err
	}
	return nil
}

func v_CheckCanDecode(buff *bytes.Buffer) (bool, error) {
	var ver uint32

	err := binary.Read(buff, binary.LittleEndian, &ver)
	if err != nil {
		return false, err
	}

	// check if the version is the same as the current version
	if ver == current {
		return true, nil
	}

	return false, errors.New("that key is for a different version")
}
