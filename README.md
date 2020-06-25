# called

[![godoc.org][godoc-badge]][godoc]

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

Analyzers ignore nodes which are annotated by [staticcheck's style comments](https://staticcheck.io/docs/#ignoring-problems) as belows.
A ignore comment includes analyzer names and reason of ignoring checking.
If you specify `called` as analyzer name, all analyzers ignore corresponding code.

```go
package main

import "log"

func main() {
	//lint:ignore called reason
	log.Fatal("hoge")
}
```

<!-- links -->
[godoc]: https://godoc.org/github.com/gostaticanalysis/called
[godoc-badge]: https://img.shields.io/badge/godoc-reference-4F73B3.svg?style=flat-square&label=%20godoc.org

