# TINJECT

TINJECT is a Go package that creates new instances by string in running time.

Getting Started
===============

## Installing

To start using TINJECT, install Go and run `go get`:

```sh
$ go get -u github.com/projectred/tinject
```

## Regist

Befort using TINJECT, you should regist the Types that you are going to create in runnig time.

```go
type A interface {
	Value() int
}

type DefaultA struct {
	V int
}

func (a *DefaultA) Value() int { return a.V }

func init() {
    // tinject.RegistOptionName isn't necessary if you dont need alias name.func Regist default use type's pkg + type.Name as key.
    if err := tinject.Regist(reflect.TypeOf((*DefaultA)(nil)), tinject.RegistOptionName("defaultA")); err != nil {
        panic(err)
	}
}
```

## Running

```go
package main

import "github.com/projectred/tinject"

func main() {
    // key: pkg + type.Name
    var a A = tinject.NewStructByKeyName("github.com/projectred/tinject.DefaultA", tinject.NewStructKvs(KV{"V", 15})).(A)
    println(a.V) // 15
    // key: alias name
	a = tinject.NewStructByKeyName("defaultA", NewStructKvs(KV{"V", 17})).(A)
    println(a.V) // 17
}
```
