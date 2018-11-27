package main

import "fmt"

type thing struct {
	id    string
	title string
	name  string
}

func main() {
	fmt.Println("Hello, Variables")
	p := thing{
		id:    "random",
		title: "This Sentence No Verb",
		name:  "Nothing to see here",
	}

	fmt.Printf("%#v\n", p)

	fmt.Println("Hello, Pointers")
	pointy := 20
	fmt.Printf("%#v\t%#v\n", pointy, &pointy)

	pointypointy := &pointy
	fmt.Printf("#v pointypointy: %#v\n", pointypointy)
	fmt.Printf("address: %+v\nvalue: %+v\npoints:%+v\n", &pointypointy, pointypointy, *pointypointy)

	fmt.Println("Hello, Control Flow")
	name := "Barney"
	age := 32

	switch name {
	case "George", "Marty":
		fmt.Println("Hey, McFly!")
	case "Fred", "Barney":
		fmt.Println("do you want a Brontosaurus Burger?")
	}

	if age > 6 {
		fmt.Println("Here's a cookie")
	}

}
