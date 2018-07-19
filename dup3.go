package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

/*
Dup3 prints the count and text of lines that
appear more than once in the named input files.
 */

 func main(){
 	counts := make(map[string]int)
 	for _, filename := range os.Args[1:]{
 		data, err := ioutil.ReadFile(filename)
 		if err != nil{
 			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
 			continue
		}

		for _, line := range strings.Split(string(data), "\n"){
			counts[line]++
		}
	}
	for line, n := range(counts){
		if n > 1{
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
 }
