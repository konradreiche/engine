package main

import (
	"fmt"
	"os"

	"github.com/october93/engine/cmd/rpcgen/parser"
)

func main() {
	p, err := parser.NewParser(os.Getenv("GOFILE"))
	if err != nil {
		fail(err)
	}
	err = p.Parse()
	if err != nil {
		fail(err)
	}
}

func fail(err error) {
	fmt.Println(err)
	os.Exit(1)
}
