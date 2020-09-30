package hello

import "fmt"
import mr "goModuleHello/morestrings"

func Hello() {
	fmt.Println("Hello")
	fmt.Println(mr.ReverseRunes("Hello"))
}
