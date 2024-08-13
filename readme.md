# DataTypes

This is used to convert a variable into a buffer for writing to a file. The current system can support upto 255 types and can have more types added in the future ( unfortantly by editing the source code. )


## Current Types

* Control Items
    * Invalid DType = iota
    * EOR           // End of recursion ( to signal the end of a repeating data structure ... used for maps)
* Basic
    * Bool
    * String
* Number Types
    * ConvInt //  ( converts to int64 )
    * Int8
    * Int16
    * Int32
    * Int64
    * Float32
    * Float64
* Complex Types // ( This may require extra supervision )
    * Array
    * Map
    * Struct


> [!WARNing]
> Structs are being supported... For encoding to binary the normal encoding will work, but for decoding a pointer must be passed meaning decodeStruct must be used. 

## Buffer table

This is a table showing how the variables are stored within the buffer. There are slightly different layouts depending on the required data to be saved. Also in golang the `int` type is saved as a `int64` to ensure no data is lost.

##### Bool, int8, int16, int32, int64, float32, float64
> [!NOTE]
> The golang type `int` is automaticly converted into a `int64` just to ensure it will fit. 

| Type | Section Name | Description |
|-|-|-|
| uint32|Version|This is used as a reference to the versions map, this is usedf to prevent items being loaded by older / newer version of the program to prevent coruption of data|
| uint8| DType | This is a iota that is used to save the type of data that is stored within the field.|
| ? ? ? | Value | This is what the `binary.Write` saves the data as. All records are written with the formatting of `binary.LittleEndian`

##### String

| Type | Section Name | Description |
|-|-|-|
| uint32|Version|This is used as a reference to the versions map, this is usedf to prevent items being loaded by older / newer version of the program to prevent coruption of data|
| uint8| DType | This is a iota that is used to save the type of data that is stored within the field.|
| uint32 | Length | This is how long the string is in bytes, this is <b>NOT</b> the number of characters in the string as unicode characters result in more than one byte.
| TEXT | Value | This is the string that has been written, it can include unicode characters.


##### Array 

> [!Note]
> Array can be used to store an array of any datatype including arrays. Also it is an artificial limitation of requiring the same type of data to be stored in a single array, this is because it will make it slightly harder to handle a mixed array after retrieving > it from the datatype decoder. ( it would probable work fine, i just dont like mixed arrays )

> [!Warning]
> Numbers need to be passed with a declaired size in order to pass the array test.


 Type | Section Name | Description |
|-|-|-|
| uint32|Version|This is used as a reference to the versions map, this is used to prevent items being loaded by older / newer version of the program to prevent coruption of data|
| uint8| DType | This is a iota that is used to save the type of data that is stored within the field.|
|uint32| Length| This is now many items are in the array.
| uint8| Dtype | This is a iota that is used to save the type of data stored in the array.|
|[]byte|Value|This is the value that is filled with other Datatypes, except they dont include `Version`. 


##### Map 
> [!NOTE]
> Maps are used to store structs within a key. When retrieving the data these is no checking if the data is in the correct type.

|Type|Section Name |Description|
|-|-|-|
| uint32|Version|This is used as a reference to the versions map, this is used to prevent items being loaded by older / newer version of the program to prevent coruption of data|
| uint8| DType | This is a iota that is used to save the type of data that is stored within the field. ( MAP ) |
| START-REPEATING | This is repeating complex| This is here to represent a repeating structure. This is repeated for each item being saved. |
| uint32 | Length | How long the name of the key is. ( This a type of string ) 
| TEXT | Value | The text value for the key|
| []byte | Value | This is the value stored by the map, this contains the binary compiled datatype |
| uint8| DType | This is the iota that is used to inform the data of a end of recusion, this is used to inform the map that the next entries are not contained within itself, this is for maps within maps |

##### Struct
> [!NOTE]
> Structures are saved using the map data structure after including the version and type `struct`. As a result of this, the user MUST pass the correct struct type as a pointer object to the decode function, as the decoder doesn't know anything about the required parameters.


> [!NOTE]
> I dont know how to make tests to automagically create and test the struct decoding. ( without manually creating the pointer manually. Should probably just use map


|Type|Section Name |Description|
|-|-|-|
| uint32|Version|This is used as a reference to the versions map, this is used to prevent items being loaded by older / newer version of the program to prevent coruption of data|
| uint8| DType | This is a iota that is used to save the type of data that is stored within, in this case this is used to inform the decoder that a conversion to a passed pointer is required after the map is decoded. ( STRUCT )
| []byte | VALUE | This is the buffer of the data type `Map` as structs can be saved as maps
