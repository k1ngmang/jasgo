package internal

import "fmt"

const StackLimit = 65535
const LocalsLimit = 65535
const ArrayDimensionsLimit = 255
const IndexLimit = 255
const MinSwitchCases = 1
const TableSwitchLowLimit = 255

func ValidateLimit(name string, value, limit int) {
	if value > limit {
		panic(
			fmt.Sprintf("limit of `%s` (%d) is exceeded by value %d.\n", name, limit, value))
	}
}

func ValidateAtLeast(name string, value, limit int) {
	if value < limit {
		panic(
			fmt.Sprintf("`%s` (%d) is less than minimal value (%d).\n", name, value, limit))
	}
}

func ValidateRange(name string, value, min, max int) {
	if value > max || value < min {
		panic(
			fmt.Sprintf("value of `%s` (%d) is not in range %d..%d\n", name, value, min, max))
	}
}
