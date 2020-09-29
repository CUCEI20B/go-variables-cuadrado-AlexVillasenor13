package main

import "fmt"

func main()  {
	var side uint64
	fmt.Scanln(&side)
	area := side*side
	fmt.Print(area)
}