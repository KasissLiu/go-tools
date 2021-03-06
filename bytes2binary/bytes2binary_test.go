package Bytes2Binary

import (
	"fmt"
	"testing"
)

func TestByte2Binary(t *testing.T) {
	a := ByteToBinaryString('0')
	fmt.Println(a)
	// Output: 00110000
}

func TestBytes2Binary(t *testing.T) {
	a := BytesToBinaryString([]byte{'0', '1'})
	fmt.Println(a)
	// Output: 00110000 00110001
}

func TestBinary2Byte(t *testing.T) {
	a := BinaryStringToBytes("00110001")
	fmt.Printf("%s\n", string(a))
	// Output: 1
}
