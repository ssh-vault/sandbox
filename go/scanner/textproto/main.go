package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
)

func main() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\n%s\n%s\n%s\n",
			"col1,col2,col3",
			"1,2,3",
			"a,b,c",
			"x,y,x",
		)
	}))
	defer ts.Close()

	client := &http.Client{}
	res, err := client.Get(ts.URL)
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
