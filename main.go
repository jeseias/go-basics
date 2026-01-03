package main

import "fmt"


func main() {
  myBill := newBill("James Bond")

  myBill.addItem("Rice", 2.99)
  myBill.addItem("Coffee", 7.50)

  myBill.updateTip(4.26)

  myBill.updateName("new name")

	fmt.Println(myBill)
}