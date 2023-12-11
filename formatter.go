package main

import "strings"
	
func format(inputString string)[]string{ 

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
				divide:=len(words[i])/maxLineLength
				mySlice=append(mySlice, tempString); tempString="";

				for y:=0; y<divide;y++{
					begin:=maxLineLength*y
					end:=maxLineLength*(y+1)
					mySlice=append(mySlice, words[i][begin:end])
				}
				temp:=divide*maxLineLength

				tempString+=words[i][temp:]
				lineLength=len(words[i])-(maxLineLength*divide)
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
		for y:=len(mySlice[i]); y<25;y++{
			mySlice[i]+=" "
		}
	}

	return mySlice
}
