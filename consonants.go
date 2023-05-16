package main

import (
	"fmt"
	"math"
)

const s = "consonant"

func main(){
   fmt.Println(s)
   const n = 500000000 
   var dd = 3e20
   fmt.Printf("%T\n",dd)
   fmt.Println(float64(dd))
   const d = 3e20/n
   
   //3e20 = 3*10power20

   fmt.Println(d)

   fmt.Println(int64(d))

   fmt.Println(math.Sin(d))

}