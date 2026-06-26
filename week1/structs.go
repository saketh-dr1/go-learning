package main

import "fmt"

type Sport struct {
	name     string
	position string
}

type Person struct {
	name string
	age  uint
	// f    func(string) string
	// favSport Sport // embedded structs
	favSports []Sport
}

func (p Person) getName() string {
	return p.name
}

func (p Person) setName(newName string) {
	p.name = newName
	fmt.Println(p)
}

func struct_main() {
	// p1 := Person{"Tim", 23}

	// var p1 Person = Person{age: 23}
	// p1.name = "Tim"
	// fmt.Printf("%T\n", p1)
	// fmt.Println(p1.age)
	// p1.f = func(x string) string {
	// 	return x + "s"
	// }
	// fmt.Println(p1.f("hello"))
	// value := p1.getName()
	// fmt.Println(value)
	// p1.setName("joey")
	// fmt.Println(p1)

	// p1 := Person{age: 23, name: "dsr", favSport: Sport{"Soccer", "D"}}
	// fmt.Println(p1.favSport.position)

	// slce of a struct type
	p1 := Person{age: 23, name: "dsr", favSports: []Sport{{"Soccer", "D"}}}
	fmt.Println(p1.favSports[0].position)
}
