package passwordhandlers

type Password string
type Handler struct {
	Check_Password func(password Password, buff []byte) (bool, error)
	Enc_Password   func(value Password) ([]byte, error)
}

type Names uint8

// var NumberNames uint8

const (
	// Fined is plain text... You will get Fined!
	Fined Names = iota
	// Default --
	bcrypt

	passLen uint8 = iota
)

var NumberNames = passLen
var HandlerList = map[Names]Handler{
	// Plain Text --- DO NOT USE ---
	Fined: Get_Fined(),
}

// It is better to use the predefined handlers, or edit the source to add your own.
// this is just because it's badly coded. ( Technically usable )
func Add(handler Handler) uint8 {
	HandlerList[Names(NumberNames)] = handler
	// Increate the number of saved names
	NumberNames++
	return NumberNames - 1
}

var Default = Fined
