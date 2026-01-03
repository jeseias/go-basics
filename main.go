package main

import (
	"bufio"
	"fmt"
	"os"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
  fmt.Print(prompt)
  input, err := r.ReadString('\n')
  return input, err
}

func main() {
  // myBill := newBill("James Bond")

  // myBill.addItem("Rice", 2.99)
  // myBill.addItem("Coffee", 7.50)

  // myBill.updateTip(4.26)

  // myBill.updateName("new name")
  r := bufio.NewReader(os.Stdin)

  value, _ := getInput("To whom does this bill belong too? ", r)

  println(value)

}