# genstack

[![Go Reference](https://pkg.go.dev/badge/github.com/markekraus/genstack/pkg.svg)](https://pkg.go.dev/github.com/markekraus/genstack/pkg)

A Generic Stack (FILO/LIFO) package for Go

```bash
 go get github.com/markekraus/genstack
```

```go
package main

import (
    "fmt"
    genstack "github.com/markekraus/genstack/pkg"
)

type mytype struct {
    a, b int
}

func main() {
    s := genstack.New[int]()
    fmt.Printf("elements: %v, want: %v\n", s.Len(), 0)
    s.Push(3)
    s.Push(2)
    s.Push(90)
    fmt.Printf("elements: %v, want: %v\n", s.Len(), 3)
    for i := s.Len(); i > 0; i-- {
        m := s.Pop().Value
        fmt.Printf("elements: %v, want: %v\n", s.Len(), i)
        fmt.Printf("Value: %v\n", m)
    }

    s2 := genstack.New[*mytype*]()
    e1 := &mytype{1, 2}
    s2.Push(&mytype{1, 2})
    e2 := q.Pop().Value
    fmt.Printf("%v\n", e2 == e1)
    fmt.Printf("%v\n", e2.a == e1.a)
    fmt.Printf("%v\n", e2.b == e1.b)
}
```
