package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	fmt.Println(resp)

	/*
		resp >> Response
		Response is struct which has field Body
		Body is of type interface ReadCloser (basically Body represents the response body.)
		ReadCloser is interface which asks for implementation for Reader -> an interface and closer -> an interface
		Reader asks for implementation for : Read(p []byte) (n int, err error) and Close() error

		http > Response (struct) > Body (ReadCloser Interface) -> Reader (access to Read)

		so this response struct's Body has implemented Reader Interface (i.e. has implemented Read() and Close() so as to satisfy contract)

	*/
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	fmt.Println("=== $$$$$$ +++++++ $$$$$$ ++++++ $$$$$$$ ========")

	//io.Copy(Something that implements Writer Interface, SOmething that implements Reader Interface)
	//io.Copy(os.Stdout, resp.Body)

	fmt.Println("=== $$$$$$ +++++++ $$$$$$ ++++++ $$$$$$$ ========")

	lg := logWriter{}
	io.Copy(lg, resp.Body)
}

type logWriter struct{}

func (logWriter) Write(bs []byte) (size int, err error) {
	fmt.Println(string(bs))
	fmt.Println("This is via custom logWriter", len(bs))
	return len(bs), nil
}
