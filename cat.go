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

	/*
		if args[1] == ">"{	//上書きの設定
			files,err := ioutil.ReadFile(args[0])
			if err != nil{
				log.Fatal(err)
			}
			ioutil.WriteFile(args[2],files,0777)

		}
	*/

	for _, path := range args {
		files, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("[" + path + "]")
		fmt.Println(string(files))
	}
}
