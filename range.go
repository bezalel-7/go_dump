package main

import "fmt"

func main(){
	m := map[string]int{"a":1,"b":2}
	for i,j := range m{
		fmt.Println(i,j)
	}

	h := [...]int{5,4,3,2,1,0}

	for  i,j := range h{
		fmt.Println(i,j)
	}

	for i,j := range "varaDudi"{
		fmt.Println(i,string(j))
	}

	s := "dafda"
    p := "bdafa"

	k := p + s
	fmt.Println(k)
}