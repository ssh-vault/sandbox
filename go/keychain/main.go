package main

import (
	"fmt"
	"log"
	"os/user"
	"path/filepath"

	keychain "github.com/ssh-vault/go-keychain"
)

func main() {
	//service, account, label, accessGroup, password := "TestGenericPasswordRef", "test", "", "", "toomanysecrets"

	//item := keychain.NewGenericPassword(service, account, label, []byte(password), accessGroup)
	//defer func() { _ = keychain.DeleteItem(item) }()
	//err := keychain.AddItem(item)
	//if err != nil {
	//log.Fatal(err)
	//}

	//passwordAfter, err := keychain.GetGenericPassword(service, account, label, accessGroup)
	//if err != nil {
	//log.Fatal(err)
	//}
	//fmt.Printf("passwordAfter = %s\n", passwordAfter)

	usr, _ := user.Current()
	dir := usr.HomeDir
	keyPath, err := filepath.Abs(filepath.Join(dir, ".ssh/id_rsa"))
	if err != nil {
		log.Fatal(err)
	}
	keyPassword, err := keychain.GetGenericPassword("OpenSSH", keyPath, "SSH: "+keyPath, "com.apple.ssh.passphrases")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("keyPassword = %s\n", keyPassword)
}
