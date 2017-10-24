package ranpar

import "testing"

func arraySum(s []int64) int64 {
	sum := int64(0)
	for _, e := range s {
		sum += e
	}

	return sum
}

func TestRanPar(t *testing.T) {
	examples := []int64{
		0, 1, 2, 10, 101,
	}

	for _, example := range examples {
		partition, err := RanPar(example)
		if err != nil {
			t.Errorf("%s", err.Error())
		}
		sum := arraySum(partition)
		if sum != example {
			t.Errorf("RanPar(%d) = %v (sum = %d)", example, partition, sum)
		}
	}
}

func TestRanParError(t *testing.T) {
	partition, err := RanPar(-1)
	if err == nil {
		t.Errorf("RanPar did not return error on negative number")
	}
	if partition != nil {
		t.Errorf("RanPar returned a partition on negative number")
	}
}
