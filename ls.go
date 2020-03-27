package main

import (
	"fmt"
//	"os"
	"io/ioutil"
//	"bufio"
	"flag"
	"log"
)


func main(){

	//コマンドライン引数を使用
	flag.Parse()
	if flag.NArg() == 0{
		log.Fatal("引数入力してー！！")
	}
	args := flag.Args()
	text := args[0]
	fmt.Println("パス:",text)

	//標準入力
	/*
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()

	fmt.Println("パス:",text)
	*/

	//ディレクトリの名前を表示
	/*	dir, err := os.Open(text)
	if err != nil{
		fmt.Println(err)
		return
	}

	defer dir.Close()

	fmt.Println(dir.Name())
	*/

	files, err := ioutil.ReadDir(text)
	if err != nil{
		log.Fatal(err)
	}

	for _,file := range files{
		fmt.Println(file.Name())
	}

	
}