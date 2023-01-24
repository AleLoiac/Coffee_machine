package main

import "fmt"

var waterSupply int = 400
var milkSupply int = 540
var coffeeSupply int = 120
var cupsSupply int = 9
var moneySupply int = 550

func stats() {
	fmt.Println("The coffee machine has:")
	fmt.Println(waterSupply, "ml of water")
	fmt.Println(milkSupply, "ml of milk")
	fmt.Println(coffeeSupply, "g of coffee beans")
	fmt.Println(cupsSupply, "disposable cups")
	fmt.Println("$", moneySupply, "of money")
}

func buy(cod int) {
	if cod == 1 {
		waterSupply = waterSupply - 250
		coffeeSupply = coffeeSupply - 16
		cupsSupply--
		moneySupply = moneySupply + 4
	} else if cod == 2 {
		waterSupply = waterSupply - 350
		milkSupply = milkSupply - 75
		coffeeSupply = coffeeSupply - 20
		cupsSupply--
		moneySupply = moneySupply + 7
	} else if cod == 3 {
		waterSupply = waterSupply - 200
		milkSupply = milkSupply - 100
		coffeeSupply = coffeeSupply - 12
		cupsSupply--
		moneySupply = moneySupply + 6
	} else {
		fmt.Println("Wrong number")
	}
}

func main() {

	stats()

	fmt.Println("Write action (buy, fill, take):")
	var action string
	fmt.Scan(&action)

	switch action {
	case "buy":
		fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
		var productCode int
		fmt.Scan(&productCode)
		buy(productCode)
		stats()
	case "fill":
		fmt.Println("Write how many ml of water you want to add:")
		var w int
		fmt.Scan(&w)
		waterSupply = waterSupply + w
		fmt.Println("Write how many ml of milk you want to add:")
		var m int
		fmt.Scan(&m)
		milkSupply = milkSupply + m
		fmt.Println("Write how many grams of coffee beans you want to add:")
		var g int
		fmt.Scan(&g)
		coffeeSupply = coffeeSupply + g
		fmt.Println("Write how many disposable cups you want to add:")
		var c int
		fmt.Scan(&c)
		cupsSupply = cupsSupply + c
		stats()
	case "take":
		fmt.Println("I gave you", moneySupply, "$")
		moneySupply = 0
		stats()
	default:
		fmt.Println("Invalid action")
		break
	}
}
