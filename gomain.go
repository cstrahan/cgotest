package main

// #cgo LDFLAGS: cmain.o
// extern int c_main();
// extern int get_num();
import "C"

import (
	"fmt"
	"os"
	"plugin"

	"github.com/cstrahan/cgotest/api"
)

func main() {
	fmt.Println("Go main running . . .")
	C.c_main()
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
