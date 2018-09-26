package bogosort

import (
	"testing"
	"github.com/made2591/go-algorithms/sorting/common"
)

func Bogosorting(vals []int) (a []int) {

	for !common.Sorted(vals) {
		common.Shuffle(vals)
	}
	return vals

}
