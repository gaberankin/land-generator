package generators

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"

	"github.com/gaberankin/land-generator/terrain"
)

// RandomHeightGenerator - generates height based on simple random seed
type RandomHeightGenerator struct {
	seed string
}

// NewRandomHeightGenerator - returns RandomHeightGenerator with seed set
func NewRandomHeightGenerator(seed string) *RandomHeightGenerator {
	h := md5.New()
	io.WriteString(h, seed)
	d := &RandomHeightGenerator{
		seed: seed,
	}
	return d
}

// Generate - generate height for point x,y
func (r RandomHeightGenerator) generateHeight(x uint, y uint) float64 {
	h := md5.New()
	io.WriteString(h, fmt.Sprintf("%s -> (%d,%d)", r.seed, x, y))
	intSeed := int64(binary.BigEndian.Uint64(h.Sum(nil)))
	s := rand.NewSource(intSeed)
	pointGenerator := rand.New(s)
	return pointGenerator.Float64()
	// pointGenerato
}

// ApplyToTerrain - applies generation algorithm to terrain object
func (r RandomHeightGenerator) ApplyToTerrain(t *terrain.Terrain) error {
	for y := uint(0); y < t.Length(); y++ {
		for x := uint(0); x < t.Width(); x++ {
			if err := t.SetHeight(x, y, r.generateHeight(x, y)); err != nil {
				return err
			}
		}
	}
	return nil
}

// Seed - getter for seed value
func (r RandomHeightGenerator) Seed() string {
	return r.seed
}
