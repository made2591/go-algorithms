package bogosort

import (
	"fmt"
	"github.com/made2591/go-algorithms/sorting/common"
)

func Bogosorting(vals []int) (a []int) {

	for s := common.Sorted(vals) {
		common.Shuffle(vals)
		s = common.Sorted(vals)
	}

	return a

}