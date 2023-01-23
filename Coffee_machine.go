package main

import (
	"fmt"
)

func main() {

	//es.2
	//fmt.Println("Starting to make a coffee")
	//fmt.Println("Grinding coffee beans")
	//fmt.Println("Boiling water")
	//fmt.Println("Mixing boiled water with crushed coffee beans")
	//fmt.Println("Pouring coffee into the cup")
	//fmt.Println("Pouring some milk into the cup")
	//fmt.Println("Coffee is ready!")

	const water int = 200
	const milk int = 50
	const beans int = 15

	var numCoffee int

	var waterSupply int
	var milkSupply int
	var beansSupply int

	fmt.Println("Write how many ml of water the coffee machine has:")
	fmt.Scan(&waterSupply)
	fmt.Println("Write how many ml of milk the coffee machine has:")
	fmt.Scan(&milkSupply)
	fmt.Println("Write how many grams of coffee beans the coffee machine has:")
	fmt.Scan(&beansSupply)

	var possibleCoffees int
	w := waterSupply / water
	m := milkSupply / milk
	b := beansSupply / beans

	// calculates  the max num of possible coffees with current supply of ingredients
	if w <= m && w <= b {
		possibleCoffees = w
	} else if m <= w && m <= b {
		possibleCoffees = m
	} else {
		possibleCoffees = b
	}

	fmt.Println("Possible coffees: ", possibleCoffees)

	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scan(&numCoffee)

	x := possibleCoffees - numCoffee

	switch {
	case possibleCoffees == numCoffee:
		fmt.Println("Yes, I can make that amount of coffee")
	case possibleCoffees > numCoffee:
		fmt.Printf("Yes, I can make that amount of coffee (and even %d more than that)\n", x)
	default:
		fmt.Printf("No, I can make only %d cups of coffee\n", possibleCoffees)
	}

	//es.1
	//fmt.Println("For", numCoffee, "cups of coffee you will need:")
	//fmt.Println(water*numCoffee, "ml of water")
	//fmt.Println(milk*numCoffee, "ml of milk")
	//fmt.Println(beans*numCoffee, "g of coffee beans")
}
