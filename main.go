package main

import "fmt"
import "os"
import "io/ioutil"

var print = fmt.Println
var printf=fmt.Printf

func main(){

	inputString:=""
	for i:=1; i<len(os.Args);i++{
		inputString+=os.Args[i]
		inputString+=" "
	}


print(" .```````````````````````````````,")
//string:=format()
mySlice:=format(inputString)

content, _ := ioutil.ReadFile("bat.txt")
//printf("' %s                        ' \n", string)
for i:=0;i<len(mySlice);i++{
	fmt.Printf("|    %s    |\n", mySlice[i])
}
fmt.Printf("%s \n",content)
}


