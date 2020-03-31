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

	const bufsize = 1024

	for _, path := range args {

		files, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer files.Close()

		text := make([]byte, bufsize)
		for {
			n, err := files.Read(text)
			if n == 0 { //内容が何もない時のループ脱出
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(text[:n]))
		}
	}
}
