package a

import (
	"b"
	"b/bsub"
)

func afunc() {
}

func main() {
	b.Func()    // want `b\.Func must not be called`
	_ = b.Func  // OK
	f := b.Func // OK
	f()         // want `b\.Func must not be called`

	new(b.Type).Method()          // want `\(\*b\.Type\)\.Method must not be called`
	_ = new(b.Type).Method        // OK
	m := new(b.Type).Method       // OK
	m()                           // want `\(\*b\.Type\)\.Method must not be called`
	(*b.Type).Method(new(b.Type)) // want `\(\*b\.Type\)\.Method must not be called`
	m2 := (*b.Type).Method        // OK
	m2(new(b.Type))               // want `\(\*b\.Type\)\.Method must not be called`

	bsub.Type{}.Method()   // want `\(b/bsub\.Type\)\.Method must not be called`
	_ = bsub.Type{}.Method // OK
	m3 := bsub.Type{}.Method
	m3()                            // want `\(b/bsub\.Type\)\.Method must not be called`
	(bsub.Type).Method(bsub.Type{}) // want `\(b/bsub\.Type\)\.Method must not be called`
	m4 := (bsub.Type).Method        // OK
	m4(bsub.Type{})                 // want `\(b/bsub\.Type\)\.Method must not be called`

	afunc() // OK
}
