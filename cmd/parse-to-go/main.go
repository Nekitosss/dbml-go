package main

import (
	"fmt"
	"os"

	"github.com/Nekitosss/dbml-go-n/parser"
	"github.com/Nekitosss/dbml-go-n/scanner"
)

func main() {
	f, _ := os.Open("test.dbml")
	s := scanner.NewScanner(f)
	parser := parser.NewParser(s)
	dbml, err := parser.Parse()
	fmt.Printf("%#v, %v", dbml, err)
}
