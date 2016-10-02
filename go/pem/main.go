package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	pem_data, err := ioutil.ReadFile("/tmp/id_rsa")
	if err != nil {
		log.Fatalf("Error reading pem file: %s", err)
	}
	block, _ := pem.Decode(pem_data)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		log.Fatal("No valid PEM data found")
	}
	pemOut, err := x509.DecryptPEMBlock(block, []byte("secret"))
	if err != nil {
		log.Fatalf("err = %+v\n", err)
	}
	key, err := x509.ParsePKCS1PrivateKey(pemOut)
	if err != nil {
		log.Fatalf("err = %+v\n", err)
	}
	pemdata := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	)
	fmt.Printf("pemdata = %s\n", pemdata)
}
