package main


import (
  "fmt"
)

type server struct {
  a int
  handlers map[string] func(a int) int
}

func (h *server) Addhandler(code string, handler func(a int) int) {
  h.handlers[code] = handler
}

func handler(a int) int {
  return a*a
}


func main() {
  s := server{3, make(map[string] func(a int) int)}
  s.Addhandler("a", handler)
  r := s.handlers["a"](3)
  fmt.Println(r)
}
