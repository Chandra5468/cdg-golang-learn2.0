package main

import (
	"fmt"
)

type order struct {
	id     string
	amount float32
	status string
	// createdAt time.Time // nanosecond precision
	// delivered bool
}

// IN GO, there is no constructor. To implement a function as constructor we can implement using this below approach

func NewOrder(id string, amount float32, status string) *order {
	return &order{
		id:     id,
		amount: amount,
		status: status,
	}
}

type xyz struct {
	abc string
	x   *order
}

func main() {
	// order := order{
	// 	id: "one",
	// }
	// // order.createdAt = time.Now().UTC()
	// order.amount = 500
	// order.status = "cooking"
	// // order.delivered = false
	// fmt.Print(order)

	// Calling constructor of a class : calling a function which gives us an instantiated struct with values

	ord := NewOrder("123", 45.4, "pending")

	fmt.Println(*ord)

	// STRUCT EMBEDDINGS (Inheritence, Composition)

	z := &xyz{
		abc: "abc",
		x: &order{
			id:     "141",
			status: "ordered",
		},
	}

	fmt.Println("This is embedded struct ", *z.x)

	// z.x.
}
