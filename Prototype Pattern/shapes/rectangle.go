package shapes

type Rectangle struct {
	Length  float64
	Breadth float64
}

func (r *Rectangle) Clone() Shape {
	return &Rectangle{
		Length:  r.Length,
		Breadth: r.Breadth,
	}
}

func (r *Rectangle) GetPerimeter() float64 {
	perimeter := 2 * (r.Length + r.Breadth)
	return perimeter
}

func (r *Rectangle) GetArea() float64 {
	area := r.Length * r.Breadth
	return area
}

func (r *Rectangle) SetMeasurements(args ...float64) {
	r.Length = args[0]
	r.Breadth = args[1]
}
