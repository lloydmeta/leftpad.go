package main

import (
	"fmt"

	"github.com/lloydmeta/leftpad.go/pkg"
)

func main() {
	fmt.Println(pkg.Leftpad("Hello world", 20, ' '))
}
