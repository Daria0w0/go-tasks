package main

import "fmt"

func main() {
	var Int8 int8 = -127
	var Int16 int16 = 32767
	var Int32 int32 = -2147483648
	var Int64 int64 = 9223372036854775807

	var Uint8 uint8 = 255
	var Uint16 uint16 = 65535
	var Uint32 uint32 = 4294967295
	var Uint64 uint64 = 18446744073709551615

	var Float32 float32 = 3.1415926535
	var Float64 float64 = 3.141592653589793

	var Compl64 complex64 = complex(3, 4)
	var Compl128 complex128 = complex(1, -1)

	var byteVar byte = 255

	var Str string = "Manul"

	var boolVar bool = true

	fmt.Printf("1. Целочисленные типы:\nint8: %d\nint16: %d\nint32: %d\nint64: %d\n\nuint8: %d\nuint16: %d\nuint32: %d\nuint64: %d\n\n", Int8, Int16, Int32, Int64, Uint8, Uint16, Uint32, Uint64)
	fmt.Printf("2. Числа с плавающей точкой:\nfloat32: %f\nfloat64: %f\n\n", Float32, Float64)
	fmt.Printf("3. Комплексные типы:\ncomplex64: %v\ncomplex128: %v\n\n", Compl64, Compl128)
	fmt.Println("4. Байт: byte:", byteVar)
	fmt.Println("5. Строки: string:", Str)
	fmt.Println("6. Логические: bool:", boolVar)
}
