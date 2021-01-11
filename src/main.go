package main

import (
	"os"

	"github.com/atotto/clipboard"
)

func main() {

	arg := os.Args[1]
	clipboard.WriteAll(arg)

}
