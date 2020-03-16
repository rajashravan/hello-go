package main
// use 'package main' for executable commands.
// is hello.go considered a package? I think it is but idk

import (
	"fmt"

	"github.com/rajashravan/hello/morestrings"
  "github.com/google/go-cmp/cmp" // remote repo example
)

func main() {
	fmt.Println("Hello, world.")
  fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
  fmt.Println('1'+2)
  fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
