package main

import (
	"fmt"
	"github.com/niwek/kwl-go/functions"
	"github.com/niwek/kwl-go/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))

	fmt.Println(fmt.Sprintf("Return from functions.TestNil(false): %+v", functions.TestNil(false)))
	fmt.Println(fmt.Sprintf("Return from functions.TestNil(true): %+v", functions.TestNil(true)))

}
