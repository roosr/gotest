package test

import (
	"fmt"
)

type B struct {
	c int
}

type A struct {
	c int
}

func (a B) Error() string {
	return "A"
}

func (a A) Error() string {
	return "A"
}

func GetA() (error, error) {

	var e0 error = A{3}
	var e2 error = B{3}

	if e0 == e2 {
		fmt.Println("r0 and r2 same shapes")
	}

	return e0, e2
}
