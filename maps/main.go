package main

import "fmt"


func main() {


	colors := map[string]string{
		"red": "#ff0000",
		"green": "#00ff00",
		"blue": "#0000ff",
		"white": "#ffffff",
	}

	fmt.Println(colors)
	printMap(colors)
}

func printMap(cm map[string]string) {
	for c, h := range(cm) {
		fmt.Println(c,h)
	}
}