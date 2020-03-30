package main

import (
	//	"os"
	"flag"
	"log"
	"fmt"
	"io/ioutil"
)

func main(){
	flag.Parse()
	if flag.NArg() == 0{
		log.Fatal("引数エラー")
	}
	args := flag.Args()
/*
	const bufsize = 1024

	files,err := os.Open("C:/gwork/LinuxMv/test.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer files.Close()

	text := make([]byte,bufsize)
	for {
		n,err := files.Read(text)
		if n == 0{
			break
		}
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(string(text[:n]))
	}
*/
	for _,path := range args{
		files,err := ioutil.ReadFile(path)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println("["+path+"]")
		fmt.Println(string(files))
	}
}