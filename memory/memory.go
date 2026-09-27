// Package memory contains experiments with Go memory and pointers.
package memory

import "fmt"

func TestMemory() {
	x := 10

	fmt.Println(x)

	changeValue(&x)

	fmt.Println(x)
}

func changeValue(number *int) {
	*number = 100
}
