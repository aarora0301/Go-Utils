package BloomFilter

import (
	"testing"
)

func TestBloomFilter(t *testing.T) {
	t.Run("Testing BloomFilter", func(t *testing.T) {
		bf := NewBloomFilter(1000)
		bf.Add([]byte("apple"))
		bf.Add([]byte("banana"))

		t.Log(bf.Contains([]byte("apple")))  // Output: true
		t.Log(bf.Contains([]byte("banana"))) // Output: true
		t.Log(bf.Contains([]byte("cherry"))) //false
		t.Log(bf.Contains([]byte("ppale")))  //false

	})
}

func TestThreadSafeBloomFilter(t *testing.T) {
	t.Run("Testing Thread safe BloomFilter", func(t *testing.T) {
		bf := NewThreadSafeBloomFilter(100)
		bf.Add([]byte("apple"))
		bf.Add([]byte("banana"))

		t.Log(bf.Contains([]byte("apple")))  // Output: true
		t.Log(bf.Contains([]byte("banana"))) // Output: true
		t.Log(bf.Contains([]byte("cherry")))
		t.Log(bf.Contains([]byte("ppale")))

	})
}
