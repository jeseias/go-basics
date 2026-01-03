package main

import (
	"bufio"
	"fmt"
	"strings"
)

func log(value string) {
  fmt.Println(value)
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
  fmt.Print(prompt)
  input, err := r.ReadString('\n')
  return strings.TrimSpace(input), err
}

func newLine() {
  log("")
}