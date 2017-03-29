package main

import (
	"fmt"
	"github.com/zjohl/go-simple-database-challenge/cmd"
)

func main() {

	parser := cmd.Parser{}
	parser.Parse("SET a 10")
	fmt.Println("10")
}
