package main

import (
        "fmt"
        "math/rand"
        "time"
)

func main() {
        // Seed the random number generator with the current time
        rand.Seed(time.Now().UnixNano())

        // Generate a random pizza size
        sizes := []string{"Small", "Medium", "Large", "Extra Large"}
        size := sizes[rand.Intn(len(sizes))]

        // Generate a random number of toppings
        numToppings := rand.Intn(5) + 2 // 2-6 toppings

        fmt.Printf("You ordered a %s pizza with %d toppings:\n", size, numToppings)

        // Generate random toppings
        toppings := []string{"Pepperoni", "Sausage", "Mushrooms", "Onions", "Bell Peppers", "Extra Cheese"}
        for i := 0; i < numToppings; i++ {
                topping := toppings[rand.Intn(len(toppings))]
                fmt.Printf("- %s\n", topping)
        }
}
