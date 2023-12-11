package main

import 
(
	"fmt"
	"os"
)

func main(){
	
	//Turning positional arguments into a string
	inputString:=""
	for i:=1; i<len(os.Args);i++{
		inputString+=os.Args[i]
		inputString+=" "
	}

fmt.Println(" .```````````````````````````````,")

//Formatting the input string
mySlice:=format(inputString)

for i:=0;i<len(mySlice);i++{
	fmt.Printf("|    %s    |\n", mySlice[i])
}
	bat()
}

