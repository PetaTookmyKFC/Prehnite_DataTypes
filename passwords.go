package datatypes

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	passwordhandlers "github.com/PetaTookmyKFC/Prehnite_DataTypes/PasswordHandlers"
)

// Passwords use bycript hashs are not returned to user via dec_Password.
func enc_Password(value passwordhandlers.Password, buff *bytes.Buffer) error {
	// Get the default item

	// Write The type as Password
	err := binary.Write(buff, binary.LittleEndian, Password)
	if err != nil {
		return err
	}

	// Write the password handler id to the file
	err = binary.Write(buff, binary.LittleEndian, passwordhandlers.Default)
	if err != nil {
		return err
	}

	// encode the password
	data, err := passwordhandlers.HandlerList[passwordhandlers.Default].Enc_Password(value)
	if err != nil {
		return err
	}

	fmt.Printf("Encoded Data Length %v \n", len(data))

	// write the length of the password
	err = binary.Write(buff, binary.LittleEndian, uint32(len(data)))
	if err != nil {
		return err
	}

	// write the data to the buffer
	_, err = buff.Write(data)
	if err != nil {
		return err
	}

	return nil
}

type PassCompareFunc func(password passwordhandlers.Password) (bool, error)

// Password hashs should not be returned!
// Returns the function to check if the password is correct
// func check_Password(value *bytes.Buffer) (func(password string) (bool, error), error) {
func check_Password(value *bytes.Buffer) (PassCompareFunc, error) {
	// Get the password handler
	var enc_Type uint8
	err := binary.Read(value, binary.LittleEndian, &enc_Type)
	if err != nil {
		return nil, errors.New("cant read the password handler type")
	}
	fmt.Printf("ENC_TYPE == %d \n", enc_Type)

	// Read the size of the password bytes
	var NumberRead uint32
	err = binary.Read(value, binary.LittleEndian, &NumberRead)
	if err != nil {
		// failed to read  password length
		return nil, errors.New("Cant read passworld length!")
	}

	if NumberRead <= 0 {
		return nil, errors.New("passsword doesn't have a set length")
	}

	fmt.Printf("Bufferlength == %d \n", NumberRead)
	fmt.Printf("Your handler has an ID of %d \n", enc_Type)

	buff := make([]byte, NumberRead)

	err = binary.Read(value, binary.LittleEndian, &buff)
	// _, err = value.Read(buff)
	if err != nil {
		return nil, err
	}

	// return string(buff[0:]), nil

	return (func(password passwordhandlers.Password) (bool, error) {
		return passwordhandlers.HandlerList[passwordhandlers.Names(enc_Type)].Check_Password(passwordhandlers.Password(password), buff)
	}), nil

}
