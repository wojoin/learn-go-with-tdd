package main

import "fmt"

func Perimeter(width float64, height float64) float64 {
    return 2*(width + height)
}

func Area(width float64, height float64) float64 {
    return width * height
}

func main() {
	fmt.Println(Perimeter(10.0, 20.0))
}