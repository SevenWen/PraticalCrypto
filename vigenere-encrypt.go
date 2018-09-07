package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"math"
	"bytes"
	"os"
)

func main(){
	//get user input
	input:=os.Args

	if len(input)!=3{
		fmt.Println("Wrong Input!!!\nInput Format: vigenere-encrypt <encipherment key> <plaintext filename>")
		os.Exit(0)
	}
	passwd:=input[1]
	passwd=strings.ToUpper(passwd)
	filename:=input[2]

	// plaintx_len:=len(plaintx)
	// passwd:="PASS"

	//read file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File Reading Error!!!")
		fmt.Println(err)
		os.Exit(0)
	}

	// fmt.Printf("File contents: %s", content)

	plaintx:=string(content[:])
	plaintx=strings.ToUpper(plaintx)
	// fmt.Println(plaintx)

	var pass_len float64 = float64(len(passwd))
	var buffer bytes.Buffer
	var shift float64
	i:=0
	//encipher
	for  _,char:=range plaintx{
		if char>=rune('A') && char<=rune('Z'){
			shift=float64(char)-float64('A')+float64(passwd[int(math.Mod(float64(i),pass_len))])-float64('A')
			buffer.WriteString(string(int(math.Mod(float64(shift),26))+int('A')))
			i++
		}
	}
	// fmt.Println()
	var cipher string = buffer.String()
	fmt.Println(cipher)

}