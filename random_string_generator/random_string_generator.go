package random_string_generator

import (
	"math/rand"
)

const MAX_SIZE int = 5

type RandomStringGenerator struct {
	CharSet string
	Length  int
}

func CreateGenerator(charSet string, length int) *RandomStringGenerator {
	return &RandomStringGenerator{CharSet: charSet, Length: length}
}

func (generator *RandomStringGenerator) GenerateRandomString() string {
	length := generator.Length
	if length > MAX_SIZE {
		length = MAX_SIZE
	}

	result := make([]byte, length)
	for i := 0; i < len(result); i++ {
		index := rand.Intn(len(generator.CharSet) - 1)
		result[i] = generator.CharSet[index]
	}
	return string(result)
}
