package main

import "fmt"

type Rectangle struct {
	Width float64
	Height float64
}

func Perimeter(r Rectangle) float64 {
    return 2*(r.Width + r.Height)
}

func Area(r Rectangle) float64 {
    return r.Width * r.Height
}

func main() {
	fmt.Println(Perimeter(Rectangle{10.0, 20.0}))
}