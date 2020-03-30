package main

import(
	"fmt"
	"flag"
	"os"
	"log"
)


func main(){

	flag.Parse()
	if flag.NArg() == 0{
		fmt.Println("コマンドライン引数を入力してください")
		return
	}
	args := flag.Args()
	dir := args[0]

	if err := os.Mkdir(dir,0777); err != nil{
		log.Fatal(err)
	}
	
}