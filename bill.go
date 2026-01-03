package main 

type Bill struct {
  name string 
  items map[string]float64 
  tip float64
}

func newBill(name string) Bill {
  b := Bill{
    name: name,
    items: map[string]float64{},
    tip: 0,
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