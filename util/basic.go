package util

import (
	"encoding/binary"
	"errors"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadTwoBytes(bytes []byte) (uint16, error) {
	if len(bytes) != 2 {
		return 0, errors.New("input length is invalid")
	}
	return binary.LittleEndian.Uint16(bytes), nil
}

func ReadFourBytes(bytes []byte) (uint32, error) {
	if len(bytes) != 4 {
		return 0, errors.New("input length is invalid")
	}

	return binary.LittleEndian.Uint32(bytes), nil
}

func ReadFourBytesBE(bytes []byte) (uint32, error) {
	if len(bytes) != 4 {
		return 0, errors.New("input length is invalid")
	}

	return binary.BigEndian.Uint32(bytes), nil
}

func LiToInt(bytes []byte) int {
	res := 0

	for i := range bytes {
		ii := len(bytes) - 1 - i
		v := int(bytes[ii])
		v = v << (8 * (ii))
		res = res | v
	}

	return res
}
