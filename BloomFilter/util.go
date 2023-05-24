package BloomFilter

import "math"

// determine bits per item
// formula :>
// m = - (n * log(p)) / (log(2)^2)
// m is the size of the bits array (number of bits)
//n is the expected number of elements to be stored in the filter (capacity)
//p is the desired false positive probability

func optimalSize(capacity int, falsePositiveProbability float64) int {
	size := int(math.Ceil((-float64(capacity) * math.Log(falsePositiveProbability)) / (math.Log(2) * math.Log(2))))

	if size < 1 {
		size = 1
	}
	return size
}

// k = (m / n) * ln(2)
//
// Where:
//
// "k" is the number of hash functions.
// "m" is the number of bits in the Bloom filter.
// "n" is the expected number of elements to be inserted into the filter.
func optimalNumHashers(capacity, size int) int {
	numHashers := int(float64(size) / float64(capacity) * math.Ln2)
	if numHashers < 1 {
		numHashers = 1
	}
	return numHashers
}
