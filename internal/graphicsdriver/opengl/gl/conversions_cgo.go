// SPDX-License-Identifier: MIT

//go:build !darwin && !windows

package gl

import (
	"unsafe"
)

// #include <stdlib.h>
import "C"

// GoStr takes a null-terminated string returned by OpenGL and constructs a
// corresponding Go string.
func GoStr(cstr *byte) string {
	return C.GoString((*C.char)(unsafe.Pointer(cstr)))
}

// CStr takes a Go string (with or without null-termination)
// and returns the C counterpart.
//
// The returned free function must be called once you are done using the string
// in order to free the memory.
func CStr(str string) (cstr *byte, free func()) {
	cs := C.CString(str)
	return cs, func() {
		C.free(unsafe.Pointer(cs))
	}
}
