package main

import "fmt"

type vendingmachine struct {
	InsertedMoney int
	ChangedMoney int
	coins	map[string]int
	items	map[string]int
}

func (m vendingmachine) InsertMoney() int {
	return m.InsertedMoney
}

func (m *vendingmachine) Insertcoin(coin string) {
	m.InsertedMoney += m.coins[coin]
}

func (m *vendingmachine) Resetcoin() {
	m.InsertedMoney = 0
}

func (m *vendingmachine) SelectSD() string {
	price := m.items["SD"]
	change := m.InsertedMoney - price
	return "SD" + m.change(change)
}

func (m *vendingmachine) SelectCC() string {
	price := m.items["CC"]
	change := m.InsertedMoney - price
	m.ChangedMoney = change
	return "CC" + m.change(change)
}

func (m *vendingmachine) SelectDW() string {
	price := m.items["DW"]
	change := m.InsertedMoney - price
	return "DW" + m.change(change)
}

func (m *vendingmachine) change(c int) string {
	var str string
	values := [...]int{10, 5, 2, 1}
	coins := [...]string{"T", "F", "TW", "O"}

	for i := 0; i < len(values); i++ {
		if c >= values[i] {
			str += ", " + coins[i]
			c -= values[i]
			i--
		}
	}
	return str
}

func (vm *vendingmachine) CoinReturn() string {
	coins := vm.change(vm.InsertedMoney)
	vm.InsertedMoney = 0
	return coins[2:len(coins)]
}


func main() {
	var coins = map[string]int{"T": 10, "F": 5, "TW": 2, "O": 1}
	var items = map[string]int{"SD": 18, "CC": 12, "DW": 7}
	vm := vendingmachine{coins: coins, items: items}
	//Inserted Money: 18
	vm.Insertcoin("T")
	vm.Insertcoin("F")
	vm.Insertcoin("TW")
	vm.Insertcoin("O")
	fmt.Println("Inserted Money:", vm.InsertMoney())
	can := vm.SelectSD()
	fmt.Println(can)

	vm.Resetcoin()
	//Inserted Money: 18
	vm.Insertcoin("T")
	vm.Insertcoin("TW")
	fmt.Println("Inserted Money:", vm.InsertMoney())
	can = vm.SelectCC()
	fmt.Println(can)

	vm.Resetcoin()
	//Inserted Money: 7
	vm.Insertcoin("F")
	vm.Insertcoin("TW")
	fmt.Println("Inserted Money:", vm.InsertMoney())
	can = vm.SelectDW()
	fmt.Println(can)

	vm.Resetcoin()
	//Inserted Money: 25
	vm.Insertcoin("T")
	vm.Insertcoin("T")
	vm.Insertcoin("F")
	fmt.Println("Inserted Money:", vm.InsertMoney())
	can = vm.SelectCC()
	fmt.Println(can)
	coin := vm.CoinReturn()
	fmt.Println("Returned Money:", coin)
}