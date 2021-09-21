package main

import "fmt"

func main() {
	// map[key type]valuetype
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "ffffff",
	}

	var colors1 map[string]string

	colors2 := make(map[int]string)
	// add the map key-value pair
	colors2[10] = "#ffffff"
	// delete key 10 pair
	delete(colors2, 10)

	printMap(colors)
	fmt.Println(colors1)
	fmt.Println(colors2)
}

// type of the map
func printMap(c map[string]string) {
	for key, value := range c {

		fmt.Println("Hex code for", key, "is", value)
	}
}
