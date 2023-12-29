package main

import (
	p "dayone/parser"
	"fmt"
)

func main() {
	res, err := p.ParseTxt("puzzle.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
