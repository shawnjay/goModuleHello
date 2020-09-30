package goModuleHello

import "fmt"
import mr "goModuleHello/morestrings"

func basicPrint() {
	fmt.Println("basic")
	fmt.Println(mr.ReverseRunes("basic"))
}
