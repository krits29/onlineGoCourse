package main

import "fmt"

func main() {
	var colors map[string]string // one way

	colors = make(map[string]string)  // another way

	colors = map[string]string{  //another way
		"red":   "#ff0000",
		"green": "#ffggg9",
	}

	colors["white"] = "#fff778"

	fmt.Println(colors)
}
