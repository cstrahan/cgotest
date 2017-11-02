package main

import "C" // required

import (
	"fmt"

	"github.com/cstrahan/cgotest/api"
)

type greeting string

// assert `greeting` is an `api.NumberGetter`
var _ api.Greeter = greeting("")

func (g greeting) Greet(str string, numGetter api.NumberGetter) {
	fmt.Printf("Hello, %s!\n", str)
	fmt.Printf("Your lucky number today is: %d\n", numGetter.GetNumber())
}

// exported
var Greeter greeting
