package passwordhandlers

import (
	"bytes"
	"fmt"
)

// Fine is Plain TEXT... You will get FINED... `if in production`!
// Only  use Fine for development

func Get_Fined() Handler {
	return Handler{
		Check_Password: func(password Password, buff []byte) (bool, error) {
			fmt.Println("password: ", password)
			fmt.Println(buff)
			return bytes.Equal([]byte(password), buff), nil
		},
		Enc_Password: func(password Password) ([]byte, error) {
			return []byte(password), nil
		},
	}
}
