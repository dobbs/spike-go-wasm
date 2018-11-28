package main

import "fmt"

type thing struct {
	id    string
	title string
	name  string
}

func helloVariables() {
	fmt.Println("Hello, Variables")
	p := thing{
		id:    "random",
		title: "This Sentence No Verb",
		name:  "Nothing to see here",
	}

	fmt.Printf("%#v\n", p)
}

func helloPointers() {
	fmt.Println("Hello, Pointers")
	pointy := 20
	fmt.Printf("%#v\t%#v\n", pointy, &pointy)

	pointypointy := &pointy
	fmt.Printf("#v pointypointy: %#v\n", pointypointy)
	fmt.Printf("address: %+v\nvalue: %+v\npoints:%+v\n", &pointypointy, pointypointy, *pointypointy)
}

func helloControlFlow() {
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

func helloSlices() {
	fmt.Println("Hello, Slices.  Let's show how to break your stuff.")

	groupA := make([]string, 4, 8)
	groupA[0] = "foo"
	groupA[1] = "bar"
	groupA[2] = "baz"
	groupA[3] = "biz"

	fmt.Printf("groupA before: %#v\n", groupA)

	groupB := groupA[1:3]
	fmt.Printf("groupB before: %#v\n", groupB)

	groupB[1] = "BAZ!" //DANGER

	fmt.Printf("groupA[3] changed: %#v\n", groupA)
	fmt.Printf("groupB: %#v\n", groupB)

	fmt.Println("Hello again, Slices.  Iteration.")
	for i, value := range groupA {
		fmt.Printf("groupA[%d] = '%s'\n", i, value)
	}
	for i, value := range groupB {
		fmt.Printf("groupB[%d] = '%s'\n", i, value)
	}
}

func main() {
	helloVariables()
	helloPointers()
	helloControlFlow()
	helloSlices()
}
