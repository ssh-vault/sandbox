package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}
	tmpfile, err := ioutil.TempFile("", "editor")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	cmd := exec.Command(editor, tmpfile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("b = %s\n", b)
}
