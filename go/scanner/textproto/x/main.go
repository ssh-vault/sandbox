package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/textproto"
	"os"
)

func main() {
	client := &http.Client{}
	res, err := client.Get("http://samplecsvs.s3.amazonaws.com/Sacramentorealestatetransactions.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer res.Body.Close()
	reader := bufio.NewReader(res.Body)
	tp := textproto.NewReader(reader)
	for {
		if line, err := tp.ReadLine(); err != nil {
			if err == io.EOF {
				// if file is emtpy
				return
			}
			return
		} else {
			fmt.Printf("%s\n\n", line)
		}
	}
}
