/*
Author:			jbreyer
Date:			20170829
Purpose:		Go echo using slices
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
