package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		log.Fatal("引数エラー")
	}
	args := flag.Args()
	path := args[0]

	const bufsize = 1024

	files, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer files.Close()

	text := make([]byte, bufsize)
	for {
		n, err := files.Read(text)
		if n == 0 {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(text[:n]))
	}
}
