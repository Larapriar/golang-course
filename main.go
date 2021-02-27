package main

import "fmt"

func main() {

	//firs way of declaring a map [keys]values
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
	}
	fmt.Println(colors)

	//second way of declaring a map
	var secondColors map[string]string
	fmt.Println(secondColors)

	//Third way
	colorsThree := make(map[string]string)
	fmt.Println(colorsThree)

	//Add values
	colors["white"] = "#ffffff"

	secondColors = map[string]string{}
	secondColors["white"] = "#ffffff"
	colorsThree["white"] = "#ffffff"
	fmt.Println(colors)
	fmt.Println(secondColors)
	fmt.Println(colorsThree)

}
