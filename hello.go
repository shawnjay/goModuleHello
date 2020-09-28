package main

import "fmt"
import mr "goModule/hello/morestrings"

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(mr.ReverseRunes("!oG, olleH"))
}
