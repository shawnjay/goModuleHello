package main 

import "fmt"
import mr "goModuleHello/morestrings"
import hr "goModuleHello/hello"

func main() {
        A := "Hello, world."
        fmt.Println(A);
        fmt.Println(mr.ReverseRunes(A))
	hr.Hello()
}
