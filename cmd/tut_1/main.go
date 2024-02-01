package main

import (
	"fmt"
	"unicode/utf8"
)




func main(){
	fmt.Println("Hello World")
	var num int16 = 10;
	fmt.Println(num)
	
	var sayHey string  = "Hello \nWorld"
	fmt.Println(sayHey)

	const PI float64 = 3.1415
	fmt.Println(PI)

	// const thisWillNotFindLengthOFTheString = len('test')

	fmt.Println(len("faizan")) // This will output number of bytes of the string and not the actual length of the string

	fmt.Println(utf8.RuneCountInString("Faizan")) // This will output the length of the string
}