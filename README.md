# genstack

A Generic Stack (FILO/LIFO) package for Go

```go
func main() {
    s := s.New[int]()
    fmt.Printf("elements: %v, want: %v\n", s.Len(), 0)
    s.Push(3)
    s.Push(2)
    s.Push(90)
    fmt.Printf("elements: %v, want: %v\n", s.Len(), 3)
    for i := s.Len(); i > 0; i-- {
        m := s.Pop()
        fmt.Printf("elements: %v, want: %v\n", s.Len(), i)
        fmt.Printf("Value: %v\n", m.Value)
    }
}
```
