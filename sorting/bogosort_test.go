package bogosort

import (
	"testing"
	"github.com/made2591/go-algorithms/sorting/common"
)

func TestBogosorting(t *testing.T) {

	ordered := Bogosorting([]int{3, 2, 1})

	if common.Equal(ordered, []int{3, 2, 1}) {
		t.Errorf("Bogosorting was incorrect, got: %d, want: %d.", ordered, []int{3, 2, 1})
	}

}