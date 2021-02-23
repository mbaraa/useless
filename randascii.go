package useless

import (
	"math/rand"
	"time"
)

// RandASCII generates a random string or char
type RandASCII struct {
	randGen *rand.Rand
}

// NewRandASCII returns a new RandASCII instance
func NewRandASCII() *RandASCII {
	return &RandASCII{
		randGen: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// GetRandomChar returns a random ASCII character
func (r *RandASCII) GetRandomChar() uint8 {
	randoms := []func() uint8{
		r.getRandomLowerChar, r.getRandomUpperChar, r.getRandomDigitChar,
	}
	r.randomSeed()

	return randoms[r.randGen.Intn(3)+0]()
}

// GetRandomString returns a random string of a given length
func (r *RandASCII) GetRandomString(length int) string {
	s := ""
	for i := 0; i < length; i++ {
		s += string(r.GetRandomChar())
	}

	return s
}

// getRandomLowerChar well it's written on the box :)
func (r *RandASCII) getRandomLowerChar() uint8 {
	r.randomSeed()
	return uint8(r.randGen.Intn(123-97) + 97)
}

// getRandomUpperChar well it's written on the box :)
func (r *RandASCII) getRandomUpperChar() uint8 {
	r.randomSeed()
	return uint8(r.randGen.Intn(91-65) + 65)
}

// getRandomDigitChar well it's written on the box :)
func (r *RandASCII) getRandomDigitChar() uint8 {
	r.randomSeed()
	return uint8(r.randGen.Intn(58-48) + 48)
}

// randomSeed seeds the random generator
func (r *RandASCII) randomSeed() {
	r.randGen.Seed(time.Now().UnixNano())
}
