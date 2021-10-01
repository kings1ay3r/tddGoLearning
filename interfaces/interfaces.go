package interfaces

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	radius float64
}
type Triangle struct {
	sides [3]float64
}
type Shape interface {
	Area() float64
	Perimeter() float64
}

func (obj Rectangle) Perimeter() float64 {
	return 2 * (obj.Height + obj.Width)
}
func (obj Circle) Perimeter() float64 {
	return 6 * (obj.radius)
}

func (obj Rectangle) Area() float64 {
	return obj.Height * obj.Width
}

func (obj Circle) Area() float64 {
	return obj.radius * obj.radius * 3
}
func (obj Triangle) Perimeter() float64 {
	return obj.sides[0] + obj.sides[1] + obj.sides[2]
}

func (obj Triangle) Area() float64 {
	return obj.sides[0] * obj.sides[1] * obj.sides[2]
}
