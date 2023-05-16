package main

import "fmt"

func main(){
	i := 1
	for i <= 3{
		fmt.Println(i)
		i = i + 1
	}

	for i=1;i<5;i++{
		fmt.Printf("%d ",i)
	}

	for {
        fmt.Println("loop")
        break
    }

	for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}