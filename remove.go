package main

import (
	"os"
	"flag"
	"log"
)

func main(){

	flag.Parse()
	if flag.NArg() == 0 {
		log.Fatal("引数を入力しなさい")
	}
	args := flag.Args()
	url := args[0]

	if err := os.Remove(url); err != nil{
		log.Fatal(err)
	}
}