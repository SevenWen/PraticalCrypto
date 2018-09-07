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
		fmt.Println("Wrong Input!!!\nInput Format: vigenere-decrypt <decipherment key> <ciphertext filename>")
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

	cipher:=string(content[:])

	var pass_len float64 = float64(len(passwd))
	var buffer bytes.Buffer
	var shift float64

	i:=0
	//decipher
	for _,char:=range cipher{
		if char>=rune('A') && char<=rune('Z'){
			shift = float64(char)-float64('A')-float64(passwd[int(math.Mod(float64(i),pass_len))])+float64('A')+26
			buffer.WriteString(string(int(math.Mod(float64(shift),26))+int('A')))
			i++
		}
	}
	fmt.Println(buffer.String())
}