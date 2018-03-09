package main

/*
#cgo linux CFLAGS: -I/usr/include -I. -I../
#include "bass.h"
*/
import "C"
import (
	"GoBass"
	"fmt"
)

func main() {
	fmt.Println(bass.RecordInit(-1))
	fmt.Println(bass.RecordStart(44100, 2, 0, MyRecordCallback))
	select {}
}

func MyRecordCallback(handle C.HRECORD, buffer *C.char, length C.DWORD, user *C.char) bool {
	fmt.Println("...")
	return true
}
