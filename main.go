package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gabriel-vasile/mimetype"
)

func main() {
	// Open File
	f, err := os.Open("crut.docx")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Get the content
	buf, _ := ioutil.ReadFile("crut.docx")
	mime, extension := mimetype.Detect([]byte(buf))
	fmt.Println("Mime: ", mime, "Extension: ", extension)
}
