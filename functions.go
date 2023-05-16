package main

import "fmt"

func plus(a int,b int)int{
	return a+b
}

func plusplus(a,b,c int)int{
	return a+b+c
}

func main(){
	fmt.Print(plus(1,2)," ")
	fmt.Print(plusplus(1,2,3))
}