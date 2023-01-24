package main

import "fmt"

var waterSupply int = 400
var milkSupply int = 540
var coffeeSupply int = 120
var cupsSupply int = 9
var moneySupply int = 550

func askAction() string {
	fmt.Println("Write action (buy, fill, take, remaining, exit):")
	var action string
	fmt.Scan(&action)
	return action
}

func remaining() {
	fmt.Println("The coffee machine has:")
	fmt.Println(waterSupply, "ml of water")
	fmt.Println(milkSupply, "ml of milk")
	fmt.Println(coffeeSupply, "g of coffee beans")
	fmt.Println(cupsSupply, "disposable cups")
	fmt.Println("$", moneySupply, "of money")
}

func checkResources(w int, m int, b int, c int) bool {
	if w > waterSupply {
		fmt.Println("Sorry, not enough water!")
		return false
	} else if m > milkSupply {
		fmt.Println("Sorry, not enough milk!")
		return false
	} else if b > coffeeSupply {
		fmt.Println("Sorry, not enough coffee beans!")
		return false
	} else if c > cupsSupply {
		fmt.Println("Sorry, not enough disposable cups!")
		return false
	}

	fmt.Println("I have enough resources, making you a coffee!")
	return true
}

func buy() {

	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	var productCode string
	fmt.Scan(&productCode)

	if productCode == "1" && checkResources(250, 0, 16, 1) {
		waterSupply -= 250
		coffeeSupply -= 16
		cupsSupply--
		moneySupply += 4
	} else if productCode == "2" && checkResources(350, 75, 20, 1) {
		waterSupply -= 350
		milkSupply -= 75
		coffeeSupply -= 20
		cupsSupply--
		moneySupply += 7
	} else if productCode == "3" && checkResources(200, 100, 12, 1) {
		waterSupply -= 200
		milkSupply -= 100
		coffeeSupply -= 12
		cupsSupply--
		moneySupply += 6
	}
}

func fill() {
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
}

func main() {

Loop:
	for {
		switch askAction() {
		case "remaining":
			remaining()
		case "buy":
			buy()
		case "fill":
			fill()
		case "take":
			fmt.Println("I gave you", moneySupply, "$")
			moneySupply = 0
		case "exit":
			break Loop
		default:
			fmt.Println("Not a valid action")
		}
	}
}
