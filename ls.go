package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"bufio"
)


func main(){

	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()

	fmt.Println("パス:",text)

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
		fmt.Println(err)
	}

	for _,file := range files{
		fmt.Println(file.Name())
	}

	

}
