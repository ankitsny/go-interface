package main

import (
	"fmt"
	"math"
)

type shape interface{
	area() float64
}

type square struct{
	side float64
}

type rect struct{
	length float64
	breadth float64
}

type circle struct{
	radius float64
}


func (s square) area() float64{
	return s.side * s.side
}


func (r rect) area() float64{
	return r.length * r.breadth
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}



func init(){
	sq := square{10}
	re := rect{12, 23}
	c := circle{17}
	fmt.Println("Square:", sq.area())
	fmt.Println("Rect: ", re.area())
	fmt.Println("Circle: ", c.area())
}