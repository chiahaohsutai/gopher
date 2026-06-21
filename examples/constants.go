//go:build variables

package main

import (
	"fmt"
	"math"
)

const s string = "string"

func main() {
	fmt.Println(s)

	const n = 500
	fmt.Println(math.Sin(n))

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
}
