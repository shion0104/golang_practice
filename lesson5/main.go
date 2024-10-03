package main

import "fmt"

func main(){
	var f164 float64 = 2.4
	fmt.Println(f164)

	fl := 3.2
	fmt.Println(f164 + fl)
	fmt.Printf("%T, %T\n", f164, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	fmt.Printf("%T\n", float64(fl32))
}