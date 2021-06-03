package main

import "fmt"

// struct
type Vertex struct {
	X, Y int
}

func (v Vertex) Area() int {
	return v.X * v.Y
}

func (v *Vertex) Scale(rate int) {
	v.X *= rate
	v.Y *= rate
}

func Area(v Vertex) int {
	return v.X * v.Y
}

// embedded
type Vertex3D struct {
	Vertex
	Z int
}

func (v Vertex3D) Area() int {
	return v.X * v.Y
}

func (v *Vertex3D) Scale(rate int) {
	v.X *= rate
	v.Y *= rate
	v.Z *= rate
}

func NewVertex3D(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v)
	fmt.Println(Area(v))
	fmt.Println(v.Area())
	v.Scale(10)
	fmt.Println(v)
	fmt.Println(Area(v))
	fmt.Println(v.Area())

	v3d := NewVertex3D(3, 4, 5)
	fmt.Println(v3d)
	fmt.Println(v3d.Area())
}
