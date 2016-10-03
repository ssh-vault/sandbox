package main

import (
	"crypto/rand"
	"fmt"
)

func main() {

	c := 64
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("64 = %X\n", b)
}
