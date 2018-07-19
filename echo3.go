package main

import (
	"os"
	"fmt"
	"strconv"
)

func main(){
	for i, arg := range os.Args[1:]{
		fmt.Println(strconv.Itoa(i) + " " + arg)
	}
}
