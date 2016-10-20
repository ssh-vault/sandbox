package main

import (
	"crypto/md5"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ssh-vault/sandbox/go/ssh/ssh"
)

func main() {

	// read in public key from file
	bytes, err := ioutil.ReadFile(os.Getenv("HOME") + "/.ssh/id_rsa.pub")
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	// decode string ssh-rsa format to native type
	pub_key, err := ssh.DecodePublicKey(string(bytes))
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	// pub_key is of type *rsa.PublicKey

	// Marshal to ASN.1 DER encoding
	pkix, err := x509.MarshalPKIXPublicKey(pub_key)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	// Encode to PEM format
	outpem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pkix,
	})

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
