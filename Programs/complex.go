package main

import (
        "fmt"
        "math/rand"
        "time"
)

type Topping struct {
        Name  string
        Price float64
}

type Pizza struct {
        Size     string
        Toppings []Topping
        Price    float64
}

func main() {
        // Seed the random number generator
        rand.Seed(time.Now().UnixNano())

        // Define pizza sizes and prices
        sizes := map[string]float64{
                "Small":     8.99,
                "Medium":    10.99,
                "Large":     12.99,
                "Extra Large": 14.99,
        }

        // Define toppings and prices
        toppings := []Topping{
                {"Pepperoni", 1.50},
                {"Sausage", 1.75},
                {"Mushrooms", 1.25},
                {"Onions", 0.75},
                {"Bell Peppers", 1.00},
                {"Extra Cheese", 1.50},
        }

        // Generate a random pizza size
        size := getRandomKey(sizes)

        // Generate a random number of toppings
        numToppings := rand.Intn(5) + 2 // 2-6 toppings

        // Create a new pizza
        pizza := Pizza{
                Size:     size,
                Toppings: make([]Topping, 0, numToppings),
                Price:    sizes[size],
        }

        // Add random toppings to the pizza
        for i := 0; i < numToppings; i++ {
                topping := toppings[rand.Intn(len(toppings))]
                pizza.Toppings = append(pizza.Toppings, topping)
                pizza.Price += topping.Price
        }

        // Print the pizza order
        fmt.Printf("You ordered a %s pizza with the following toppings:\n", pizza.Size)
        for _, topping := range pizza.Toppings {
                fmt.Printf("- %s\n", topping.Name)
        }
        fmt.Printf("Total price: $%.2f\n", pizza.Price)
}

func getRandomKey(m map[string]float64) string {
        keys := make([]string, 0, len(m))
        for key := range m {
                keys = append(keys, key)
        }
        return keys[rand.Intn(len(keys))]
}
