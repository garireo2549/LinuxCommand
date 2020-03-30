package main

import (
	"flag"
	"log"
	"os"
)

func main() {

	flag.Parse()
	if flag.NArg() != 2 {
		log.Fatal("引数を二つ入力して！")
	}
	args := flag.Args()
	file1 := args[0]
	file2 := args[1]

	if err := os.Rename(file1, file2); err != nil {
		log.Fatal(err)
	}

}
