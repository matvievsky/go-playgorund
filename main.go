package main

// #include <helloworld.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	C.cHello()
	myStr := C.CString("LOL KEK CHEBUREK")
	defer C.free(unsafe.Pointer(myStr))
	C.printMessage(myStr)
	fmt.Println("Hello World")
}
