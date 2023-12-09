package main

import ( 
	"strings" 
)
	func divideString(str string) (string, string) {
		newString:=str[:20]
		oldString:=str[20:]
		return newString, oldString
	}



func format(fullString string) string{

// Your original string
//fullString := "hello my name is not alksdlkf aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa something that you would want to know but for me it is the most important thing out therea"

outString := ""

	// Split the string into words
	words := strings.Fields(fullString)

	linelength:=0

	for i:=0; i<len(words); i++{
		linelength+=len(words[i])
		if linelength>=20{

			if len(words[i])>20{
				word:=words[i]
				for len(word)>20{
					var add string;
					add, word = divideString(word)
					outString+="\n"
					outString+="' "
					outString+=add
					if(len(word)<20){
						outString+="\n"
						outString+="' "
						outString+=word
						outString+=" "
						linelength=len(word)
					}else{linelength=0}
				}
			}else{
			outString+="\n"
			outString+="' "
			linelength=0
		}
		}else{
		outString+=words[i]
		outString+=" "
	}
	}

	//fmt.Println(outString)
	return outString

}
