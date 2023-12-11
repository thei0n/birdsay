package main

import ( 
	"strings" 
	//"fmt"
)
	func divideString(str string) (string, string) {
		newString:=str[:20]
		oldString:=str[20:]
		return newString, oldString
	}




func format(inputString string)[]string{ 
	//inputString:="Hello there I want to tell you that there exists a buch of stuff that I don't want to talk about like this alksjdflkjasl dkfjlaksjdlfkjalskdjflaksjdlfkajsldkjflaksjdlf"
	//inputString:="salam"

	//inputString:= "aaaa a a a a a a a a a a aaaa a a a a a a  a a a aa a a a a a a a a a a a"
	lineLength:=0
	maxLineLength:=20
	tempString:=""
 	words := strings.Fields(inputString)
	mySlice := []string{}

	
	for i:=0; i<len(words); i++{
		lineLength+=len(words[i])
		
		if (lineLength>=maxLineLength){
			//The case in which the length of the word is more than max line length
			if(len(words[i])>maxLineLength){
				divide:=float64(len(words[i]))/float64(maxLineLength)
				mySlice=append(mySlice, tempString); tempString="";
				

				for y:=0; y<int(divide);y++{
					begin:=maxLineLength*y
					end:=maxLineLength*(y+1)
					mySlice=append(mySlice, words[i][begin:end])
				}
				temp:=int(divide)*maxLineLength

				tempString+=words[i][temp:]
				lineLength=len(words[i])-(maxLineLength*int(divide))
				continue

			}else{
				mySlice=append(mySlice, tempString); tempString="";
				lineLength=len(words[i]);
			}


		}
 
		tempString+=words[i]; tempString+=" ";

	}
	mySlice=append(mySlice, tempString);


	for i:=0;i<len(mySlice);i++{
		//fmt.Println(mySlice[i])
		//for y:=len(mySlice);  
		for y:=len(mySlice[i]); y<25;y++{
			//mySlice=append(mySlice, "")	
			mySlice[i]+=" "
		}

		//fmt.Printf("|    %s    |\n", mySlice[i])

	}
	return mySlice


}

