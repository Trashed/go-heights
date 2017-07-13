package hmapgen

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	roughness = 0.7
)

// Terrain object holds the data needed for generating the height map.
type Terrain struct {
	size int
	max  int
	data [][]float64
}

// New creates and returns a new Terrain object.
func New(detail int) *Terrain {
	t := &Terrain{}
	t.size = int(math.Pow(2.0, float64(detail))) + 1
	t.max = t.size - 1
	t.data = make([][]float64, t.size)
	for i := range t.data {
		t.data[i] = make([]float64, t.size)
	}
	return t
}

// Generate generates the data for height map. First, the corner values are
// randomized after which the Diamon-Square algorithm is used to generate rest
// of the data points.
func (t *Terrain) Generate() {
	rand.Seed(time.Now().UTC().UnixNano())

	// Set corner values
	t.data[0][0] = randInRange()
	t.data[t.max][0] = randInRange()
	t.data[0][t.max] = randInRange()
	t.data[t.max][t.max] = randInRange()

	// Divide array recursively into smaller squares.
	divide(t, t.max)
}

// Data returns terrain's height map info.
func (t *Terrain) Data() [][]float64 {
	return t.data
}

func divide(t *Terrain, size int) {
	x, y, half := 0, 0, size/2
	scale := roughness * float64(size)

	if half < 1 {
		return
	}

	// Square step
	for y = half; y < t.max; y += size {
		for x = half; x < t.max; x += size {
			square(t, x, y, half, rand.Float64()*scale*2-scale)
		}
	}

	// Diamond step
	for y = 0; y <= t.max; y += half {
		for x = (y + half) % size; x <= t.max; x += size {
			diamond(t, x, y, half, rand.Float64()*scale*2-scale)
		}
	}

	divide(t, size/2)
}

func diamond(t *Terrain, x, y, size int, offset float64) {
	fmt.Printf("== diamond ==\nterrain: %v\nx: %d, y: %d, size: %d, offset: %f\n\n", t, x, y, size, offset)

	lx := t.data[x-size][y] // left
	rx := t.data[x+size][y] // right
	ty := t.data[x][y-size] // top
	by := t.data[x][y+size] // bottom

	av := avg(lx, rx, ty, by)
	t.data[x][y] = av + offset

	fmt.Printf("== diamond ==\nterrain: %v\nx: %d, y: %d, size: %d, offset: %f\n\n", t, x, y, size, offset)
}

func square(t *Terrain, x, y, size int, offset float64) {
	fmt.Printf("== square ==\nterrain: %v\nx: %d, y: %d, size: %d, offset: %f\n\n", t, x, y, size, offset)

	a := t.data[x-size][y-size] // upper left
	b := t.data[x+size][y-size] // upper right
	c := t.data[x-size][y+size] // lower left
	d := t.data[x+size][y+size] // lower right

	av := avg(a, b, c, d)
	t.data[x][y] = av + offset

	fmt.Printf("== square ==\nterrain: %v\nx: %d, y: %d, size: %d, avg: %f, offset: %f\n\n", t, x, y, size, av, offset)
}

func avg(vals ...float64) float64 {

	sum, l := 0.0, len(vals)
	for _, val := range vals {
		sum += val
	}

	return sum / float64(l)
}

func randInRange() (rnd float64) {

	r := rand.Intn(200)

	if r < 100 || r > 100 {
		rnd = float64(r-100) / 100.0
	} else {
		rnd = 0.5
	}

	return
}
