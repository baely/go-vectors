package main

import (
	"awesomeProject/vector"
	"fmt"
)

func main() {
	v1 := vector.Create(1, 3)
	v2 := vector.Create(2, 4)

	v10 := vector.Create(0, 0, 0)
	v11 := vector.Create(0, 1, 0)

	v3, err1 := vector.Add(v1, v2)
	v4, err2 := vector.Subtract(v1, v2)
	v5, err3 := vector.DotProduct(v1, v2)
	v6 := vector.Multiply(v1, 10)

	if err1 != nil || err2 != nil || err3 != nil {
		return
	}

	v12, err4 := vector.Add(v1, v10)

	if err4 != nil {
		fmt.Println(err4)
	}

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(v4)
	fmt.Println(v5)
	fmt.Println(v6)
}
