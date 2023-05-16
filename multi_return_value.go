package main

import "fmt"

func vals(a,b int)(int,int){
   return a+b,a*b
}

func main(){
   a,b  := vals(1,2)
   fmt.Println(a,b)
}
