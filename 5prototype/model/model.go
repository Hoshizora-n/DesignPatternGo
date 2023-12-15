package model

type ShapeI interface {
	GetId() ShapeType // Get the Shape Id
	// PrintTypeProp()   // used for printing proerty values of the Shape
	Clone() Shape // used for getting DeepCopy
}

type ShapeType int // Define Custom type here

const (
	CircleType ShapeType = 1
	SquareType ShapeType = 2
)

type Shape struct {
	shape         ShapeType
	Width         int
	Height        int
	Circumference int
}

func NewCircle() *Shape {
	return &Shape{
		shape:         CircleType,
		Circumference: 30,
	}
}

func NewSquare() *Shape {
	return &Shape{
		shape:  SquareType,
		Width:  15,
		Height: 15,
	}
}

func (s Shape) GetId() ShapeType {
	return s.shape
}

func (s Shape) Clone() *Shape {
	return &Shape{
		shape:         s.shape,
		Width:         s.Width,
		Height:        s.Height,
		Circumference: s.Circumference,
	}
}
