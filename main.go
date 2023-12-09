package main

import "fmt"
import "os"
import "io/ioutil"

var print = fmt.Println
var printf=fmt.Printf

func main(){

str:="CiCi"

print( " .```````,")

printf("' %s    ' \n", str)

file, _ := os.Open("bat.txt")
out, _ :=ioutil.ReadAll(file)
output := string(out[:])
fmt.Println(output)

}


