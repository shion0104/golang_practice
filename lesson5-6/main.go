package main

import "fmt"

func main() {
	var x interface{}
	fmt.Println(x)

	fmt.Printf("%T\n", x)

	x = 1
	fmt.Println(x)

	x = 3.14
	fmt.Println(x)

	x = "A"
	fmt.Println(x)

	x = [3]int{1, 2, 3}
	fmt.Println(x)
 
	// インターフェイス型に計算式は使えない
	// x = 2
	// fmt.Println(x + 3)

	fmt.Printf("%T\n", x)
}