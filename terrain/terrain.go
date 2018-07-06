package terrain

import (
	"fmt"
)

type Terrain struct {
	field  [][]float64
	width  uint
	length uint
	// generator Generator
}

// NewTerrain - constructor for terrain
func NewTerrain(w uint, l uint) *Terrain {
	var arr = [][]float64{}
	for i := uint(0); i < w; i++ {
		arr = append(arr, []float64{})
		for j := uint(0); j < l; j++ {
			arr[i] = append(arr[i], 0)
		}
	}
	return &Terrain{
		field:  arr,
		width:  w,
		length: l,
	}
}

// Width - width accessor
func (t Terrain) Width() uint {
	return t.width
}

// Length - length accessor
func (t Terrain) Length() uint {
	return t.length
}

// SetHeight - add a position pointer
func (t *Terrain) SetHeight(x uint, y uint, h float64) error {
	if x >= t.width {
		return fmt.Errorf("Invalid x position %d.  must not be greater than or equal to width %d", x, t.width)
	}
	if y >= t.length {
		return fmt.Errorf("Invalid y position %d.  must not be greater than or equal to length %d", y, t.length)
	}
	t.field[x][y] = h
	return nil
}

// SetHeightIfBlank - similar to SetHeight, but only updates height if height value is considered blank.
func (t *Terrain) SetHeightIfBlank(x uint, y uint, h float64) error {
	if x >= t.width {
		return fmt.Errorf("Invalid x position %d.  must not be greater than or equal to width %d", x, t.width)
	}
	if y >= t.length {
		return fmt.Errorf("Invalid y position %d.  must not be greater than or equal to length %d", y, t.length)
	}
	if t.field[x][y] < 0.0 {
		t.field[x][y] = h
	}
	return nil
}

// MarkBlanks - should be used by terrain randomizers that need a 'blank' state.  should be used in conjunction with SetHeightIfBlank()
func (t *Terrain) MarkBlanks() {
	for x, w := uint(0), t.Width(); x < w; x++ {
		for y, l := uint(0), t.Length(); y < l; y++ {
			t.field[x][y] = -1.0
		}
	}
}

// CleanBlanks - sets untouched blank elements to 0.0 if they are blank.
func (t *Terrain) CleanBlanks() {
	for x, w := uint(0), t.Width(); x < w; x++ {
		for y, l := uint(0), t.Length(); y < l; y++ {
			if t.field[x][y] < 0.0 {
				t.field[x][y] = 0.0
			}
		}
	}
}

// GetHeight - get z value for particular x/y position
func (t Terrain) GetHeight(x uint, y uint) (float64, error) {
	if x >= t.width {
		return 0, fmt.Errorf("Invalid x position %d (index-out-of-range).  must not be greater than or equal to width %d", x, t.width)
	}
	if y >= t.length {
		return 0, fmt.Errorf("Invalid y position %d (index-out-of-range).  must not be greater than or equal to length %d", y, t.length)
	}
	return t.field[x][y], nil
}

func (t Terrain) String() string {
	return fmt.Sprintf("Terrain (%d, %d)", t.width, t.length)
}
