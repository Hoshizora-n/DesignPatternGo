package prototype

import (
	"design-pattern/5prototype/model"
	"fmt"
)

func Run() {
	circle := model.NewCircle()
	square := model.NewSquare()

	fmt.Printf("%+v\n", circle)
	fmt.Printf("%+v\n", square)

	newCircle := circle.Clone()
	circle.Circumference = 45
	fmt.Printf("%+v | %+v\n", newCircle, circle)

	// example shallow copy, when update newcircle, circle updated too
	newCircle1 := circle
	newCircle1.Circumference = 99
	fmt.Printf("%+v | %+v\n", newCircle1, circle)

	// example creating new object, can't assign private variable
	newCircle2 := &model.Shape{
		Circumference: circle.Circumference,
	}
	newCircle2.Circumference = 88
	fmt.Printf("%+v | %+v\n", newCircle2, circle)

}
