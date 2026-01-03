package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
  "strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
  newLine()
  fmt.Print(prompt)
  input, err := r.ReadString('\n')
  return strings.TrimSpace(input), err
}

func newLine() {
  log("")
}

func askForTip(b *Bill, r *bufio.Reader) {
  input, _ := getInput("How much tip", r)

  p, err := strconv.ParseFloat(input, 64)

  if err != nil {
    fmt.Println("The price must be a number")
    askForTip(b, r)
  }

  b.updateTip(p)
}

func askForItem(b *Bill, r *bufio.Reader) {
  name, _ := getInput("What is the item name? ", r)   
  price, _ := getInput("What is the item price? ", r)
  
  p, err := strconv.ParseFloat(price, 64)

  if err != nil {
    log("The price must be a number")
    askForItem(b, r)
  }

  b.addItem(name, p)
}

func log(value string) {
  fmt.Println(value)
}

func promptOptions(b Bill) {
  r := bufio.NewReader(os.Stdin)
  option, _ := getInput("Select an action (a - add item, t - add tip, n - update name s-save, x - exit) ", r)

  switch option {
    case "x":
      newLine()
      log("Goodbye")
    case "a":
      askForItem(&b, r)
      promptOptions(b)
  case "t":
    askForTip(&b, r)
  case "":
    log("Tip was saved with success")
  default:  
    log("That was an invalid option")
    promptOptions(b)
  }

  fmt.Println(b)
}

func main() {
  r := bufio.NewReader(os.Stdin)

  name, _ := getInput("To whom does this bill belong too? ", r)
  myBill := newBill(name)
  newLine()

  promptOptions(myBill)

}