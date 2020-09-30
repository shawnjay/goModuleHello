package goModuleHello 

import "fmt"
import mr "goModuleHello/morestrings"
import hr "goModuleHello/hello"

func test_main() {
        A := "Hello, world."
        fmt.Println(A);
        fmt.Println(mr.ReverseRunes(A))
	hr.Hello()
}
