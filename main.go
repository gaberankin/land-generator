package main

import (
	"fmt"

	"github.com/gaberankin/land-generator/terrain"
	"github.com/gaberankin/land-generator/terrain/generators"
)

func main() {
	// rg := generators.NewRandomHeightGenerator("Gabe")
	rg := generators.NewDiamondSquareGenerator("Gabe", 1.1)
	field := terrain.NewTerrain(15, 15)
	rg.ApplyToTerrain(field)

	for x := uint(0); x < field.Width(); x++ {
		for y := uint(0); y < field.Length(); y++ {
			h, _ := field.GetHeight(x, y)
			fmt.Printf("%0.2f	", h)
		}
		fmt.Println("")
	}

}
