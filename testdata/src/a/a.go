package a

import (
	"b"
	"b/bsub"

	"github.com/gostaticanalysis/c"
	"github.com/gostaticanalysis/c/csub"
)

func afunc() {
}

func main() {
	b.Func()    // want `b\.Func must not be called`
	_ = b.Func  // OK
	f := b.Func // OK
	f()         // want `b\.Func must not be called`

	//lint:ignore called OK
	b.Func() // OK
	//lint:ignore called OK
	f() // OK

	new(b.Type).Method()          // want `\(\*b\.Type\)\.Method must not be called`
	_ = new(b.Type).Method        // OK
	m := new(b.Type).Method       // OK
	m()                           // want `\(\*b\.Type\)\.Method must not be called`
	(*b.Type).Method(new(b.Type)) // want `\(\*b\.Type\)\.Method must not be called`
	m2 := (*b.Type).Method        // OK
	m2(new(b.Type))               // want `\(\*b\.Type\)\.Method must not be called`

	//lint:ignore called OK
	new(b.Type).Method() // OK
	//lint:ignore called OK
	m() // OK
	//lint:ignore called OK
	(*b.Type).Method(new(b.Type)) // OK
	//lint:ignore called OK
	m2(new(b.Type)) // OK

	bsub.Type{}.Method()   // want `\(b/bsub\.Type\)\.Method must not be called`
	_ = bsub.Type{}.Method // OK
	m3 := bsub.Type{}.Method
	m3()                            // want `\(b/bsub\.Type\)\.Method must not be called`
	(bsub.Type).Method(bsub.Type{}) // want `\(b/bsub\.Type\)\.Method must not be called`
	m4 := (bsub.Type).Method        // OK
	m4(bsub.Type{})                 // want `\(b/bsub\.Type\)\.Method must not be called`

	//lint:ignore called OK
	bsub.Type{}.Method() // OK
	//lint:ignore called OK
	m3() // OK
	//lint:ignore called OK
	(bsub.Type).Method(bsub.Type{}) // OK
	//lint:ignore called OK
	m4(bsub.Type{}) // OK

	c.Func()    // want `github\.com/gostaticanalysis/c\.Func must not be called`
	_ = c.Func  // OK
	g := c.Func // OK
	g()         // want `github\.com/gostaticanalysis/c\.Func must not be called`

	new(c.Type).Method()          // want `\(\*github\.com/gostaticanalysis/c\.Type\)\.Method must not be called`
	_ = new(c.Type).Method        // OK
	m5 := new(c.Type).Method      // OK
	m5()                          // want `\(\*github\.com/gostaticanalysis/c\.Type\)\.Method must not be called`
	(*c.Type).Method(new(c.Type)) // want `\(\*github\.com/gostaticanalysis/c\.Type\)\.Method must not be called`
	m6 := (*c.Type).Method        // OK
	m6(new(c.Type))               // want `\(\*github\.com/gostaticanalysis/c\.Type\)\.Method must not be called`

	csub.Type{}.Method()            // want `\(github\.com/gostaticanalysis/c/csub\.Type\)\.Method must not be called`
	_ = csub.Type{}.Method          // OK
	m7 := csub.Type{}.Method        // OK
	m7()                            // want `\(github\.com/gostaticanalysis/c/csub\.Type\)\.Method must not be called`
	(csub.Type).Method(csub.Type{}) // want `\(github\.com/gostaticanalysis/c/csub\.Type\)\.Method must not be called`
	m8 := (csub.Type).Method        // OK
	m8(csub.Type{})                 // want `\(github\.com/gostaticanalysis/c/csub\.Type\)\.Method must not be called`

	afunc() // OK
}
