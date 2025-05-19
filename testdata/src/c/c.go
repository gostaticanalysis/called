package c

import "example.com/foo"

func main() {
	foo.Func()          // want `example.com/foo.Func must not be called`
	foo.Type{}.Method() // want `(example.com/foo.Type).Method must not be called`
}
