package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0 // can also be x or x/2
	for i:= 0; i < 10; i++ {
		pre_z := z
		z -= (z*z -x) / (2*z)
		// If the difference is small enough we got it
		if math.Abs(pre_z - z) < 0.0000000001 {
			fmt.Printf("Got it in %v tries \n", i)
			return z, nil
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
