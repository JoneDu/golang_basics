package pointer

import (
	"fmt"
	"testing"
)

func TestIntPointerInc(t *testing.T) {
	a := 22
	IntPointerAddTen(&a)
	fmt.Printf("a: %+v\n", a)
}

func TestSliceMulTow(t *testing.T) {
	slice := []int{1, 2, 3}
	fmt.Println("before mul tow:", slice)
	SliceMulTow(slice)
	fmt.Println("after mul tow:", slice)

	slice1 := []int{45, 6, 90}
	fmt.Println("before slice tow:", slice1)
	SlicePointerMulTow(&slice1)
	fmt.Println("after slice tow:", slice1)
}
