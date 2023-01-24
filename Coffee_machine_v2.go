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

func buy() {

	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	var productCode string
	fmt.Scan(&productCode)

	if productCode == "1" && waterSupply >= 250 && coffeeSupply >= 16 && cupsSupply >= 1 {
		waterSupply -= 250
		coffeeSupply -= 16
		cupsSupply--
		moneySupply += 4
		fmt.Println("I have enough resources, making you a coffee!")
	} else if productCode == "2" && waterSupply >= 350 && milkSupply >= 75 && coffeeSupply >= 20 && cupsSupply >= 1 {
		waterSupply -= 350
		milkSupply -= 75
		coffeeSupply -= 20
		cupsSupply--
		moneySupply += 7
		fmt.Println("I have enough resources, making you a coffee!")
	} else if productCode == "3" && waterSupply >= 200 && milkSupply >= 100 && coffeeSupply >= 12 && cupsSupply >= 1 {
		waterSupply -= 200
		milkSupply -= 100
		coffeeSupply -= 12
		cupsSupply--
		moneySupply += 6
		fmt.Println("I have enough resources, making you a coffee!")

	} else if productCode == "1" && waterSupply < 250 {
		fmt.Println("Sorry, not enough water!")
	} else if productCode == "1" && coffeeSupply < 16 {
		fmt.Println("Sorry, not enough coffee!")
	} else if productCode == "1" && cupsSupply < 1 {
		fmt.Println("Sorry, not enough cups!")

	} else if productCode == "2" && waterSupply < 350 {
		fmt.Println("Sorry, not enough water!")
	} else if productCode == "2" && milkSupply < 75 {
		fmt.Println("Sorry, not enough milk!")
	} else if productCode == "2" && coffeeSupply < 20 {
		fmt.Println("Sorry, not enough coffee!")
	} else if productCode == "2" && cupsSupply < 1 {
		fmt.Println("Sorry, not enough cups!")

	} else if productCode == "2" && waterSupply < 200 {
		fmt.Println("Sorry, not enough water!")
	} else if productCode == "2" && milkSupply < 100 {
		fmt.Println("Sorry, not enough milk!")
	} else if productCode == "2" && coffeeSupply < 12 {
		fmt.Println("Sorry, not enough coffee!")
	} else if productCode == "2" && cupsSupply < 1 {
		fmt.Println("Sorry, not enough cups!")

	} else if productCode == "back" {

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
