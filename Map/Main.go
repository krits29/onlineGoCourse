package main

import "fmt"

func main() {
	var colors map[string]string //can skip this

	colors = map[string]string{
		"red":   "#ff0000",
		"green": "#ffggg",
	}

	fmt.Println(colors)
}
