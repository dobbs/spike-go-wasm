package main

import "fmt"

type thing struct {
	id    string
	title string
	name  string
}

func main() {
	fmt.Println("Hello, World!")
	p := thing{
		id:    "random",
		title: "This Sentence No Verb",
		name:  "Nothing to see here",
	}

	fmt.Printf("%#v\n", p)
	fmt.Println("Hello again, World")
}
