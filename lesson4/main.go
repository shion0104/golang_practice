package main

import "fmt"

func main(){
	var i int = 100

	var i2 int64 = 200

	fmt.Println(i + 50)

	// %Tは型を指定する書式
	fmt.Printf("%T\n",i2)

	fmt.Printf("%T\n",int32(i2))

	fmt.Println(i + int(i2))
}