package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{2, 4, 6} // array
	s1 := []int{2, 4, 6}  // slice
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// Slice define um pedaço de um array
	s2 := a2[1:3]

	s3 := a2[:2] // novo slice, mas aponta para o mesmo array

	fmt.Println(s2, s3)

}
