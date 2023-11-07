package shapes

type Circle struct {
	Radius float64
}

func (c *Circle) Clone() Shape {
	return &Circle{
		Radius: c.Radius,
	}
}

func (c *Circle) GetPerimeter() float64 {
	perimeter := 2 * 3.14 * c.Radius
	return perimeter
}

func (c *Circle) GetArea() float64 {
	area := 3.14 * c.Radius * c.Radius
	return area
}

func (c *Circle) SetMeasurements(args ...float64) {
	c.Radius = args[0]
}