package generators

/**
 * Translated from https://github.com/sebastianrosik/diamond-square-algorithm/, with additional work to support non-square fields
 */

import (
	"fmt"
	"math"

	"github.com/gaberankin/land-generator/terrain"
	"github.com/gaberankin/land-generator/terrain/helpers"
)

// DiamondSquareGenerator - extension of the random height generator that takes nearby points into account when calculating height
type DiamondSquareGenerator struct {
	RandomHeightGenerator
	Noise float64
}

// NewDiamondSquareGenerator -
func NewDiamondSquareGenerator(seed string, noise float64) *DiamondSquareGenerator {
	rhg := NewRandomHeightGenerator(seed)
	d := DiamondSquareGenerator{
		RandomHeightGenerator: *rhg,
		Noise: noise,
	}
	return &d
}

// ApplyToTerrain -
func (dsg DiamondSquareGenerator) ApplyToTerrain(t *terrain.Terrain) error {

	var xys = [][]uint{
		[]uint{0, 0},
		[]uint{t.Width() - 1, 0},
		[]uint{t.Width() - 1, t.Length() - 1},
		[]uint{0, t.Length() - 1},
		[]uint{uint(math.Floor(float64(t.Width()) / 2)), uint(math.Floor(float64(t.Length()) / 2))},
	}
	t.MarkBlanks()
	for _, xy := range xys {
		t.SetHeightIfBlank(xy[0], xy[1], dsg.generateHeight(xy[0], xy[1]))
	}

	dsg.subdivide(t, 0, 0, t.Width(), t.Length(), 1)

	// TODO - at this point, the algorithm will have left untouched cells.  we need to calculate those cells and fill them in.

	t.CleanBlanks()
	return nil
}

func (dsg DiamondSquareGenerator) displace(t *terrain.Terrain, x, y uint, s uint, roughness float64) float64 {
	max := float64(s) / float64(t.Width()+t.Length()) * roughness // idk...
	r := dsg.generateHeight(x, y)
	return (r - 0.5) * max

}

func (dsg DiamondSquareGenerator) subdivide(t *terrain.Terrain, x uint, y uint, sx uint, sy uint, level uint) {
	if sx > 1 || sy > 1 {
		fmt.Printf("[%04d] (%d, %d) [%d, %d]\n", level, x, y, sx, sy)
		halfX := uint(math.Floor(float64(sx) / 2))
		halfY := uint(math.Floor(float64(sy) / 2))

		midpointX := x + halfX
		midpointY := y + halfY

		roughness := dsg.Noise / float64(level)

		// Diamond Stage
		tpLf, _ := t.GetHeight(x, y)
		tpRg, _ := t.GetHeight(x+sx, y)
		btLf, _ := t.GetHeight(x, y+sy)
		btRg, _ := t.GetHeight(x+sx, y+sy)

		midpointValue := (tpLf + tpRg + btRg + btLf) / 4
		midpointValue = midpointValue + dsg.displace(t, midpointX, midpointY, halfX+halfY, roughness)
		midpointValue = helpers.Constrain(midpointValue)

		t.SetHeightIfBlank(midpointX, midpointY, midpointValue)

		//Square stage
		tpX := x + halfX
		tpY := y

		rgX := x + sx
		rgY := y + halfY

		btX := x + halfX
		btY := y + sy

		lfX := x
		lfY := y + halfY

		tVal := helpers.Constrain(((tpLf + tpRg) / 2) + dsg.displace(t, tpX, tpY, halfX+halfY, roughness))
		rVal := helpers.Constrain(((tpRg + btRg) / 2) + dsg.displace(t, rgX, rgY, halfX+halfY, roughness))
		bVal := helpers.Constrain(((btLf + btRg) / 2) + dsg.displace(t, btX, btY, halfX+halfY, roughness))
		lVal := helpers.Constrain(((tpLf + btLf) / 2) + dsg.displace(t, lfX, lfY, halfX+halfY, roughness))

		t.SetHeightIfBlank(tpX, tpY, tVal)
		t.SetHeightIfBlank(rgX, rgY, rVal)
		t.SetHeightIfBlank(btX, btY, bVal)
		t.SetHeightIfBlank(lfX, lfY, lVal)

		dsg.subdivide(t, x, y, halfX, halfY, level+1)
		dsg.subdivide(t, x, midpointY, halfX, halfY, level+1)
		dsg.subdivide(t, midpointX, midpointY, halfX, halfY, level+1)
		dsg.subdivide(t, midpointX, y, halfX, halfY, level+1)
	}
}
