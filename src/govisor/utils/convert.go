package utils

import ()

func Int16_Uint16(i int16) uint16 {
	return uint16(i) & 0xffff
}

func Uint16_Int16(u uint16) int16 {
	return int16(u)
}

func Int32_Uint32(i int32) uint32 {
	return uint32(i) & 0xffffffff
}

func Uint32_Int32(u uint32) int32 {
	return int32(u)
}
