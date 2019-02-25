package wurfl

/*
#include <stdlib.h>
*/
import "C"

import "unsafe"

var (
	unsafePointerPool = make(chan unsafe.Pointer, 1000)
)

func getPointer() (p unsafe.Pointer) {
	select {
	case p = <-unsafePointerPool:
	default:
	}
	return p
}

func putPointer(p unsafe.Pointer) {
	select {
	case unsafePointerPool <- p:
	default:
		C.free(p)
	}
}
