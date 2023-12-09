package main

import "fmt"
//import "os"
import "io/ioutil"

var print = fmt.Println
var printf=fmt.Printf

func main(){

print(" .```````````````````````````,")
string:=format("lalkasjdflkjasdljfalksdjflakjsdlfjalskdjflkasjdlfkjals alksdj flkasjd lfkajsldkfjalasjdlfaskldfjlaksdjf")
content, _ := ioutil.ReadFile("bat.txt")
printf("' %s                        ' \n", string)
fmt.Printf("%s \n",content)
}


