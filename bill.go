package main

import (
	"fmt"
	"os"
)

type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) Bill {
	b := Bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

func (b *Bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *Bill) updateTip(value float64) {
	b.tip = value
}

func (b *Bill) updateName(value string) {
	b.name = value
}

func (b *Bill) format() string {
	fs := "Bill breakdown:\n"
	var total float64 = 0

	for key, value := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n",key+":", value)
    total += value
	}

  fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)

	return fs
}

func (b *Bill) save() {
	data := []byte(b.format())
  err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
  if err != nil {
    panic(err)
  }

  log("Bill saved to file")
}