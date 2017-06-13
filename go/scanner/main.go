package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var b bytes.Buffer
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if len(bytes.TrimSpace(scanner.Bytes())) > 0 {
			b.WriteString(fmt.Sprintf("%s\n", scanner.Text()))
		}
	}
	fmt.Printf("%s", b.Bytes())
}
