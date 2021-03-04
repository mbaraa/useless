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

// GetRandomAlphanumChar returns a random alphanumeric ASCII character
func (r *RandASCII) GetRandomAlphanumChar() uint8 {
	randoms := []func() uint8{
		r.getRandomLowerChar, r.getRandomUpperChar, r.getRandomDigitChar,
	}
	r.randomSeed()

	return randoms[r.randGen.Intn(3)+0]()
}

// GetRandomChar returns a random (alphanumeric and special) ASCII character
func (r *RandASCII) GetRandomChar() uint8 {
	r.randomSeed()
	randSpecialChars := []func() uint8{
		r.getRandomSpecialCharSet1, r.getRandomSpecialCharSet2, r.getRandomSpecialCharSet3,
		r.getRandomSpecialCharSet4, r.getRandomUpperChar, r.getRandomLowerChar, r.getRandomDigitChar}

	return randSpecialChars[r.randGen.Intn(7)+0]()
}

// GetRandomString returns a random alphanumeric ASCII string of a given length
func (r *RandASCII) GetRandomAlphanumString(length int) string {
	return r.getRandomString(length, r.GetRandomAlphanumChar)
}

// GetRandomString returns a random ASCII string of a given length
func (r *RandASCII) GetRandomString(length int) string {
	return r.getRandomString(length, r.GetRandomChar)
}

// getRandomString returns a random ASCII string of a given length and the given random character generator
func (r RandASCII) getRandomString(length int, randCharGen func() uint8) string {
	s := ""
	for i := 0; i < length; i++ {
		s += string(randCharGen())
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

// getRandomSpecialCharSet1 well it's written on the box :)
func (r *RandASCII) getRandomSpecialCharSet1() uint8 {
	r.randomSeed()
	return uint8(r.randGen.Intn(48-33) + 33)
}

// getRandomSpecialCharSet2 well it's written on the box :)
func (r *RandASCII) getRandomSpecialCharSet2() uint8 {
	r.randomSeed()
	return uint8(r.randGen.Intn(65-58) + 58)
}

// getRandomSpecialCharSet3 well it's written on the box :)
func (r *RandASCII) getRandomSpecialCharSet3() uint8 {
	r.randomSeed()
	return uint8(r.randGen.Intn(97-91) + 91)
}

// getRandomSpecialCharSet4 well it's written on the box :)
func (r *RandASCII) getRandomSpecialCharSet4() uint8 {
	r.randomSeed()
	return uint8(r.randGen.Intn(127-123) + 123)
}

// randomSeed seeds the random generator
func (r *RandASCII) randomSeed() {
	r.randGen.Seed(time.Now().UnixNano())
}
