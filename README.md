# called

[![pkg.go.dev][gopkg-badge]][gopkg]

`called` called find callings specified by called.funcs flag.

```go
package main

import "log"

func main() {
	log.Fatal("hoge")
}
```

```sh
$ go vet -vettool=`which called` -called.funcs="log.Fatal" main.go
./main.go:6:11: log.Fatal must not be called
```

## Ignore Checks

Analyzers ignore nodes annotated with [staticcheck's style comments](https://staticcheck.io/docs/#ignoring-problems) as below.
An ignore comment includes the analyzer names and a reason for disabling the check.
If you specify `called` as an analyzer name, all analyzers ignore the corresponding code.

```go
package main

import "log"

func main() {
	//lint:ignore called reason
	log.Fatal("hoge")
}
```

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gostaticanalysis/called
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gostaticanalysis/called?status.svg
