package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {

	// ssh-keygen -t rsa -f /tmp/key
	key, err := ioutil.ReadFile("/tmp/key")
	if err != nil {
		log.Fatal(err)
	}

	i, err := ssh.ParseRawPrivateKeyWithPassphrase(key, []byte("secret"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("i= %+v\n", i)
}
