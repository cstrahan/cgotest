# About

This is a proof-of-concept of a Go program (`gomain.go`) which
statically links in a separate C program (`cmain.c`), which in turn
defers to a Go plugin via one of `gomain.go`'s exported symbols
(`DoGreet`).

The `DoGreet` function first constructs a 
(`plugin.go`/`plugin.so`), on which it will later invoke the `Greet`
function. First it constructs an instance of the `api.NumberGetter`
interface, which itself defers to `cmain.c`'s `int get_num();` function.
The `api.Greeter`'s `Greet` function is the invoked.

As you can see, this effectively demonstrates passing handles to C
function from the plugin host (via `interface` wrappers).

Unfortunately, what we can't do is use the `c-archive` build mode and
then link the code into a C program, otherwise we'd run into [this open
issue](https://github.com/golang/go/issues/18123). The solution is to
let Go generate the executable and link to the C code.
