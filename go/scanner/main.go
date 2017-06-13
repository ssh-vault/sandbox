package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type Scanner struct {
	*bufio.Scanner
	N int
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{bufio.NewScanner(r), 0}
}

func (s *Scanner) Scan() bool {
	for s.Scanner.Scan() {
		s.N++
		if len(bytes.TrimSpace(s.Bytes())) > 0 {
			return true
		}
	}
	return false
}

func (s *Scanner) LineNumber() int {
	return s.N
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	fmt.Printf("s.LineNumber() = %+v\n", s.LineNumber())
}
