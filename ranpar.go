package ranpar

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type int64arr []int64

func (a int64arr) Len() int {
	return len(a)
}

func (a int64arr) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a int64arr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// RanPar generates a random partition
// the sum of all elements in the partion equals n
func RanPar(n int64) ([]int64, error) {
	if n < 0 {
		err := fmt.Errorf("cannot partion n < 0 (n = %d)", n)
		return nil, err
	}

	if n == 0 || n == 1 {
		return []int64{n}, nil
	}

	randomGenerator := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	//	splits := int64(math.Floor(math.Log2(float64(n))))
	splits := n
	rangeUpperLimit := make(int64arr, splits)

	for i := int64(0); i < splits; i++ {
		r := randomGenerator.Int63n(n)
		rangeUpperLimit[i] = r
	}

	sort.Sort(rangeUpperLimit)

	partitions := make([]int64, splits+1)
	partitions[0] = rangeUpperLimit[0]

	for i := 1; i < rangeUpperLimit.Len(); i++ {
		partitions[i] = rangeUpperLimit[i] - rangeUpperLimit[i-1]
	}
	partitions[rangeUpperLimit.Len()] = n - rangeUpperLimit[rangeUpperLimit.Len()-1]

	return partitions, nil
}
