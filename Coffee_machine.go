package main

import "fmt"

func main() {
	fmt.Println("Starting to make a coffee")
	fmt.Println("Grinding coffee beans")
	fmt.Println("Boiling water")
	fmt.Println("Mixing boiled water with crushed coffee beans")
	fmt.Println("Pouring coffee into the cup")
	fmt.Println("Pouring some milk into the cup")
	fmt.Println("Coffee is ready!")

	var numCoffee int
	const water int = 200
	const milk int = 50
	const beans int = 15

	fmt.Println("Write how many cups of coffee you will need:")

	fmt.Scan(&numCoffee)
	fmt.Println("For", numCoffee, "cups of coffee you will need:")
	fmt.Println(water*numCoffee, "ml of water")
	fmt.Println(milk*numCoffee, "ml of milk")
	fmt.Println(beans*numCoffee, "g of coffee beans")
}
