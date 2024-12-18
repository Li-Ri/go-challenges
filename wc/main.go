package main

import (
	"fmt"
	"os"
	"strings"
)

type WcReader struct {
  content string
}

func main() {

  filename := os.Args[1:][0]

  if(len(filename) == 0) {
    panic("no file name given")
  }
  data, err := os.ReadFile(filename)

  if(err != nil) {
    panic("unable to read file")
  }
  content := string(data)
  reader := NewWcReader(content)

  fmt.Printf("%d, %d, %d %s \n", reader.getNumberOfLines(), reader.getWordCount(), reader.getBytes(), filename)
}

func (r WcReader) getNumberOfLines() int {
  lines := strings.Split(r.content, "\n")
  return len(lines[:len(lines) -1])
}

func (r WcReader) getBytes() int {
  return len(r.content)
}

func (r WcReader) getWordCount() int {
  return len(strings.Split(r.content, " "))
}

func NewWcReader(content string) WcReader {
  return WcReader{ 
    content,
  }
}
