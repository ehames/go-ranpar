package ranpar

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func getKeys(m map[int]struct{}) []int {
	keys := make([]int, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

func generateCoefficients(n int) []int {
	randomGenerator := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	splits := int(math.Floor(math.Log2(float64(n))))

	fmt.Printf("n = %d splits = %d\n", n, splits)

	coefficients := make(map[int]struct{}, splits)

	for i := 0; i < splits; i++ {
		r := randomGenerator.Intn(n)
		coefficients[r] = struct{}{}
	}

	keys := getKeys(coefficients)
	sort.Ints(keys)

	return keys
}

func calculatePartition(n int, coefficients []int) []int {
	partition := make([]int, len(coefficients)+1)

	partition[0] = coefficients[0]
	for i := 1; i < len(coefficients); i++ {
		partition[i] = coefficients[i] - coefficients[i-1]
	}
	partition[len(coefficients)] = n - coefficients[len(coefficients)-1]

	fmt.Printf("partitions = %v\n", partition)
	return partition
}

// RanPar generates a random partition
// the sum of all elements in the partion equals n
func RanPar(n int) ([]int, error) {
	if n < 0 {
		err := fmt.Errorf("cannot partion n < 0 (n = %d)", n)
		return nil, err
	}

	if n == 0 || n == 1 {
		return []int{n}, nil
	}

	coefficients := generateCoefficients(n)
	partition := calculatePartition(n, coefficients)

	return partition, nil
}
