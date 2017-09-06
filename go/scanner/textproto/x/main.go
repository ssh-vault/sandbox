package main

import (
	"bufio"
	"fmt"
	"net/http"
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
	scanner := bufio.NewScanner(res.Body)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		c := scanner.Text()
		switch c {
		case "\r":
			fmt.Println()
		default:
			fmt.Printf("%s", c)
		}
	}
}
