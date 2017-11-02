package api

type NumberGetter interface {
	GetNumber() int
}

type Greeter interface {
	Greet(str string, numGetter NumberGetter)
}
