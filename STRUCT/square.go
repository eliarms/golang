package main

import (
	"fmt"
	"log"
)

type Square struct {
	Center Point
	Length int
}

//NewSquare returns a new square
func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be > 0")
	}
	s := &Square{
		Center: Point{x, y},
		Length: length,
	}
	return s, nil
}

//Move moves the square
func (s *Square) Move(dx int, dy int) {
	s.CenterMove(dx, dy)
}

//Area return the  square
func (s *Square) Area() int {
	return s.Length * s.Length
}

func main() {
	s, err := NewSquare(1, 1, 10)
	if err != nill {
		log.Fatalf("Error: can't create Square")
	}
	s.Move(2, 3)
	fmt.Printf("%v\n", s)
	fmt.Println(s.Area())
}
