package main

import (
	"fmt"
)

func main(){
	a := 5
	if a>10 {
		fmt.Println("a>10",a)
	}else if a>15{
		fmt.Println("a>15",a)
	}else if a==5{
		fmt.Println("a=5",a)
	}
}