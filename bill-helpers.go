package main

import (
	"bufio"
	"strconv"
)

func askForTip(b *Bill, r *bufio.Reader) {
  newLine()
  input, _ := getInput("How much tip", r)

  p, err := strconv.ParseFloat(input, 64)

  if err != nil {
    log("The price must be a number")
    askForTip(b, r)
  }

  b.updateTip(p)
}

func askForItem(b *Bill, r *bufio.Reader) {
  newLine()
  name, _ := getInput("What is the item name? ", r)   
  price, _ := getInput("What is the item price? ", r)
  
  p, err := strconv.ParseFloat(price, 64)

  if err != nil {
    log("The price must be a number")
    askForItem(b, r)
  }

  b.addItem(name, p)
}