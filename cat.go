package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		log.Fatal("引数エラー")
	}
	args := flag.Args()

	for _, path := range args {
		files, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("[" + path + "]")
		fmt.Println(string(files))
	}
}
