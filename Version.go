package datatypes

import (
	"bytes"
	"encoding/binary"
	"errors"
)

const current = uint32(1)

// This was originally used incase the encoding types where to change the data would be updated and recovered to the new format...
// I know this is not used, don't really know what it was made for, but it's here now and here it will stay, including `SizeofUint` and the versions map
const (
	SizeofUint8 uint8 = iota
	SizeOfUint16
	SizeOfUint32
	SizeOfUint64
)

// Append to me when i change the handing of decoding and encoding the binary files.
// This is most important if i change the size of the data types as i need more than 255 types.
// This is unlickly, but it might happen or that the decoding of the variables is changes slightly that
// may corrupt some of the results from the binary files.
var Versions = map[uint32]Version{
	0: {Name: "PreDevelopment-0", DataTypeSize: SizeofUint8},
	1: {Name: "Added-Uints", DataTypeSize: SizeofUint8},
}

type Version struct {
	Name         string
	DataTypeSize uint8
	// BackwardCompatible bool
}

// It is so in the future you can inform the decoder if the data types are backwards compatible or not.
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
