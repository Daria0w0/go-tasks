package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("Размер int8: %d байт\n", unsafe.Sizeof(int8(100)))
	fmt.Printf("Размер int16: %d байт\n", unsafe.Sizeof(int16(0)))
	fmt.Printf("Размер int32: %d байт\n", unsafe.Sizeof(int32(0)))
	fmt.Printf("Размер int64: %d байт\n\n", unsafe.Sizeof(int64(0)))
	fmt.Printf("Размер uint8: %d байт\n", unsafe.Sizeof(uint8(0)))
	fmt.Printf("Размер uint16: %d байт\n", unsafe.Sizeof(uint16(0)))
	fmt.Printf("Размер uint32: %d байт\n", unsafe.Sizeof(uint32(0)))
	fmt.Printf("Размер uint64: %d байт\n\n", unsafe.Sizeof(uint64(0)))

	fmt.Printf("Размер float32: %d байт\n", unsafe.Sizeof(float32(0)))
	fmt.Printf("Размер float64: %d байт\n\n", unsafe.Sizeof(float64(0)))

	fmt.Printf("Размер complex64: %d байт\n", unsafe.Sizeof(complex64(0)))
	fmt.Printf("Размер complex128: %d байт\n\n", unsafe.Sizeof(complex128(0)))

	fmt.Printf("Размер byte: %d байт\n\n", unsafe.Sizeof(byte(0)))

	fmt.Printf("Размер string: %d байт\n\n", unsafe.Sizeof("Manul"))

	fmt.Printf("Размер bool: %d байт\n\n", unsafe.Sizeof(true))
}
