package main

// #cgo LDFLAGS: cmain.o
// extern int c_main(int argc, char** argv);
// extern int get_num();
import "C"

import (
	"fmt"
	"os"
	"plugin"
	"unsafe"

	"github.com/cstrahan/cgotest/api"
)

func main() {
	fmt.Println("Go main running . . .")

	argc := len(os.Args)
	argv := toCArgv(os.Args)
	C.c_main(C.int(argc), argv)
}

type numGetter struct{}

func (self *numGetter) GetNumber() int {
	return int(C.get_num())
}

//export DoGreet
func DoGreet(name string) {
	plug, err := plugin.Open("plugin.so")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	symGreeter, err := plug.Lookup("Greeter")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var greeter api.Greeter
	greeter, ok := symGreeter.(api.Greeter)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	greeter.Greet(name, &numGetter{})
}

func toCArgv(argv []string) **C.char {
	ptrSize := unsafe.Sizeof(uintptr(0))
	array := (**C.char)(C.malloc(C.size_t(len(argv)) * C.size_t(ptrSize)))

	for idx, arg := range argv {
		elem := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(array)) + ptrSize*uintptr(idx)))
		*elem = C.CString(arg)
	}

	return array
}
