package main

import "fmt"

type Command interface {
    execute()
}

type BurgerCommand struct {
    burger *Burger
}

func (c *BurgerCommand) execute() {
    c.burger.makeBurger()
}

type Burger struct {
    Name       string
    Price      float64
    Ingredients []string
}

func (b *Burger) makeBurger() {
    fmt.Printf("Making a %s burger...\n", b.Name)
    fmt.Println("Ingredients:")
    for _, ingredient := range b.Ingredients {
        fmt.Printf("- %s\n", ingredient)
    }
    fmt.Printf("Price: $%.2f\n", b.Price)
    fmt.Println("Burger is ready!\n")
}

type Waiter struct {
    orders []Command
}

func (w *Waiter) takeOrder(order Command) {
    w.orders = append(w.orders, order)
}

func (w *Waiter) serveOrders() {
    fmt.Println("Waiter is serving orders:")
    for _, order := range w.orders {
        order.execute()
    }
}

func main() {
    burger1 := &Burger{
        Name: "Cheeseburger",
        Price: 5.99,
        Ingredients: []string{"Beef patty", "Cheese", "Bun", "Lettuce", "Tomato", "Onion"},
    }

    burger2 := &Burger{
        Name: "Vegetarian Burger",
        Price: 4.99,
        Ingredients: []string{"Veggie patty", "Bun", "Lettuce", "Tomato", "Onion"},
    }

    order1 := &BurgerCommand{burger: burger1}
    order2 := &BurgerCommand{burger: burger2}

    waiter := &Waiter{}
    waiter.takeOrder(order1)
    waiter.takeOrder(order2)

    waiter.serveOrders()
}
