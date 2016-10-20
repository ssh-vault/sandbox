package main

import (
	"crypto/md5"
	"encoding/pem"
	"fmt"
	"os"
	"strings"

	"github.com/ssh-vault/ssh2pem"
)

func main() {

	// read in public key from file
	key := os.Getenv("HOME") + "/.ssh/id_rsa.pub"
	outpem, err := ssh2pem.GetPem(key)
	if err != nil {
		panic(err)
	}

	p, _ := pem.Decode(outpem)
	if p == nil {
		panic("No PEM found")
	}

	fmt.Printf("%s", outpem)
	fingerPrint := md5.New()
	fingerPrint.Write(p.Bytes)
	fp := strings.Replace(fmt.Sprintf("% x",
		fingerPrint.Sum(nil)),
		" ",
		":",
		-1)
	fmt.Printf("fp = %+v\n", fp)
}
