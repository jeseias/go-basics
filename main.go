package main

import (
	"bufio"
	"fmt"
	"os"
)

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
    promptOptions(b)
  case "s":
    b.save()
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