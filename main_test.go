package datatypes

import (
	"reflect"
	"testing"
)

type args struct {
	value any
}
type testStruct struct {
	name     string
	args     args
	wantErr  bool
	wantType DType
}

func Test_String(t *testing.T) {
	tests := []testStruct{
		{
			name: "Testing plain String",
			args: args{
				value: "Testing Plain Test",
			},
			wantErr:  false,
			wantType: String,
		},
		{
			name: "Testing Unicode String",
			args: args{
				value: "❤️",
			},
			wantErr:  false,
			wantType: String,
		},
		{
			name: "Testing Mixed String",
			args: args{
				value: "Testing ❤️",
			},
			wantErr:  false,
			wantType: String,
		},
	}
	_Test_Generic(tests, t)
}
func Test_Bool(t *testing.T) {

	tests := []testStruct{
		{
			name: "Testing Bool",
			args: args{
				value: false,
			},
			wantErr:  false,
			wantType: Bool,
		},
		{
			name: "Testing Bool 2",
			args: args{
				value: true,
			},
			wantErr:  false,
			wantType: Bool,
		},
	}

	_Test_Generic(tests, t)
}
func Test_Number(t *testing.T) {

	tests := []testStruct{
		{
			name: "Testing Int32",
			args: args{
				value: int32(26),
			},
			wantErr:  false,
			wantType: Int32,
		},
		{
			name: "Testing Int8",
			args: args{
				value: int8(26),
			},
			wantErr:  false,
			wantType: Int8,
		},
		{
			name: "Testing Int16",
			args: args{
				value: int16(1066),
			},
			wantErr:  false,
			wantType: Int16,
		},
		{
			name: "Testing Int",
			args: args{
				value: 53180084,
			},
			wantErr:  false,
			wantType: Int64,
		},
		{
			name: "Testing int 32",
			args: args{
				value: int32(5318008),
			},
			wantErr:  false,
			wantType: Int32,
		},
		{
			name: "Testing int 64",
			args: args{
				value: int64(453245623546),
			},
			wantErr:  false,
			wantType: Int64,
		},
		{
			name: "Testing float 32",
			args: args{
				value: float32(45.5),
			},
			wantType: Float32,
			wantErr:  false,
		},
		{
			name: "Testing float 64",
			args: args{
				value: float64(3.14159),
			},
			wantType: Float64,
			wantErr:  false,
		},
	}

	_TestNumber(tests, t)
}
func Test_Array(t *testing.T) {
	tests := []testStruct{{
		name: "Testing Array",
		args: args{
			value: []interface{}{
				"Testing",
				"Testing1",
				"Testing2",
			},
		},
		wantType: Array,
		wantErr:  false,
	},
		{
			name: "TESTING 2dArray",
			args: args{
				value: []interface{}{
					[]interface{}{
						"Testing",
						"Testing1",
						"Testing2",
					},
					[]interface{}{
						"1Testing",
						"1Testing1",
						"1Testing2",
					},
					[]interface{}{
						"2Testing",
						"2Testing1",
						"2Testing2",
					},
				},
			},
			wantType: Array,
			wantErr:  false,
		},
		{
			name: "TESTING a number",
			args: args{
				value: []interface{}{
					int32(5),
					int32(35),
					int32(82),
				},
			},
			wantErr:  false,
			wantType: Array,
		},
	}
	_Test_Generic(tests, t)
}
func Test_Map(t *testing.T) {
	tests := []testStruct{{
		name: "Testing Map 1",
		args: args{
			value: map[string]interface{}{
				"TestingTEXT":    "Testing",
				"TestingNumber":  202,
				"TestingBoolean": true,
			},
		},
		wantType: Map,
		wantErr:  false,
	},
		{
			name: "Testing Map 2",
			args: args{
				value: map[string]interface{}{
					"TestingTEXT":    "Testing",
					"TestingNumber":  202,
					"TestingBoolean": false,
					"TestingMap": map[string]interface{}{
						"Testing-1":      "Testing",
						"Testing-2":      "Testing2",
						"TestingNumnber": 4,
						"HOLY RECURSION BATMAN": []interface{}{
							"Testing",
							"GOD NO!",
						},
						"OMGoodness": map[string]interface{}{
							"Float": float64(3.14159),
							"int32": int32(5318008),
							"Bool":  false,
						},
					},
				},
			},
			wantType: Map,
			wantErr:  false,
		}}

	_TestMaps(tests, t)
}
func TestEncodeDecode(t *testing.T) {

	tests := []struct {
		name     string
		args     args
		wantErr  bool
		wantType DType
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encode(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			dec, tpe, err := Decode(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tpe != tt.wantType {
				t.Errorf("ReturnType() = %v, want %v", tpe, tt.wantType)
				return
			}
			// Dont bother checking if they're equal if there is an error they're not
			if err != nil {
				return
			}

			if !reflect.DeepEqual(tt.args.value, dec) {
				// Check if items are a map
				if tpe == Map || tt.wantType == Map {

					// Workarround for map comparision ( deepEqual doesn't work quite right for maps)
					if CheckMapsEqual(tt.args.value.(map[string]interface{}), dec.(map[string]interface{}), t) {
						return
					} else {
						t.Errorf("DontMatch AsString () = %v, want %v", tt.args.value, dec)
						return
					}
				}

				// Workarrount to chage type from int to int64 for the type of int
				if tpe != Int64 || tt.wantType != Int64 {
					t.Errorf("AreEqual() = %v, want %v", tt.args.value, dec)
					return
				}
				if int64(dec.(int64)) != int64(tt.args.value.(int)) {
					t.Logf("Are not Equal after conversion () = %v, want %v", tt.args.value, dec)
					return
				}
				return
			}
		})
	}
}

// Numbers and maps shouldn't be tested here
func _Test_Generic(tests []testStruct, t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make Encoded Buffer
			got, err := Encode(tt.args.value)
			// Check if an error should have occoured
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Decode Encoded Buffer
			dec, tpe, err := Decode(got)

			// Check if an error should have occoured
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Check if the types match
			if tpe != tt.wantType {
				t.Errorf("ReturnType() = %v, want %v", tpe, tt.wantType)
				return
			}
			// Check if they are equal
			if !reflect.DeepEqual(tt.args.value, dec) {
				t.Errorf("AreEqual() = %v, want %v", tt.args.value, dec)
				return
			}

		})
	}
}
func _TestNumber(tests []testStruct, t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encode(tt.args.value)
			// Check if an error should have occoured
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Decode Encoded Buffer
			dec, tpe, err := Decode(got)
			// Check if an error should have occoured
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Check if the types match
			if tpe != tt.wantType {
				t.Errorf("ReturnType() = %v, want %v", tpe, tt.wantType)
				return
			}

			if reflect.DeepEqual(tt.args.value, dec) {
				return
			}

			// Convert both to int64 for comparison ( not taking into account type)

			var a interface{}
			var b interface{}

			a = tt.args.value
			b = dec

			if GetType(a) == ConvInt {
				a = int64(a.(int))
			}
			if GetType(b) == ConvInt {
				b = int64(b.(int))
			}

			if !reflect.DeepEqual(a, b) {
				t.Errorf("AreEqual() = %v, want %v", a, b)
				return
			}

		})
	}
}
func _TestMaps(tests []testStruct, t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make Encoded Buffer
			got, err := Encode(tt.args.value)
			// Check if an error should have occoured
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Decode Encoded Buffer
			dec, tpe, err := Decode(got)
			// Check if an error should have occoured
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Check if the types match
			if tpe != tt.wantType {
				t.Errorf("ReturnType() = %v, want %v", tpe, tt.wantType)
			}

			if !CheckMapsEqual(tt.args.value.(map[string]interface{}), dec.(map[string]interface{}), t) {
				t.Errorf("AreEqual() = %v, want %v", tt.args.value, dec)
				return
			}

		})
	}
}

func Test_Struct(t *testing.T) {
	type Type1 struct {
		TestingTEXT    string
		TestingNumber  int64
		TestingBoolean bool
	}

	tests := []testStruct{
		{
			name: "Testing Struct",
			args: args{
				value: Type1{
					TestingTEXT:    "Testing",
					TestingNumber:  202,
					TestingBoolean: true,
				},
			},
			wantType: Struct,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Make Encoded Buffer
			got, err := Encode(tt.args.value)
			// Check if an error should have occoured
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			target := new(Type1)

			tpe, err := DecodeStruct(got, target)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tpe != tt.wantType {
				t.Errorf("ReturnType() = %v, want %v", tpe, tt.wantType)
				return
			}

			// a := *target

			if tt.args.value == *target {
				t.Log("Everything Equal")
				return
			} else {
				t.Log("Not Equal")
			}
			// Check if they are equal
			if !reflect.DeepEqual(tt.args.value, target) {
				t.Errorf("AreEqual() = %v, want %v", target, tt.args.value)
				return
			}

		})
	}
}

func CheckMapsEqual(a, b map[string]interface{}, t *testing.T) bool {
	if len(a) != len(b) {
		t.Error("Maps are not equal length")
		return false
	}
	for k, va := range a {
		vb, ok := b[k]
		if !ok {
			t.Error("Maps are not equal ok")
			return false
		}
		if !reflect.DeepEqual(va, vb) {
			if GetType(va) == Map {
				return CheckMapsEqual(va.(map[string]interface{}), vb.(map[string]interface{}), t)
			}
			if GetType(va) == ConvInt {
				va = int64(va.(int))
			}
			if GetType(vb) == ConvInt {
				vb = int64(vb.(int))
			}

			if !reflect.DeepEqual(va, vb) {
				t.Error("ints are not equal deep")
				return false
			}
		}
	}
	return true
}

// func makeEmptyStruct(src interface{}) interface{} {

// 	s := reflect.ValueOf(src).Elem()
// 	res := reflect.New(reflect.TypeOf(struct{}{}))

// 	for i := 0; i < s.NumField(); i++ {
// 		res.
// 	}

// 	return nil
// }
