package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"io"
)

type name struct{
	fname string
	lname string
}

func fixLongName(s string) string {
	if len(s)>20 {
       return string(s[0:20])
    } else {
       return s 
    }
}

func main() {

	path:=""
    fmt.Println("Please input the file path: ")
    fmt.Scanln(&path)

	var Name[] name


	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Failed to open file: %s", err)
	}

	defer file.Close()
	
	rd := bufio.NewReader(file)
	
	for {
        line, _, err:= rd.ReadLine()
        
        if err != nil || io.EOF == err {
            break
        }       
            words:=strings.Split(string(line)," ")
            tpwords:= name {
                fixLongName(words[0]),
                fixLongName(words[len(words)-1]),
            } 
            Name=append(Name,tpwords)
        } 

    for i:=0;i<len(Name);i++{
        fmt.Println(i+1)
        fmt.Println("First Name:"+Name[i].fname)
        fmt.Println("Last Name:"+Name[i].lname)
    } 

}


