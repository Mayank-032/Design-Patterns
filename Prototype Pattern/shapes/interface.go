package shapes

type Shape interface {
	Clone() Shape
	SetMeasurements(args ...float64)
	GetPerimeter() float64
	GetArea() float64
}
