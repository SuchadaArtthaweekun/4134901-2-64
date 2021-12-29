package structs

import (
	"fmt"
)

type Student struct{
	id int
	name string
	age int
}

func Learn()  {
	
	james := Student{
		id:1,
		name:"James",
		age:15,
	}

	students := []Student{
		{id:2, name:"John", age: 20},
		{id:3, name:"Jeaw", age: 21},
	}

	fmt.Println(james.name)
	fmt.Println(students[1].name)

	v := Vertex{3.21,4.3}
	fmt.Printf("%v",v.Area()) 

	fmt.Println(v)
	v.Scale(10)
	fmt.Println(v)
}

/////////////////////////// Area /////////////////////////////////
type Vertex struct{
	X,Y float32
}

func (v Vertex)Area() float32{
	return v.X * v.Y
}

func (v *Vertex)Scale(f float32){
	v.X = v.X * f
	v.Y = v.Y * f
}