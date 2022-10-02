package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(slice) //[1 2 3 4 5]

	slice = append(slice, 6)
	fmt.Println(slice) //[1 2 3 4 5 6]

	fmt.Println(len(slice)) //6

	fmt.Println(slice[1:3]) //[2 3]

}
