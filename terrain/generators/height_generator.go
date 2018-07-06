package generators

import "github.com/gaberankin/land-generator/terrain"

// HeightGenerator - interface for generation
type HeightGenerator interface {
	ApplyToTerrain(t terrain.Terrain)
}
