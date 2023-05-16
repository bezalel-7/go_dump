package main

import "fmt"

func main(){
	s := make([]string,3)
    s[0] = "a"
	s[1] = "b"
	s[2] = "c"
    s = append(s,"d")
    fmt.Println(s)
    i := make([]int,len(s))

	for j:=0;j<4;j++{
		i[j] = j;
	}

	fmt.Println(i)
	twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

}