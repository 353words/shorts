package main

import "fmt"

// Zero returns the zero value for type T.
func Zero[T any]() (_ T) {
	return
}

// Max returns the maximal value in values.
func Max[T int | float64](values []T) (T, error) {
	if len(values) == 0 {
		return Zero[T](), fmt.Errorf("Max of empty slice")
	}

	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}

	return m, nil
}

func main() {
	fmt.Println(Max([]int{}))
}
