package main

import (
	"time"
	"fmt"
)

func main() {
	s := time.Second;
	fmt.Println(s)

	t:=time.Unix(10,10)
	fmt.Println(t)
}
