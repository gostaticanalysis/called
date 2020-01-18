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

<!-- links -->
[godoc]: https://godoc.org/github.com/gostaticanalysis/called
[godoc-badge]: https://img.shields.io/badge/godoc-reference-4F73B3.svg?style=flat-square&label=%20godoc.org

