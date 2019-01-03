package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))


	s1 := rand.NewSource(42)
	r1 := rand.New(s1)
	fmt.Println(r1)
}
